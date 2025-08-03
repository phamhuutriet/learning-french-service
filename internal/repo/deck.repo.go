package repo

import (
	"context"
	"fmt"
	"learning-french-service/global"
	"learning-french-service/internal/ent"
	"learning-french-service/internal/ent/deck"
	"learning-french-service/internal/ent/user"
	"learning-french-service/internal/types"
)

type IDeckRepository interface {
	GetDecks(userID int, category, level string, active *bool) ([]types.Deck, int, error)
	CreateDeck(userID int, name, description, category, targetLevel string) (*types.Deck, error)
	GetDeck(userID, deckID int) (*types.Deck, error)
	UpdateDeck(userID, deckID int, name, description, category, targetLevel string) (*types.Deck, error)
	DeleteDeck(userID, deckID int) error
}

type DeckRepository struct{}

func NewDeckRepository() IDeckRepository {
	return &DeckRepository{}
}

// ---------- Helper ----------

func entDeckToTypeDeck(d *ent.Deck) *types.Deck {
	return &types.Deck{
		ID:          fmt.Sprintf("%d", d.ID),
		Name:        d.Name,
		Description: d.Description,
		Category:    d.Category,
		TargetLevel: d.TargetLevel,
		WordCount:   d.WordCount,
		IsActive:    d.IsActive,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

// ---------- Methods ----------

func (dr *DeckRepository) GetDecks(userID int, category, level string, active *bool) ([]types.Deck, int, error) {
	ctx := context.Background()

	query := global.EntClient.Deck.Query().
		Where(deck.HasUserWith(user.IDEQ(userID)))

	if category != "" {
		query = query.Where(deck.CategoryEQ(category))
	}
	if level != "" {
		query = query.Where(deck.TargetLevelEQ(level))
	}
	if active != nil {
		query = query.Where(deck.IsActiveEQ(*active))
	}

	decks, err := query.All(ctx)
	if err != nil {
		return nil, 0, err
	}

	result := make([]types.Deck, len(decks))
	for i, d := range decks {
		result[i] = *entDeckToTypeDeck(d)
	}

	return result, len(result), nil
}

func (dr *DeckRepository) CreateDeck(userID int, name, description, category, targetLevel string) (*types.Deck, error) {
	ctx := context.Background()

	deck, err := global.EntClient.Deck.Create().
		SetName(name).
		SetDescription(description).
		SetCategory(category).
		SetTargetLevel(targetLevel).
		SetIsActive(true).
		SetWordCount(0).
		SetUserID(userID). // Make sure this matches your schema
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return entDeckToTypeDeck(deck), nil
}

func (dr *DeckRepository) GetDeck(userID, deckID int) (*types.Deck, error) {
	ctx := context.Background()

	deck, err := global.EntClient.Deck.Query().
		Where(
			deck.IDEQ(deckID),
			deck.HasUserWith(user.IDEQ(userID)),
		).Only(ctx)
	if err != nil {
		return nil, err
	}

	return entDeckToTypeDeck(deck), nil
}

func (dr *DeckRepository) UpdateDeck(userID, deckID int, name, description, category, targetLevel string) (*types.Deck, error) {
	ctx := context.Background()

	// Ensure deck belongs to user
	deck, err := global.EntClient.Deck.Query().
		Where(
			deck.IDEQ(deckID),
			deck.HasUserWith(user.IDEQ(userID)),
		).Only(ctx)
	if err != nil {
		return nil, err
	}

	update := deck.Update()
	if name != "" {
		update.SetName(name)
	}
	if description != "" {
		update.SetDescription(description)
	}
	if category != "" {
		update.SetCategory(category)
	}
	if targetLevel != "" {
		update.SetTargetLevel(targetLevel)
	}

	updatedDeck, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}

	return entDeckToTypeDeck(updatedDeck), nil
}

func (dr *DeckRepository) DeleteDeck(userID, deckID int) error {
	ctx := context.Background()

	// Ensure deck belongs to user
	deck, err := global.EntClient.Deck.Query().
		Where(
			deck.IDEQ(deckID),
			deck.HasUserWith(user.IDEQ(userID)),
		).Only(ctx)
	if err != nil {
		return err
	}

	return global.EntClient.Deck.DeleteOne(deck).Exec(ctx)
}
