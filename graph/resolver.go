package graph

import "github.com/derivedpuma7/go-graphQL/internal/database"

type Resolver struct{
	CategoryDB *database.Category
	CourseDB *database.Course
}
