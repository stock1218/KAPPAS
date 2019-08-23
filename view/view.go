package view

// View is an interface that defines the functions for Views
type View interface {
	GetObservers() []func(string) // Get the observer functions registered to the view
}
