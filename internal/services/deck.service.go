package services

import (
	"learning-french-service/internal/repo"
	"learning-french-service/internal/types"
)

type IDeckService interface {
	GetDecks(userID int, category, level string, active *bool) ([]types.Deck, int, error)
	CreateDeck(userID int, name, description, category, targetLevel string) (*types.Deck, error)
	GetDeck(userID int, deckID int) (*types.Deck, error)
	UpdateDeck(userID int, deckID int, name, description, category, targetLevel string) (*types.Deck, error)
	DeleteDeck(userID int, deckID int) error
}

type DeckService struct {
	deckRepo repo.IDeckRepository
}

func NewDeckService(deckRepo repo.IDeckRepository) IDeckService {
	return &DeckService{
		deckRepo: deckRepo,
	}
}

func (ds *DeckService) GetDecks(userID int, category, level string, active *bool) ([]types.Deck, int, error) {
	return ds.deckRepo.GetDecks(userID, category, level, active)
}

func (ds *DeckService) CreateDeck(userID int, name, description, category, targetLevel string) (*types.Deck, error) {
	return ds.deckRepo.CreateDeck(userID, name, description, category, targetLevel)
}

func (ds *DeckService) GetDeck(userID int, deckID int) (*types.Deck, error) {
	return ds.deckRepo.GetDeck(userID, deckID)
}

func (ds *DeckService) UpdateDeck(userID int, deckID int, name, description, category, targetLevel string) (*types.Deck, error) {
	return ds.deckRepo.UpdateDeck(userID, deckID, name, description, category, targetLevel)
}

func (ds *DeckService) DeleteDeck(userID int, deckID int) error {
	return ds.deckRepo.DeleteDeck(userID, deckID)
}
