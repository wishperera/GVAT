package container

type Module interface {
	Init(c Container) error
}
