package entity

type Dictionary struct {
	Word     string `json:"word"`
	Entry    string `json:"entry"`
	Type     string `json:"type"`
	Tier     string `json:"tier"`
	Semantic string `json:"semantic"`
}
