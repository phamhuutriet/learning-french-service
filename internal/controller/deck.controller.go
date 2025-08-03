package controller

import (
	"learning-french-service/internal/services"
	"learning-french-service/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeckController struct {
	deckService services.IDeckService
}

func NewDeckController(deckService services.IDeckService) *DeckController {
	return &DeckController{
		deckService: deckService,
	}
}

// GetDecks handles GET /decks
func (dc *DeckController) GetDecks(c *gin.Context) {
	// TODO: Get user ID from auth middleware
	userID := 1 // This should come from auth middleware

	category := c.Query("category")
	level := c.Query("level")
	activeStr := c.Query("active")

	var active *bool
	if activeStr != "" {
		activeVal, err := strconv.ParseBool(activeStr)
		if err == nil {
			active = &activeVal
		}
	}

	decks, total, err := dc.deckService.GetDecks(userID, category, level, active)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError)
		return
	}

	response.SuccessResponse(c, http.StatusOK, gin.H{
		"decks": decks,
		"total": total,
	})
}

// CreateDeck handles POST /decks
func (dc *DeckController) CreateDeck(c *gin.Context) {
	// TODO: Get user ID from auth middleware
	userID := 1 // This should come from auth middleware

	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Category    string `json:"category"`
		TargetLevel string `json:"target_level"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest)
		return
	}

	deck, err := dc.deckService.CreateDeck(userID, req.Name, req.Description, req.Category, req.TargetLevel)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError)
		return
	}

	response.SuccessResponse(c, http.StatusCreated, deck)
}

// GetDeck handles GET /decks/:id
func (dc *DeckController) GetDeck(c *gin.Context) {
	// TODO: Get user ID from auth middleware
	userID := 1 // This should come from auth middleware

	deckID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest)
		return
	}

	deck, err := dc.deckService.GetDeck(userID, deckID)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound)
		return
	}

	response.SuccessResponse(c, http.StatusOK, deck)
}

// UpdateDeck handles PUT /decks/:id
func (dc *DeckController) UpdateDeck(c *gin.Context) {
	// TODO: Get user ID from auth middleware
	userID := 1 // This should come from auth middleware

	deckID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest)
		return
	}

	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Category    string `json:"category"`
		TargetLevel string `json:"target_level"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest)
		return
	}

	deck, err := dc.deckService.UpdateDeck(userID, deckID, req.Name, req.Description, req.Category, req.TargetLevel)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound)
		return
	}

	response.SuccessResponse(c, http.StatusOK, deck)
}

// DeleteDeck handles DELETE /decks/:id
func (dc *DeckController) DeleteDeck(c *gin.Context) {
	// TODO: Get user ID from auth middleware
	userID := 1 // This should come from auth middleware

	deckID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest)
		return
	}

	err = dc.deckService.DeleteDeck(userID, deckID)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound)
		return
	}

	response.SuccessResponse(c, http.StatusNoContent, nil)
}
