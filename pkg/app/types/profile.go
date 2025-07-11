package types

import "github.com/nleiva/go-todo-api/pkg/app/model"

// ProfileData represents user profile with statistics
type ProfileData struct {
	Account *model.Account `json:"account"`
	Stats   ProfileStats   `json:"stats"`
}

// ProfileStats represents user statistics
type ProfileStats struct {
	TotalTodos     int64 `json:"total_todos"`
	CompletedTodos int64 `json:"completed_todos"`
	PendingTodos   int64 `json:"pending_todos"`
	CompletionRate int   `json:"completion_rate"` // percentage
}
