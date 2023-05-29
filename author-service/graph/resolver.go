package graph

import "github.com/smallretardedfish/gql-federation/author-service/author_storage"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthorStorage *author_storage.InMemoryAuthorStorage
}
