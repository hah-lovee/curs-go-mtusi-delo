package main

import (
	"bufio"
	"fmt"
	"go-mtuci/folder"
	"go-mtuci/task"
	"os"
	"strconv"
	"strings"
)

func readLine(in *bufio.Scanner) string {
	in.Scan()                           // ждём, пока пользователь нажмёт Enter
	return strings.TrimSpace(in.Text()) // строка целиком, без пробелов по краям
}

func main() {
	var command string
	var pickFolder string
	var pickTask string
	var statusProgram string = "main"
	var currentFolder *folder.Folder
	var currentTasks []*task.Task

	var in = bufio.NewScanner(os.Stdin)

	folders, mapFolders := folder.CreateMainFolder()

	for {
		switch statusProgram {
		case "main":
			fmt.Println("=== папки ===")

			for i, folder := range folders {
				fmt.Printf("%d. %s\n", i+1, folder.GetName())
			}
			fmt.Println("0. Выход")

			pickFolder = readLine(in)
			switch pickFolder {
			case "0":
				return
			default:
				statusProgram = "folder"
			}
		case "folder":
			// получаем ответ пользователя
			// повторяющийся код
			number, err := strconv.Atoi(pickFolder)
			if err != nil {
				fmt.Println("Некорректный ввод")
				statusProgram = "main"
				continue
			}

			if number > len(folders) || number < 0 {
				fmt.Println("Нет папки с таким номером")
				statusProgram = "main"
				continue
			}
			// работаем внутри папки
			currentFolder = folders[number-1]
			fmt.Println("[" + currentFolder.GetName() + "]")
			fmt.Println("1. Создать папку")
			fmt.Println("2. Удалить папку")
			fmt.Println("3. Переименовать папку")
			fmt.Println("4. Создать задачу")
			fmt.Println("5. Показать список")
			fmt.Println("6. Назад")
			fmt.Println("0. Выход")

			command = readLine(in)

			switch command {
			case "1":
				fmt.Print("Напишите название папки: ")
				name := readLine(in)
				newFolder, err := folder.CreateFolder(name, mapFolders)
				if err != nil {
					fmt.Println("Создание папки:", err)
				} else {
					folders = append(folders, newFolder)
					mapFolders[name] = newFolder
					fmt.Println("Папка создана!")
				}
			case "2":
				folders, mapFolders, err = currentFolder.DeleteFolder(folders, mapFolders)
				if err != nil {
					fmt.Println("Удаление папки:", err)
				} else {
					fmt.Println("Папка удалена!")
					statusProgram = "main"
				}
			case "3":
				fmt.Print("Напишите новое название папки: ")
				name := readLine(in)
				err = currentFolder.RenameFolder(mapFolders, name)

				if err != nil {
					fmt.Println("Переименовывание папки:", err)
				} else {
					fmt.Println("Папка переименована!")
				}
			case "4":
				fmt.Print("Напишите название задачи: ")
				name := readLine(in)
				newTask, err := task.CreateTask(name)
				if err != nil {
					fmt.Println("Создание задачи:", err)
				} else {
					err = currentFolder.AddTask(newTask)
					if err != nil {
						fmt.Println("Создание задачи: ", err)
					}
				}

			case "5":
				statusProgram = "tasks"
			case "6":
				statusProgram = "main"
			case "0":
				return
			default:
				fmt.Print("неизвестная команда")
			}

		case "tasks":
			currentTasks = currentFolder.GetTasks()
			fmt.Println(" === Задачи в папке " + currentFolder.GetName() + " ===")
			for i, task := range currentTasks {
				fmt.Printf("%d. %s\n", i+1, task.GetName())
			}
			fmt.Printf("%d. Назад\n", len(currentTasks)+1)
			fmt.Println("0. Выход")

			pickTask = readLine(in)
			switch pickTask {
			case strconv.Itoa(len(currentTasks) + 1):
				statusProgram = "folder"
			case "0":
				return
			default:
				statusProgram = "task"
			}
		case "task":
			// получаем ответ пользователя
			number, err := strconv.Atoi(pickTask)
			if err != nil {
				fmt.Println("Некорректный ввод")
				statusProgram = "main"
				continue
			}

			if number > len(folders) || number < 0 {
				fmt.Println("Нет папки с таким номером")
				statusProgram = "main"
				continue
			}
			currentTask := currentTasks[number-1]
			fmt.Printf(" === задача %s === \n", currentTask.GetName())
			fmt.Println("1. Изменить название")
			fmt.Println("2. Изменить Статус")
			fmt.Println("3. Удалить задачу")
			fmt.Println("4. Назад")
			fmt.Println("0. Выход")

			command = readLine(in)

			switch command {
			case "1":
				fmt.Println("Напишите новое название задачи: ")
				name := readLine(in)
				err = currentTask.ChangeName(name)

				if err != nil {
					fmt.Println("Переименовывание задачи:", err)
				} else {
					fmt.Println("Задача переименована!")
				}
			case "2":
				fmt.Print("Задайте новый статус: ")
				fmt.Println("статус может быть: In work/Done")
				newStatus := readLine(in)

				err = currentTask.ChangeStatus(newStatus)

				if err != nil {
					fmt.Println("Смена статуса:", err)
				} else {
					fmt.Println("Статус изменен на " + newStatus)
				}
			case "3":
				currentFolder.DeleteTask(number - 1)
				fmt.Println("Задача успешно удалена")
				statusProgram = "tasks"
			case "4":
				statusProgram = "tasks"
			case "0":
				return
			}
		}
	}
}
