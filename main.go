package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

var (
	path           = flag.String("path", "", "help")
	data           = make(map[string][]string)
	typedefs_basic = make(map[string]string)
	typedefs       = make(map[string]string)
	interfaces     = make(map[string]map[string]bool)
)

func init() {
	flag.Parse()
}

/*

	&Struct{simpleFiled: org.simpleFiled,
		list: cloneList(org.list),
		simplePointerField: newSimplePointer(org.simplePointerField)
		structPointerField: clone_Struct(org.otherStruct)
		LATER:chan: make(chan type)
		}


*/

func filter(fileInfo os.FileInfo) bool {
	return !strings.HasSuffix(fileInfo.Name(), "_clone.go")
}

func findInterfaces(pkgMap map[string]*ast.Package) {
	for _, pkg := range pkgMap {
		interfaces[pkg.Name] = make(map[string]bool)
		for _, file := range pkg.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				switch x := n.(type) {
				case *ast.TypeSpec:
					_, isInterface := x.Type.(*ast.InterfaceType)
					if isInterface {
						interfaces[pkg.Name][x.Name.String()] = true
					}
				}
				return true
			})
		}
	}
}

func main() {
	fset := token.NewFileSet()
	pkgMap, err := parser.ParseDir(fset, *path, filter, parser.AllErrors)
	if err != nil {
		log.Println(err)
		return
	}

	findInterfaces(pkgMap)

	for _, pkg := range pkgMap {
		for _, file := range pkg.Files {

			currentStruct := ""

			ast.Inspect(file, func(n ast.Node) bool {
				switch x := n.(type) {
				case *ast.TypeSpec:
					simple, isTypedef := x.Type.(*ast.Ident)

					if isTypedef {
						if simple.Obj != nil {
							typedefs[x.Name.Name] = simple.Name
						} else {
							typedefs_basic[x.Name.Name] = simple.Name
						}

					} else {
						currentStruct = x.Name.String()
					}

				case *ast.StructType:
					for _, field := range x.Fields.List {
						switch x := field.Type.(type) {
						case *ast.Ident:
							if x.Obj != nil && len(field.Names) == 0 {
								// fmt.Println(x.Obj.Name, x.Obj.Kind)
							}
							for _, name := range field.Names {

								if x.Obj != nil {
									if _, ok := interfaces[pkg.Name][x.Obj.Name]; ok {
										add(currentStruct, fmt.Sprintf("%v: Clone_%v(obj.%v),", name.Name, x.Obj.Name, name.Name))
									} else {
										add(currentStruct, fmt.Sprintf("%v: *obj.%v.Clone(),", name.Name, name.Name))
									}
								} else {
									add(currentStruct, fmt.Sprintf("%v: obj.%v,", name.Name, name.Name))
								}
							}
						case *ast.StarExpr:

							for _, name := range field.Names {

								s, ok := x.X.(*ast.Ident)
								if ok && s.Obj == nil {
									add(currentStruct, fmt.Sprintf("%v: deepcopy.Clone_%v(obj.%v),", name.Name, x.X, name.Name))
								} else {
									add(currentStruct, fmt.Sprintf("%v: obj.%v.Clone(),", name.Name, name.Name))
								}

							}

						case *ast.ArrayType:
							for _, name := range field.Names {
								ptr, isPtr := x.Elt.(*ast.StarExpr)
								if isPtr {
									add(currentStruct, fmt.Sprintf("%v: cloneListPtr_%v(obj.%v),", name.Name, ptr.X, name.Name))
								} else {
									s, ok := x.Elt.(*ast.Ident)
									if ok && s.Obj == nil {
										add(currentStruct, fmt.Sprintf("%v: deepcopy.CloneList_%v(obj.%v),", name.Name, x.Elt, name.Name))
									} else {
										add(currentStruct, fmt.Sprintf("%v: cloneList_%v(obj.%v),", name.Name, x.Elt, name.Name))
									}
								}
							}

						case *ast.ChanType:
							// TODO

						case *ast.SelectorExpr:
							add(currentStruct, fmt.Sprintf("%v: obj.%v,", x.Sel.Name, x.Sel.Name))
						}
					}

				}
				return true
			})

			outfile, err := os.Create(*path + "/" + file.Name.Name + "_clone.go")
			if err != nil {
				log.Println(err)
			}

			outfile.WriteString("package " + pkg.Name + "\n\n")
			outfile.WriteString("import \"github.com/wookesh/deepcopy/deepcopy\"\n\n")

			for interface_ := range interfaces[pkg.Name] {
				generateInterface(outfile, interface_)
			}

			for class, strList := range data {
				generateFunctions(outfile, class, strList)
			}

			for typedef, class := range typedefs {
				generateFunctions(outfile, typedef, data[class])
			}

			for typedef, _ := range typedefs_basic {
				generateTypedefBasic(outfile, typedef)
			}
		}

	}

}

func generateInterface(outfile *os.File, interf string) {

	outfile.WriteString("func Clone_" + interf + "(obj " + interf + ") " + interf + " {\n")
	outfile.WriteString("\tx, ok := interface{}(obj).(deepcopy.DeepCopier)\n")
	outfile.WriteString("\tif ok {\n")
	outfile.WriteString("\t\ti, ok := x.CloneInterface().(" + interf + ")\n")
	outfile.WriteString("\t\tif ok {\n")
	outfile.WriteString("\t\t\treturn i\n")
	outfile.WriteString("\t\t}\n")
	outfile.WriteString("\t\treturn nil\n")
	outfile.WriteString("\t}\n")
	outfile.WriteString("\treturn nil\n")
	outfile.WriteString("}\n\n")
}

func generateFunctions(outfile *os.File, class string, fields []string) {
	// SingleGenerator
	outfile.WriteString("func (obj *" + class + ") Clone() *" + class + " {\n")
	outfile.WriteString("\tif obj == nil {\n")
	outfile.WriteString("\t\treturn nil\n")
	outfile.WriteString("\t}\n")
	outfile.WriteString("\treturn &" + class + "{\n")
	for _, str := range fields {
		outfile.WriteString("\t\t" + str + "\n")
	}
	outfile.WriteString("\t}\n")
	outfile.WriteString("}\n\n")

	// InterfaceCloneGenerator
	outfile.WriteString("func (obj *" + class + ") CloneInterface() deepcopy.DeepCopier {\n")
	outfile.WriteString("\treturn obj.Clone()\n")
	outfile.WriteString("}\n\n")

	// ListGenerator
	outfile.WriteString("func cloneList_" + class + "(objs []" + class + ") []" + class + " {\n")
	outfile.WriteString("\tretList := make([]" + class + ", 0, len(objs))\n")
	outfile.WriteString("\tfor _, obj := range objs {\n")
	outfile.WriteString("\t\tretList = append(retList, *obj.Clone())\n")
	outfile.WriteString("\t}\n")
	outfile.WriteString("\treturn retList\n")
	outfile.WriteString("}\n\n")

	// ListPtrGenerator
	outfile.WriteString("func cloneListPtr_" + class + "(objs []*" + class + ") []*" + class + " {\n")
	outfile.WriteString("\tretList := make([]*" + class + ", 0, len(objs))\n")
	outfile.WriteString("\tfor _, obj := range objs {\n")
	outfile.WriteString("\t\tretList = append(retList, obj.Clone())\n")
	outfile.WriteString("\t}\n")
	outfile.WriteString("\treturn retList\n")
	outfile.WriteString("}\n\n")
}

func generateTypedefBasic(outfile *os.File, class string) {
	// SingleGenerator
	outfile.WriteString("func (obj *" + class + ") Clone() *" + class + " {\n")
	outfile.WriteString("\tif obj == nil {\n")
	outfile.WriteString("\t\treturn nil\n")
	outfile.WriteString("\t}\n")
	outfile.WriteString("\tret := " + class + "(*obj)\n")
	outfile.WriteString("\treturn &ret\n")
	outfile.WriteString("}\n\n")

	// ListGenerator
	outfile.WriteString("func cloneList_" + class + "(objs []" + class + ") []" + class + " {\n")
	outfile.WriteString("\tretList := make([]" + class + ", 0, len(objs))\n")
	outfile.WriteString("\tfor _, obj := range objs {\n")
	outfile.WriteString("\t\tretList = append(retList, *obj.Clone())\n")
	outfile.WriteString("\t}\n")
	outfile.WriteString("\treturn retList\n")
	outfile.WriteString("}\n\n")

	// ListPtrGenerator
	outfile.WriteString("func cloneListPtr_" + class + "(objs []*" + class + ") []*" + class + " {\n")
	outfile.WriteString("\tretList := make([]*" + class + ", 0, len(objs))\n")
	outfile.WriteString("\tfor _, obj := range objs {\n")
	outfile.WriteString("\t\tretList = append(retList, obj.Clone())\n")
	outfile.WriteString("\t}\n")
	outfile.WriteString("\treturn retList\n")
	outfile.WriteString("}\n\n")
}

func add(class, field string) {
	l, ok := data[class]
	if !ok {
		l = make([]string, 0)
	}
	data[class] = append(l, field)
}
