package graph

import "github.com/smallretardedfish/gql-federation/patient/storage"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PatientStore storage.PatientStore
}
