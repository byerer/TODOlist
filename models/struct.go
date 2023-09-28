package models

type TODO struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
