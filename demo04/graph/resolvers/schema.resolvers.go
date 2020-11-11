package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"lai.com/gqlgen_study/demo04/graph/generated"
	"lai.com/gqlgen_study/demo04/graph/model"
)

func (r *queryResolver) Persons(ctx context.Context) ([]*model.Person, error) {
	name := "小灰灰"
	age := 19
	p := &model.Person{
		Name: &name,
		Age:  &age,
	}
	var persons []*model.Person
	persons = append(persons, p)
	return persons, nil
}

func (r *queryResolver) Student(ctx context.Context, id string) (*model.Student, error) {
	name := "小灰灰"
	age := 19
	class := "初三5班"
	stu := &model.Student{
		Name:  &name,
		Age:   &age,
		Class: &class,
	}
	return stu, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
