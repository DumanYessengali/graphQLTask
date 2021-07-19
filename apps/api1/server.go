package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"twoBinPJ/adapters"
	graph1 "twoBinPJ/apps/api1/graph"
	generated1 "twoBinPJ/apps/api1/graph/generated"
	"twoBinPJ/domains/auth"
	"twoBinPJ/domains/project"
	"twoBinPJ/domains/report"
	"twoBinPJ/domains/user"
	"twoBinPJ/domains/vulnerability"
	"twoBinPJ/middleware"
)

func main() {
	config := adapters.ParseConfig()
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Webport1
	}

	router := chi.NewRouter()
	db, err := adapters.InitDB(adapters.Config{
		Host:     config.DataBaseHost,
		Port:     config.DataBasePort,
		Username: config.DataBaseUsername,
		Password: config.DataBasePassword,
		DBName:   config.DataBaseDbname,
		SSLMode:  config.DataBaseSslmode,
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
		return
	}

	authModule := auth.NewAuthModule(db)
	userModule := user.NewUserModule(db)
	vulnerabilityModule := vulnerability.NewVulnerabilityModule(db)
	projectModule := project.NewProjectModule(db)
	reportModule := report.NewReportModule(db)

	router.Use(middleware.Middleware())
	srv := handler.NewDefaultServer(generated1.NewExecutableSchema(generated1.Config{Resolvers: &graph1.Resolver{
		AuthModule:          authModule,
		UserModule:          userModule,
		VulnerabilityModule: vulnerabilityModule,
		ProjectModule:       projectModule,
		ReportModule:        *reportModule,
	}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
