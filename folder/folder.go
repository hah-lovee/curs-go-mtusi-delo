package folder

import (
	"errors"
	"go-mtuci/task"
	"slices"
)

var (
	ErrEmptyNameFolder = errors.New("Пустое имя папки")
	ErrAlreadyExists   = errors.New("Папка уже существует")
	ErrAlreadyOnList   = errors.New("Папки нет в списке")
)

type Folder struct {
	name  string
	tasks []*task.Task
}

func CreateMainFolder() ([]*Folder, map[string]*Folder) {
	folders := make([]*Folder, 0, 3)

	folders = append(folders, &Folder{name: "Входящие"})
	folders = append(folders, &Folder{name: "Работа"})
	folders = append(folders, &Folder{name: "Личное"})

	mapFolders := map[string]*Folder{
		"Входящие": folders[0],
		"Работа":   folders[1],
		"Личное":   folders[2],
	}
	return folders, mapFolders
}

func (f *Folder) GetName() string {
	return f.name
}

func CreateFolder(name string, mapfolders map[string]*Folder) (*Folder, error) {
	if name == "" {
		return &Folder{}, ErrEmptyNameFolder
	}
	if _, ok := mapfolders[name]; ok {
		return &Folder{}, ErrAlreadyExists
	}

	newFolder := Folder{
		name: name,
	}
	return &newFolder, nil
}

func (f *Folder) DeleteFolder(folders []*Folder, mapFolders map[string]*Folder) ([]*Folder, map[string]*Folder, error) {

	if i := slices.Index(folders, f); i >= 0 {
		delete(mapFolders, f.name)
		folders = slices.Delete(folders, i, i+1)
		return folders, mapFolders, nil
	} else {
		return folders, mapFolders, ErrAlreadyOnList
	}
}

func (f *Folder) RenameFolder(mapFolders map[string]*Folder, name string) error {
	if name == "" {
		return ErrEmptyNameFolder
	}
	if _, ok := mapFolders[name]; ok {
		return ErrAlreadyExists
	}

	f.name = name
	return nil

}

func (f *Folder) AddTask(t *task.Task) error {
	f.tasks = append(f.tasks, t)
	return nil
}
func (f *Folder) DeleteTask(index int) {
	f.tasks = slices.Delete(f.tasks, index, index+1)
}

func (f *Folder) GetTasks() []*task.Task {
	return f.tasks
}
