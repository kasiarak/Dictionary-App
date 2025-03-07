package graphql

import (
	"github.com/graphql-go/graphql"
)

var WordType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Word",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"word": &graphql.Field{Type: graphql.String},
		},
	},
)

var SentenceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Sentence",
		Fields: graphql.Fields{
			"id":             &graphql.Field{Type: graphql.Int},
			"translation_id": &graphql.Field{Type: graphql.Int},
			"sentence":       &graphql.Field{Type: graphql.String},
		},
	},
)

var TranslationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Translation",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"word_id":     &graphql.Field{Type: graphql.Int},
			"translation": &graphql.Field{Type: graphql.String},
			"sentences":   &graphql.Field{Type: graphql.NewList(SentenceType)},
		},
	},
)
