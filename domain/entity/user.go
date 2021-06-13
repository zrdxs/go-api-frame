package entity

// User holds user struct
type User struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Age    int64  `json:"age"`
	Email  string `json:"email"`
}
