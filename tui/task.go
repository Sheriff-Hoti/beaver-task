package tui

import (
	"github.com/Sheriff-Hoti/beaver-task/database"
	"github.com/charmbracelet/bubbles/list"
)

type Task struct {
	ID              int64
	TaskTitle       string
	TaskDescription string
}

func (t Task) FilterValue() string { return t.TaskTitle }
func (t Task) Title() string       { return t.TaskTitle }
func (t Task) Description() string { return t.TaskDescription }

func fromDatabaseTask(task *database.Task) *Task {
	return &Task{
		ID:              task.ID,
		TaskTitle:       task.Title,
		TaskDescription: task.Description.String,
	}
}

func fromDatabaseTasks(tasks []database.Task) []list.Item {
	items := make([]list.Item, len(tasks))
	for i, task := range tasks {
		items[i] = fromDatabaseTask(&task)
	}
	return items
}
