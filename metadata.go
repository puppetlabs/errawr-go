package errawr

type HTTPMetadata interface {
	Status() int
}

type Metadata interface {
	HTTP() (HTTPMetadata, bool)
}
