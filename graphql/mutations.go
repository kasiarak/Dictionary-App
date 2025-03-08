package graphql

import (
	"dictionary-app/db"
	"fmt"

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

					if len(wordText) > 50 {
						return nil, fmt.Errorf("word is too long, maximum length is 50 characters")
					}
					if len(translationText) > 50 {
						return nil, fmt.Errorf("translation is too long, maximum length is 50 characters")
					}
					if len(sentenceText) > 100 {
						return nil, fmt.Errorf("sentence is too long, maximum length is 100 characters")
					}

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
