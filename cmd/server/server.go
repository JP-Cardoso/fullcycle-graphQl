package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/JP-Cardoso/13-graphQL/graph"
	"github.com/JP-Cardoso/13-graphQL/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	// Abrindo a conexão com o banco de dados
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	// Fechando a comunicação com o banco de dados
	defer db.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	//Esse cara que faz a manipulação no banco
	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)


	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDb,
		CourseDB: courseDb,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
