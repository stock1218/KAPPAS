package model

// PAN is a Key type that encapsulates PAN data.
//
type PAN struct{}

// NewPAN returns a pointer to a new PAN struct.
//
func NewPAN() *PAN {
	return new(PAN)
}
