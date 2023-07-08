package repository

import (
	"database/sql"
	"log"

	"github.com/SkadisCode/meaning/entity"
)

type DictionaryRepository struct {
	db *sql.DB
}

func NewDictionaryRepository(db *sql.DB) *DictionaryRepository {
	return &DictionaryRepository{db: db}
}

func (r *DictionaryRepository) Save(dictionary *entity.Dictionary) error {
	// Implement the logic to save the dictionary entry to the database
	_, err := r.db.Exec("INSERT INTO dictionary (word, entry, type, tier, semantic) VALUES ($1, $2, $3, $4, $5)",
		dictionary.Word, dictionary.Entry, dictionary.Type, dictionary.Tier, dictionary.Semantic)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *DictionaryRepository) Find(word string) (*entity.Dictionary, error) {
	// Implement the logic to find a dictionary entry by word in the database
	row := r.db.QueryRow("SELECT word, entry, type, tier, semantic FROM dictionary WHERE word = $1", word)
	dictionary := &entity.Dictionary{}
	err := row.Scan(&dictionary.Word, &dictionary.Entry, &dictionary.Type, &dictionary.Tier, &dictionary.Semantic)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if the entry is not found
		}
		log.Println(err)
		return nil, err
	}
	return dictionary, nil
}
