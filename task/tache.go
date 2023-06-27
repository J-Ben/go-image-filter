package task

type Filter interface {
	Process(cheminSource, cheminDest string) error
}

type Tasker interface {
	Process() error
}
