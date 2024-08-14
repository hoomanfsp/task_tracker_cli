package input

import (
	"bufio"
	"database/sql"
	"fmt"
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
		err := addTask(db, data[1], data[2])
		if err != nil {
			panic(err)
		}

	case "delete":
		deleteTask(db, data[1])
	case "update":
		updateTaskDescription(db, data[1], data[2])
	case "mark-in-progress":
		updateTaskStatus(db, data[1], "pending")
	case "mark-done":
		updateTaskStatus(db, data[1], "completed")
	case "list":
		ListOfTasks, err := getAllTasks(db)
		if err != nil {
			panic(err)
		}
		listprint(ListOfTasks)
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

func updateTaskStatus(db *sql.DB, id, status string) error {
	query := fmt.Sprintf("UPDATE tasks SET status = %s WHERE id = %s", status, id)
	_, err := db.Exec(query, status, id)
	return err
}

func deleteTask(db *sql.DB, id string) error {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}

func updateTaskDescription(db *sql.DB, id, description string) error {
	query := "UPDATE tasks SET description = ? WHERE id = ?"
	_, err := db.Exec(query, description, id)
	return err
}
func listprint(tasks []Task) {
	for _, task := range tasks {
		output := fmt.Sprintf(`
		id : %d
		status : %s
		creater at : %s
		title : %s
		description :%s
		`, task.ID, task.Status, task.CreatedAt, task.Title, task.Description)
		fmt.Println(output)

	}
}
