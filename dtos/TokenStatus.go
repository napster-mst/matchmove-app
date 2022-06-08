package dtos

type TokenStatusResponse struct {
	ActiveTokens   int `json:"active-tokens"`
	InActiveTokens int `json:"inactive-tokens"`
}
