package model

type Task struct {
	ID          string `csv:"ID"`
	Description string `csv:"Description"`
	CreatedAt   string `csv:"CreatedAt"`
	IsCompleted bool   `csv:"IsCompleted"`
}
