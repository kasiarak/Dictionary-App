package graphql

import (
	"dictionary-app/db"
	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addWordTranslation": &graphql.Field{
				Type: TranslationType,
				Args: graphql.FieldConfigArgument{
					"word":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"translation": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"sentence":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db.Mutex.Lock()
					defer db.Mutex.Unlock()

					wordText := p.Args["word"].(string)
					translationText := p.Args["translation"].(string)
					sentenceText := p.Args["sentence"].(string)

					var word db.Word
					db.DB.FirstOrCreate(&word, db.Word{Word: wordText})

					translation := db.Translation{WordID: word.ID, Translation: translationText}
					db.DB.Create(&translation)

					sentence := db.Sentence{TranslationID: translation.ID, Sentence: sentenceText}
					db.DB.Create(&sentence)

					translation.Sentences = []db.Sentence{sentence}
					return translation, nil
				},
			},
		},
	},
)
