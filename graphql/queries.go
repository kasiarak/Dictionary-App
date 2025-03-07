package graphql

import (
	"dictionary-app/db"
	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getTranslation": &graphql.Field{
				Type: graphql.NewList(TranslationType),
				Args: graphql.FieldConfigArgument{
					"word": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					word := p.Args["word"].(string)
					var translations []db.Translation
					db.DB.Preload("Sentences").Joins("JOIN words ON words.id = translations.word_id").Where("words.word = ?", word).Find(&translations)
					return translations, nil
				},
			},
		},
	},
)
