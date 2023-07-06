package graph

import "github.com/JP-Cardoso/13-graphQL/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Fazendo uma injeção de depêndencia
type Resolver struct{
	CategoryDB *database.Category
	CourseDB *database.Course 
}
