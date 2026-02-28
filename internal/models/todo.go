package models

import "time"

type Todo struct {
	Id      int64     `json:"id"`
	Name    string    `binding:"required" json:"name"`
	Desc    string    `binding:"required" json:"description"`
	Done    *bool     `json:"isDone,omitempty"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
