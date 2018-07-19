package uc

// Handler handles all your use cases
type Handler interface {
	Health() error
}

// Interactor is the struct that will have as properties all the IMPLEMENTED
// interfaces in order to provide them to its methods: the use cases, so we can
// actually implement the Handler interface
type Interactor struct{}

// NewInteractor : the Interactor constructor, use this in order to avoid nil
// pointers at runtime
func NewInteractor() Interactor {
	return Interactor{}
}

// Put your interfaces here and insert them in the Interactor struct (as well as
// in the constructor)
