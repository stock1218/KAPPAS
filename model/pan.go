package model

// PAN is a Key type that encapsulates PAN data.
//
type PAN struct {
	payload string
}

// NewPAN returns a pointer to a new PAN struct.
//
func NewPAN() *PAN {
	return new(PAN)
}

// GetPayload will return the a string representation of the data stored in this Data.
//
func (pan PAN) GetPayload() string {
	return pan.payload
}

// SetPayload will set the payload of the current Data.
//
func (pan *PAN) SetPayload(newPayload string) {
	pan.payload = newPayload
}
