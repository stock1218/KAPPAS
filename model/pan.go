package model

// PAN is a Key type that encapsulates PAN data.
//
type PAN struct{}

// NewPAN returns a pointer to a new PAN struct.
//
func NewPAN() *PAN {
	return new(PAN)
}

// GetPayload will return the a string representation of the data stored in this Data.
//
func (pan PAN) GetPayload() string {
	return "payload"
}
