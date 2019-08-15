package model

import "github.com/stock1218/KAPPAS/model/data"

// Model is an interface that defines the functions for Model types
//
type Model interface {
	PutPAN(*data.PAN) error           // Put a Data value into the AWS model, and update the Data ID
	GetPAN(string) (data.Data, error) // Get a Data value form the AWS model
}
