package graph

import "github.com/smallretardedfish/gql-federation/storage"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoStore storage.TodoStore
}
