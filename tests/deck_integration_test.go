package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	baseURL = "http://localhost:8080"
	apiURL  = baseURL + "/api/v1/decks"
)

// ---------- Models ----------

type Deck struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	TargetLevel string    `json:"target_level"`
	WordCount   int       `json:"word_count"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type APIResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type DecksResponse struct {
	Decks []Deck `json:"decks"`
	Total int    `json:"total"`
}

// ---------- Utility ----------

func skipIfServerDown(t *testing.T) {
	resp, err := http.Get(baseURL + "/api/v1/health")
	if err != nil {
		t.Skip("Server not running, skipping integration test")
	}
	resp.Body.Close()
}

func makeRequest(method, url string, body interface{}) (*http.Response, error) {
	var reqBody []byte
	if body != nil {
		reqBody, _ = json.Marshal(body)
	}
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return http.DefaultClient.Do(req)
}

func parseResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return json.Unmarshal(body, target)
}

// ---------- Reusable Operations ----------

func createDeck(t *testing.T, name, level string) Deck {
	deckData := map[string]interface{}{
		"name":         name,
		"description":  "Test deck",
		"category":     "test",
		"target_level": level,
	}

	resp, err := makeRequest("POST", apiURL, deckData)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var apiResp APIResponse
	_ = parseResponse(resp, &apiResp)
	assert.Equal(t, 201, apiResp.Code)

	var deck Deck
	_ = json.Unmarshal(apiResp.Data, &deck)
	return deck
}

func getDeckByID(t *testing.T, id string) (Deck, int) {
	resp, err := makeRequest("GET", fmt.Sprintf("%s/%s", apiURL, id), nil)
	assert.NoError(t, err)

	var apiResp APIResponse
	_ = parseResponse(resp, &apiResp)

	var deck Deck
	_ = json.Unmarshal(apiResp.Data, &deck)
	return deck, apiResp.Code
}

func getAllDecks(t *testing.T) ([]Deck, int) {
	resp, err := makeRequest("GET", apiURL, nil)
	assert.NoError(t, err)

	var apiResp APIResponse
	_ = parseResponse(resp, &apiResp)

	var decksResp DecksResponse
	_ = json.Unmarshal(apiResp.Data, &decksResp)
	return decksResp.Decks, apiResp.Code
}

func updateDeck(t *testing.T, id string, updates map[string]interface{}) Deck {
	resp, err := makeRequest("PUT", fmt.Sprintf("%s/%s", apiURL, id), updates)
	assert.NoError(t, err)

	var apiResp APIResponse
	_ = parseResponse(resp, &apiResp)

	var deck Deck
	_ = json.Unmarshal(apiResp.Data, &deck)
	return deck
}

func deleteDeck(t *testing.T, id string) int {
	resp, err := makeRequest("DELETE", fmt.Sprintf("%s/%s", apiURL, id), nil)
	assert.NoError(t, err)

	var apiResp APIResponse
	_ = parseResponse(resp, &apiResp)
	return apiResp.Code
}

// ---------- Tests ----------

func TestCreateDeck(t *testing.T) {
	skipIfServerDown(t)

	deck := createDeck(t, "Test Create Deck", "A1")
	assert.Equal(t, "Test Create Deck", deck.Name)
	assert.Equal(t, "A1", deck.TargetLevel)

	deleteDeck(t, deck.ID)
}

func TestGetDecks(t *testing.T) {
	skipIfServerDown(t)

	decks, code := getAllDecks(t)
	assert.Equal(t, 200, code)
	assert.NotNil(t, decks)
}

func TestGetDeckByID(t *testing.T) {
	skipIfServerDown(t)

	createdDeck := createDeck(t, "Test Get Deck", "A2")

	fetchedDeck, code := getDeckByID(t, createdDeck.ID)
	assert.Equal(t, 200, code)
	assert.Equal(t, createdDeck.ID, fetchedDeck.ID)
	assert.Equal(t, "Test Get Deck", fetchedDeck.Name)

	deleteDeck(t, createdDeck.ID)
}

func TestUpdateDeck(t *testing.T) {
	skipIfServerDown(t)

	deck := createDeck(t, "Old Deck", "B1")

	updated := updateDeck(t, deck.ID, map[string]interface{}{
		"name":        "Updated Deck",
		"description": "Updated description",
		"category":    "updated",
	})
	assert.Equal(t, "Updated Deck", updated.Name)
	assert.Equal(t, "Updated description", updated.Description)
	assert.Equal(t, "updated", updated.Category)

	deleteDeck(t, deck.ID)
}

func TestDeleteDeck(t *testing.T) {
	skipIfServerDown(t)

	deck := createDeck(t, "To Be Deleted", "C1")

	code := deleteDeck(t, deck.ID)
	assert.Equal(t, 204, code)

	_, getCode := getDeckByID(t, deck.ID)
	assert.Equal(t, 404, getCode)
}
