package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/KenxinKun/gqlgenplayground/graph/model"
)

// Author is the resolver for the author field.
func (r *bookResolver) Author(ctx context.Context, obj *model.Book) (*model.Author, error) {
	return &model.Author{
		Name: "author name from author resolver for book: " + obj.Title,
	}, nil
}

// Fractal is the resolver for the fractal field.
func (r *fractalResolver) Fractal(ctx context.Context, obj *model.Fractal) (*model.Fractal, error) {
	return &model.Fractal{
		Level: obj.Level + 1,
	}, nil
}

// Books is the resolver for the books field.
func (r *libraryResolver) Books(ctx context.Context, obj *model.Library) ([]*model.Book, error) {
	return []*model.Book{
		{
			Title: "book title from books resolver",
		},
	}, nil
}

// Libraries is the resolver for the libraries field.
func (r *queryResolver) Libraries(ctx context.Context) ([]*model.Library, error) {
	return []*model.Library{
		{
			Address: "library address from libraries resolver",
		},
	}, nil
}

// Fractal is the resolver for the fractal field.
func (r *queryResolver) Fractal(ctx context.Context) (*model.Fractal, error) {
	return &model.Fractal{
		Level: 0,
	}, nil
}

// Book returns BookResolver implementation.
func (r *Resolver) Book() BookResolver { return &bookResolver{r} }

// Fractal returns FractalResolver implementation.
func (r *Resolver) Fractal() FractalResolver { return &fractalResolver{r} }

// Library returns LibraryResolver implementation.
func (r *Resolver) Library() LibraryResolver { return &libraryResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type bookResolver struct{ *Resolver }
type fractalResolver struct{ *Resolver }
type libraryResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }