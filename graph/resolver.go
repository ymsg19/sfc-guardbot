package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/bwmarrin/discordgo"
	"github.com/ymsg19/sfc-guardbot/ent"
	"github.com/ymsg19/sfc-guardbot/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client  *ent.Client
	session *discordgo.Session
}

func NewSchema(client *ent.Client, session *discordgo.Session) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &Resolver{
				client,
				session,
			},
		},
	)
}
