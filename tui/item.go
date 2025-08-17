package tui

import (
	"github.com/Sheriff-Hoti/beaver-task/database"
	"github.com/charmbracelet/bubbles/list"
)

type Item struct {
	ID              int64
	TaskTitle       string
	TaskDescription string
}

func (t *Item) FilterValue() string { return t.TaskTitle }
func (t *Item) Title() string       { return t.TaskTitle }
func (t *Item) Description() string { return t.TaskDescription }

func fromDatabaseTask(task *database.Task) *Item {
	return &Item{
		ID:              task.ID,
		TaskTitle:       task.Title,
		TaskDescription: task.Description.String,
	}
}

func FromDatabaseTasks(tasks []database.Task) []list.Item {
	items := make([]list.Item, len(tasks))
	for i, task := range tasks {
		items[i] = fromDatabaseTask(&task)
	}
	return items
}
