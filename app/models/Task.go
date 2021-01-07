package models

type Task struct {
	Id        int64    `json:"id"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	CreatedAt string `json:"created_at" binding:"required"`
	UpdatedAt string `json:"updated_at"`
	DeleteAt  string `json:"deleted_at"`
}
