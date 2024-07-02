package store

import (
	"errors"

	"github.com/majimaccho/go-todo-app/entity"
)

var (
	Tasks = &TaskStore{Tasks: make(map[entity.TaskID]*entity.Task)}

	ErrNotFount = errors.New("not found")
)

type TaskStore struct {
	LastId entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastId++
	t.ID = ts.LastId
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.Task, len(ts.Tasks))
	for i, t := range ts.Tasks {
		tasks[i-1] = t
	}

	return tasks
}
