package task

import "time"

type Task struct {
	ID          string    `form:"id" json:"id" xml:"id"`
	Content     string    `form:"content" json:"content" xml:"content" binding:"required"`
	UserID      string    `form:"user_id" json:"user_id" xml:"user_id"`
	CreatedDate time.Time `form:"created_date" json:"created_date" xml:"created_date" time_format:"2006-01-02"`
}

type TaskFilter struct {
	CreatedDate time.Time `form:"created_date" json:"created_date" xml:"created_date" time_format:"2006-01-02"`
}
