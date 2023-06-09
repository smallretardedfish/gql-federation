package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"github.com/smallretardedfish/gql-federation/book-service/graph/model"
)

// Books is the resolver for the books field.
func (r *authorResolver) Books(ctx context.Context, obj *model.Author) ([]*model.Book, error) {
	var res []*model.Book
	books, err := r.BookStorage.GetAll()
	if err != nil {
		return nil, err
	}
	for i := range books {
		if books[i].Author.ID == obj.ID {
			res = append(res, books[i])
		}
	}

	return res, nil
}

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input *model.NewBook) (*model.Book, error) {
	b, err := r.BookStorage.Create(model.Book{
		Title:  input.Title,
		Year:   input.Year,
		Author: &model.Author{ID: input.AuthorID},
	})
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return r.BookStorage.GetAll()
}

// Author returns AuthorResolver implementation.
func (r *Resolver) Author() AuthorResolver { return &authorResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type authorResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
