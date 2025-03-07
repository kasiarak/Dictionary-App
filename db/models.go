package db

type Word struct {
	ID   int    `json:"id"`
	Word string `json:"word" gorm:"uniqueIndex"`
}

type Translation struct {
	ID          int        `json:"id"`
	WordID      int        `json:"word_id"`
	Translation string     `json:"translation"`
	Sentences   []Sentence `json:"sentences" gorm:"foreignKey:TranslationID"`
}

type Sentence struct {
	ID            int    `json:"id"`
	TranslationID int    `json:"translation_id"`
	Sentence      string `json:"sentence"`
}
