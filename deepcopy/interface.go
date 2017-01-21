package deepcopy

type DeepCopier interface {
	CloneInterface() DeepCopier
}
