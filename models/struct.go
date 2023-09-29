package models

type User struct {
	UserID   int64  `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TODO struct {
	ID      int64  `json:"id"` //update
	UserID  int64  `json:"userID"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
