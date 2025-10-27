package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

type Command struct {
	name               string
	desc               string
	argsNum            int
	command_to_execute func()
}

type Task struct {
	ID     int    `json:"id"`
	DESC   string `json:"desc"`
	STATUS string `json:"status"`
}

// global array has all the available commands
var commands = []Command{
	{"add", "add task to the db", 3, func() { addTask(os.Args[2]) }},
	{"update", "update a task", 4, func() { updateTask(os.Args[2], os.Args[3]) }},
	{"delete", "remove a task from db", 3, func() { deleteTask(os.Args[2]) }},
	{"mark-in", "mark a task as in-progress", 3, func() { markInTask(os.Args[2]) }},
	{"mark-done", "mark a task as done", 3, func() { markDoneTask(os.Args[2]) }},
	{"list", "list all the tasks", 0, func() {
		if len(os.Args) > 2 {
			listTasks(os.Args[2:]...) // pass remaining args to variadic function
		} else {
			listTasks() // no arguments
		}
	}},
}

var tasks []Task
var lastId int = 0

// Define colors
var (
	todoColor    = color.New(color.FgYellow).SprintFunc()
	doneColor    = color.New(color.FgGreen).SprintFunc()
	inprogColor  = color.New(color.FgCyan).SprintFunc()
	headerColor  = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	borderColor  = color.New(color.FgHiBlack).SprintFunc()
	messageColor = color.New(color.FgHiMagenta).SprintFunc()
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("too few arguments")
		return
	}

	tasks = loadTasks()

	command, err := validator(os.Args)
	if err != nil  {
		fmt.Println("argument is not correct")
		return
	}

	execute_command(command)
	saveTasks()
}

func loadTasks() []Task {
	path := "/home/mada/go/db.json"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("Failed to open the file")
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		fmt.Println("Failed to open the file")
	}

	if info.Size() == 0 {
		return []Task{}
	}

	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("Failed to decode tasks")
	}
	return tasks
}

func saveTasks() {
	file, err := os.Create("/home/mada/go/db.json")
	if err != nil {
		fmt.Println("Failed to open file db.json")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Failed to encode tasks")
	}
}

func execute_command(command Command) {
	command.command_to_execute()
}

func addTask(task string) {
	for i := range tasks {
		if task == tasks[i].DESC {
			fmt.Println("this task already exists")
			return
		}
	}

	newTask := Task{len(tasks), task, "todo"}
	tasks = append(tasks, newTask)
	lastId += 1

	fmt.Println(messageColor("✅ Task added successfully!"))
}

func updateTask(id, task string) {
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Failed to convert ID")
		return
	}
	for i := range tasks {
		if num == tasks[i].ID {
			tasks[i].DESC = task
			fmt.Println(messageColor("✏️ Task updated successfully!"))
			return
		}
	}
	fmt.Println("Couldn't find a task with such ID")
}

func deleteTask(id string) {
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Failed to convert ID")
		return
	}
	for i := range tasks {
		if num == tasks[i].ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println(messageColor("🗑️ Task deleted successfully!"))
			return
		}
	}
	fmt.Println("This ID doesn't exist")
}

func markInTask(id string) {
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Failed to convert ID")
		return
	}
	for i := range tasks {
		if num == tasks[i].ID {
			tasks[i].STATUS = "in-progress"
			fmt.Println(messageColor("🚧 Task marked as in-progress"))
			return
		}
	}
	fmt.Println("This ID doesn't exist")
}

func markDoneTask(id string) {
	num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Failed to convert ID")
		return
	}
	for i := range tasks {
		if num == tasks[i].ID {
			tasks[i].STATUS = "done"
			fmt.Println(messageColor("✅ Task marked as done"))
			return
		}
	}
	fmt.Println("This ID doesn't exist")
}

func listTasks(props ...string) {
	fmt.Println()
	fmt.Println(borderColor("+----+---------------------------+-------------+"))
	fmt.Printf("| %s | %s | %s |\n",
		headerColor("ID"),
		headerColor("Description"),
		headerColor("Status"),
	)
	fmt.Println(borderColor("+----+---------------------------+-------------+"))

	count := 0
	for _, task := range tasks {
		if len(props) == 0 || task.STATUS == props[0] {
			status := task.STATUS
			switch status {
			case "todo":
				status = todoColor(status)
			case "done":
				status = doneColor(status)
			case "in-progress":
				status = inprogColor(status)
			}
			fmt.Printf("| %-2d | %-25s | %-11s |\n", task.ID, task.DESC, status)
			count++
		}
	}

	if count == 0 {
		if len(props) == 0 {
			fmt.Println("| No tasks found                              |")
		} else {
			fmt.Printf("| No tasks found with status '%s'             |\n", props[0])
		}
	}

	fmt.Println(borderColor("+----+---------------------------+-------------+"))
	fmt.Println()
}

func validator(args []string) (Command, error) {
	for i := range commands {
		if args[1] == commands[i].name {
			return commands[i], nil
		}
	}
	return Command{}, fmt.Errorf("unknown command %s", args[1])
}

