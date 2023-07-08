package usecase

import (
	"github.com/SkadisCode/meaning/entity"
	"github.com/SkadisCode/meaning/repository"
)

type DictionaryUsecase struct {
	dictionaryRepo repository.DictionaryRepository
}

func NewDictionaryUsecase(dictionaryRepo repository.DictionaryRepository) *DictionaryUsecase {
	return &DictionaryUsecase{dictionaryRepo: dictionaryRepo}
}

func (uc *DictionaryUsecase) SaveDictionary(dictionary *entity.Dictionary) error {
	// Implement the business logic for saving a dictionary entry
	// You can perform any validation or additional operations here
	err := uc.dictionaryRepo.Save(dictionary)
	if err != nil {
		return err
	}
	return nil
}

func (uc *DictionaryUsecase) GetDictionary(word string) (*entity.Dictionary, error) {
	// Implement the business logic for retrieving a dictionary entry
	dictionary, err := uc.dictionaryRepo.Find(word)
	if err != nil {
		return nil, err
	}
	return dictionary, nil
}
