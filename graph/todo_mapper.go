package graph

import (
	"github.com/smallretardedfish/gql-federation/graph/model"
	"github.com/smallretardedfish/gql-federation/storage"
)

func TodoToGraph(todo storage.Todo) model.Todo {
	var user *model.User
	if todo.User != nil {
		res := UserToGraph(*todo.User)
		user = &res
	}

	return model.Todo{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
		User: user,
	}
}

func UserToGraph(u storage.User) model.User {
	return model.User{
		ID:   u.ID,
		Name: u.Name,
	}
}
