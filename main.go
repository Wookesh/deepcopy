package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"
	"strings"
)

var (
	path           = flag.String("path", "", "help")
	data           = make(map[string][]string)
	typedefs_basic = make(map[string]string)
	typedefs       = make(map[string]string)
	interfaces     = make(map[string]bool)

	structs       = make(map[string][]FieldCloner)
	currentStruct string
)

func init() {
	flag.Parse()
}

/*
	type Struct struct {
		AnonymousStruct
		*AnonymousStruct2
		package.ExternalAnonymous
		simpleField             int
		simpleFieldPtr          *int
		structField             Struct
		structFieldPtr          *StructPtr
		interface               Interface
		externalStructField:    package.ExternalStruct
		externalStructFieldPtr: *package.ExternalStruct
		externalInterface       package.ExternalInterface
	}


	&Struct{
		AnonymousStruct:           *org.AnonymousStruct.Clone(),
		AnonymousStruct2:          org.AnonymousStruct2.Clone(),
		package.ExternalAnonymous: org.ExternalAnonymous, || [TODO] *org.ExternalAnonymous.Clone(),
		simpleField:               org.simmpleField,
		simpleFieldPtr:            deepcopy.clone_int(org.simpleFieldPtr),
		structField"               *org.structField.Clone(),
		structFieldPtr:            org.structFieldPtr.Clone(),
		interface                  clone_InterfaceType(org.interface),
		externalStructField:       [TODO] clone_by_reflect(org.externalStructField),    || *org.externalStructField.Clone(),
		externalStructFieldPtr:    [TODO] clone_by_reflect(org.externalStructFieldPtr), || org.externalStructFieldPtr.Clone(),
		externalInterface:         [TODO] clone_by_reflect(org.externalInterface),      || package.Clone_ExternalInterface(org.externalInterface),
		}


*/

func filter(fileInfo os.FileInfo) bool {
	return !strings.HasSuffix(fileInfo.Name(), "_clone.go")
}

func run(pkgMap map[string]*ast.Package) {
	for _, pkg := range pkgMap {
		for fileName, file := range pkg.Files {
			log.Println(fileName)

			ast.Inspect(file, findStructs)

		}
	}
}

func findStructs(n ast.Node) bool {
	switch x := n.(type) {
	case *ast.TypeSpec:

		switch t := x.Type.(type) {
		case *ast.Ident: // typedef
			if t.Obj == nil {
				typedefs_basic[x.Name.Name] = t.Name
			}
			log.Printf("\t\t%#v", t)
		case *ast.StructType:
			currentStruct = x.Name.Name
			log.Printf("\t\t%v", x.Name.Name)
		case *ast.InterfaceType:
			interfaces[x.Name.Name] = true
		}
	case *ast.StructType:
		getherFields(x)
	}

	return true
}

func addField(f FieldCloner) {
	if _, ok := structs[currentStruct]; !ok {
		structs[currentStruct] = make([]FieldCloner, 0)
		structs[currentStruct] = append(structs[currentStruct], f)
	} else {
		structs[currentStruct] = append(structs[currentStruct], f)
	}
}

type FieldCloner interface {
	GenerateStr() string
}

type ClonerBase struct {
	fieldName string
	type_     string
	isList    bool
	isMap     bool
	isPtr     bool
}

type BasicCloner struct {
	ClonerBase
}

func (c *BasicCloner) GenerateStr() string {
	if c.isList {
		if c.isPtr {
			return fmt.Sprintf("%v: deepcopy.CloneListPtr_%v(obj.%v),", c.fieldName, c.type_, c.fieldName)
		} else {
			return fmt.Sprintf("%v: deepcopy.CloneList_%v(obj.%v),", c.fieldName, c.type_, c.fieldName)
		}
	}
	// TODO: map
	if c.isPtr {
		return fmt.Sprintf("%v: deepcopy.Clone_%v(obj.%v),", c.fieldName, c.type_, c.fieldName)
	} else {
		return fmt.Sprintf("%v: obj.%v,", c.fieldName, c.fieldName)
	}
}

type StructCloner struct {
	ClonerBase
}

type ShallowCloner struct {
	BasicCloner
}

func (c *ShallowCloner) GenerateStr() string {
	return c.BasicCloner.GenerateStr() + "// Shallow Copy"
}

func (c *StructCloner) GenerateStr() string {
	if c.isList {
		if c.isPtr {
			return fmt.Sprintf("%v: CloneListPtr_%v(obj.%v),", c.fieldName, c.type_, c.fieldName)
		} else {
			return fmt.Sprintf("%v: CloneList_%v(obj.%v),", c.fieldName, c.type_, c.fieldName)
		}
	}
	// TODO: when is map
	if c.isPtr {
		return fmt.Sprintf("%v: obj.%v.Clone(),", c.fieldName, c.fieldName)
	} else {
		return fmt.Sprintf("%v: *obj.%v.Clone(),", c.fieldName, c.fieldName)
	}
}

type InterfaceCloner struct {
	ClonerBase
	interfaceName string
}

func (c *InterfaceCloner) GenerateStr() string {
	if c.isList {
		return fmt.Sprintf("%v: CloneList_%v(obj.%v),", c.fieldName, c.interfaceName, c.fieldName)
	}
	return fmt.Sprintf("%v: Clone_%v(obj.%v),", c.fieldName, c.interfaceName, c.fieldName)
}

func sepLine() {
	log.Println("--------------------------")
}

func generateBasic(type_ string, field *ast.Field, isPtr, isList bool) {
	for _, fieldName := range field.Names {
		s := &BasicCloner{ClonerBase{type_: type_, fieldName: fieldName.Name, isList: isList, isPtr: isPtr}}
		addField(s)
		fmt.Println("\t\t" + s.GenerateStr())
	}
}

func generateAnonymous(obj *ast.Object, isPtr, isList bool) {
	s := &StructCloner{ClonerBase: ClonerBase{fieldName: obj.Name, isList: isList, isPtr: isPtr}}
	addField(s)
	fmt.Println("\t\t" + s.GenerateStr())
}

func generateStruct(type_ string, field *ast.Field, isPtr, isList bool) {
	for _, fieldName := range field.Names {
		s := &StructCloner{ClonerBase: ClonerBase{type_: type_, fieldName: fieldName.Name, isList: isList, isPtr: isPtr}}
		addField(s)
		fmt.Println("\t\t" + s.GenerateStr())
	}
}

func generateTypedef(type_ string, field *ast.Field, isPtr, isList bool) {
	for _, fieldName := range field.Names {
		s := &StructCloner{ClonerBase: ClonerBase{type_: type_, fieldName: fieldName.Name, isPtr: isPtr, isList: isList}}
		addField(s)
		fmt.Println("\t\t" + s.GenerateStr())
	}
}

func generateInterface(e *ast.Ident, field *ast.Field) {
	for _, fieldName := range field.Names {
		s := &InterfaceCloner{ClonerBase: ClonerBase{fieldName: fieldName.Name}, interfaceName: e.Name}
		addField(s)
		fmt.Println("\t\t" + s.GenerateStr())
	}
}

func generateFromIdent(e *ast.Ident, field *ast.Field, isPtr, isList bool) {
	if obj := e.Obj; obj == nil { // basicType
		generateBasic(e.Name, field, isPtr, isList)
	} else { // structure, typedef, anonymous field, interface,
		if len(field.Names) == 0 { // anonymous field
			decl, ok := obj.Decl.(*ast.TypeSpec)
			if ok {
				switch decl.Type.(type) {
				case *ast.StructType:
					generateAnonymous(obj, isPtr, isList)
				}
			} else {
				log.Printf("\t\tERROR, cannot cast to decl")
			}

		} else { // structure, typedef, interface
			decl, ok := obj.Decl.(*ast.TypeSpec)
			if ok {
				switch decl.Type.(type) {
				case *ast.StructType: // structure
					generateStruct(decl.Name.Name, field, isPtr, isList)
				case *ast.Ident: // typedef
					generateTypedef(decl.Name.Name, field, isPtr, isList)
				case *ast.InterfaceType: // interface
					generateInterface(e, field)
				default:
					log.Printf("\t\tERROR, unknown field type: %#v", decl)
				}
			} else {
				log.Printf("\t\tERROR, cannot cast to decl")
			}

		}
	}
}

func generateFromStar(e *ast.StarExpr, field *ast.Field, isList bool) {
	switch x := e.X.(type) {
	case *ast.Ident:
		generateFromIdent(x, field, true, isList)
	default:
		log.Printf("\t\tnot implemented yet *ast.StarExpr: %#v", x)
	}
}

func generateFromSelector(e *ast.SelectorExpr, field *ast.Field) {
	// Temporary solution
	s := &ShallowCloner{BasicCloner: BasicCloner{ClonerBase: ClonerBase{fieldName: e.Sel.Name}}}
	addField(s)
	fmt.Println("\t\t" + s.GenerateStr())
}

func generateFromArray(e *ast.ArrayType, field *ast.Field) {
	switch x := e.Elt.(type) {
	case *ast.Ident:
		generateFromIdent(x, field, false, true)
	case *ast.StarExpr:
		generateFromStar(x, field, true)
	default:
		log.Printf("\t\tnot implemented yet *ast.ArrayField: %#v", x)
	}
}

func generateLine(field *ast.Field) {
	switch x := field.Type.(type) {
	case *ast.Ident:
		generateFromIdent(x, field, false, false)
		sepLine()
	case *ast.StarExpr:
		generateFromStar(x, field, false)
		sepLine()
	case *ast.SelectorExpr:
		generateFromSelector(x, field)
		sepLine()
	case *ast.ArrayType:
		generateFromArray(x, field)
		sepLine()
	case *ast.ChanType:
		log.Printf("\tnot implemented yet:%#v", x)
		sepLine()
	case *ast.MapType:
		log.Printf("\tnot implemented yet:%#v", x)
		sepLine()
	default:
		log.Println("\tnot implemented yet:", reflect.TypeOf(x))
	}
}

func getherFields(n *ast.StructType) {
	for _, field := range n.Fields.List {
		generateLine(field)
	}
}

func main() {
	fset := token.NewFileSet()
	pkgMap, err := parser.ParseDir(fset, *path, filter, parser.AllErrors)
	if err != nil {
		log.Println(err)
		return
	}

	run(pkgMap)

	for _, pkg := range pkgMap {
		for _, file := range pkg.Files {

			outfile, err := os.Create(*path + "/" + file.Name.Name + "_clone.go")
			if err != nil {
				log.Println(err)
			}

			outfile.WriteString("package " + pkg.Name + "\n\n")
			outfile.WriteString("import \"github.com/wookesh/deepcopy/deepcopy\"\n\n")

			for interface_ := range interfaces {
				generateInterfaceCloneCode(outfile, interface_)
			}

			for class, fieldCloners := range structs {
				generateCloneCode(outfile, class, fieldCloners)
			}

			for typedef, class := range typedefs {
				generateCloneCode(outfile, typedef, structs[class])
			}

			for typedef, _ := range typedefs_basic {
				generateTypedefBasicCode(outfile, typedef)
			}
		}

	}

}

func generateInterfaceCloneCode(outfile *os.File, interf string) {
	outfile.WriteString("func Clone_" + interf + "(obj " + interf + ") " + interf + " {\n")
	outfile.WriteString("\tx, ok := interface{}(obj).(deepcopy.DeepCopier)\n")
	outfile.WriteString("\tif ok {\n")
	outfile.WriteString("\t\ti, ok := x.CloneInterface().(" + interf + ")\n")
	outfile.WriteString("\t\tif ok {\n")
	outfile.WriteString("\t\t\treturn i\n")
	outfile.WriteString("\t\t}\n")
	outfile.WriteString("\t\treturn obj\n")
	outfile.WriteString("\t}\n")
	outfile.WriteString("\treturn obj\n")
	outfile.WriteString("}\n\n")
}

func generateCloneCode(outfile *os.File, class string, fields []FieldCloner) {
	// SingleGenerator
	outfile.WriteString("func (obj *" + class + ") Clone() *" + class + " {\n")
	outfile.WriteString("\tif obj == nil {\n")
	outfile.WriteString("\t\treturn nil\n")
	outfile.WriteString("\t}\n")
	outfile.WriteString("\treturn &" + class + "{\n")
	for _, field := range fields {
		outfile.WriteString("\t\t" + field.GenerateStr() + "\n")
	}
	outfile.WriteString("\t}\n")
	outfile.WriteString("}\n\n")

	// InterfaceCloneGenerator
	outfile.WriteString("func (obj *" + class + ") CloneInterface() deepcopy.DeepCopier {\n")
	outfile.WriteString("\treturn obj.Clone()\n")
	outfile.WriteString("}\n\n")

	// ListGenerator
	generateListCloner(outfile, class, false)

	// ListPtrGenerator
	generateListCloner(outfile, class, true)
}

func generateListCloner(outfile *os.File, class string, isPtr bool) {
	ptr := "*"
	name := "CloneList"
	typ := class
	if isPtr {
		typ = "*" + class
		ptr = ""
		name = name + "Ptr"
	}
	outfile.WriteString("func " + name + "_" + class + "(objs []" + typ + ") []" + typ + " {\n")
	outfile.WriteString("\tretList := make([]" + typ + ", 0, len(objs))\n")
	outfile.WriteString("\tfor _, obj := range objs {\n")
	outfile.WriteString("\t\tretList = append(retList, " + ptr + "obj.Clone())\n")
	outfile.WriteString("\t}\n")
	outfile.WriteString("\treturn retList\n")
	outfile.WriteString("}\n\n")
}

func generateTypedefBasicCode(outfile *os.File, class string) {
	// SingleGenerator
	outfile.WriteString("func (obj *" + class + ") Clone() *" + class + " {\n")
	outfile.WriteString("\tif obj == nil {\n")
	outfile.WriteString("\t\treturn nil\n")
	outfile.WriteString("\t}\n")
	outfile.WriteString("\tret := " + class + "(*obj)\n")
	outfile.WriteString("\treturn &ret\n")
	outfile.WriteString("}\n\n")

	// ListGenerator
	generateListCloner(outfile, class, false)

	// ListPtrGenerator
	generateListCloner(outfile, class, true)
}

func add(class, field string) {
	l, ok := data[class]
	if !ok {
		l = make([]string, 0)
	}
	data[class] = append(l, field)
}
