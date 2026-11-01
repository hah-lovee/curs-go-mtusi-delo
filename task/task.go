package task

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrEmptyNameTask = errors.New("empty name")
	ErrUnknownStatus = errors.New("unknown status")
)

type statusTask string

const (
	statusCreated statusTask = "Created"
	statusInWork  statusTask = "In work"
	statusDone    statusTask = "Done"
)

type Task struct {
	statusTask
	name string
	id   string
}

func CreateTask(name string) (*Task, error) {
	if strings.Compare(name, "") == 0 {
		return nil, ErrEmptyNameTask
	}

	task := Task{name: name, statusTask: "Created", id: uuid.New().String()}
	return &task, nil
}

func (t *Task) ChangeName(newName string) error {
	if strings.Compare(newName, "") == 0 {
		return ErrEmptyNameTask
	}
	t.name = newName
	return nil
}

func (t *Task) ChangeStatus(newStatus string) error {
	switch status := statusTask(newStatus); status {
	case statusCreated, statusInWork, statusDone:
		t.statusTask = status
		return nil
	default:
		return ErrUnknownStatus
	}
}

func (t *Task) GetId() string {
	return t.id
}

func (t *Task) GetName() string {
	return t.name
}

// status in string
func (t *Task) GetStatus() string {
	return string(t.statusTask)
}
