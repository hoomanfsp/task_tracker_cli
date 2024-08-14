package input

import (
	"bufio"
	"database/sql"
	"os"
	"simp_task_cli/database"
	"strings"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   string
}

func Start() {
	reader := bufio.NewReader(os.Stdin)
	strin, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	order := strings.Split(strin, " ")
	db := database.Start()
	procces(db, order)
}
func procces(db *sql.DB, data []string) {
	opr := data[0]
	switch opr {
	case "add":
		addTask(db, data[1], data[2])
	case "update":

	case "mark-in-progress":

	case "mark-done":

	case "list":
	}
}

func addTask(db *sql.DB, title, description string) error {
	query := "INSERT INTO tasks (title, description) VALUES (?, ?)"
	_, err := db.Exec(query, title, description)
	return err
}

func getAllTasks(db *sql.DB) ([]Task, error) {
	query := "SELECT id, title, description, status, created_at FROM tasks"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
