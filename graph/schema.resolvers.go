package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"
	"github.com/ArianNexux/go-graphql/graph/model"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category, err := r.CategoryDB.Create(input.Name, *input.Description)

	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID: category.ID,
		Name: category.Name,
		Description: &category.Description,
	}, nil
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	course, err := r.CourseDB.Create(input.Name, *input.Description, input.CategoryID)

	if err != nil {
		return nil, err
	}

	return &model.Course{
		ID: course.ID,
		Name: course.Name,
		Description: &course.Description,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categories, err := r.CategoryDB.FindAll()

	if err != nil {
		return nil, err
	}

	var categoriesModel []*model.Category

	for _, c := range categories {

		categoriesModel = append(categoriesModel, &model.Category{
			ID: c.ID,
			Name: c.Name,
			Description: &c.Description,
		})

	}

	return categoriesModel, nil
}

// Course is the resolver for the course field.
func (r *queryResolver) Course(ctx context.Context) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindAll()

	if err != nil {
		return nil, err
	}

	var coursesModel []*model.Course

	for _, c := range courses {

		coursesModel = append(coursesModel, &model.Course{
			ID: c.ID,
			Name: c.Name,
			Description: &c.Description,
		})

	}

	return coursesModel, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
