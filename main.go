package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/ymsg19/sfc-guardbot/ent"
	"github.com/ymsg19/sfc-guardbot/ent/migrate"
	"github.com/ymsg19/sfc-guardbot/graph"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

const (
	ENV_DEVELOPMENT = "development"
	ENV_PRODUCTION  = "production"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = ENV_DEVELOPMENT
	}

	if env == ENV_DEVELOPMENT {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	client, err := func() (*ent.Client, error) {
		if env == ENV_DEVELOPMENT {
			return ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
		} else {
			database := os.Getenv("DATABASE_URL")
			if database == "" {
				return nil, errors.New("DATABASE_URL is not set")
			}
			return ent.Open(dialect.Postgres, os.Getenv("DATABASE_URL"))
		}
	}()
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	defer client.Close()

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	srv := handler.NewDefaultServer(graph.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
