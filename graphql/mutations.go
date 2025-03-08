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
					if err := db.DB.FirstOrCreate(&word, db.Word{Word: wordText}).Error; err != nil {
						return nil, err
					}

					translation := db.Translation{WordID: word.ID, Translation: translationText}
					if err := db.DB.Create(&translation).Error; err != nil {
						return nil, err
					}

					sentence := db.Sentence{TranslationID: translation.ID, Sentence: sentenceText}
					if err := db.DB.Create(&sentence).Error; err != nil {
						return nil, err
					}

					translation.Sentences = []db.Sentence{sentence}
					return translation, nil
				},
			},
		},
	},
)
