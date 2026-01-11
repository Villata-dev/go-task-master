package db

import (
	"database/sql"
	"go-task-master/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(filepath string) error {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		return err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"description" TEXT,
		"completed" BOOLEAN
	);`

	_, err = DB.Exec(createTableSQL)
	return err
}

func GetTasksDB() ([]models.Task, error) {
	rows, err := DB.Query("SELECT id, title, description, completed FROM tasks ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func CreateTaskDB(task models.Task) (int64, error) {
	result, err := DB.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", task.Title, task.Description, task.Completed)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func SeedTasksDB() error {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		tasks := []models.Task{
			{Title: "Complete project proposal", Description: "Write and submit the project proposal for the new client.", Completed: false},
			{Title: "Team meeting", Description: "Attend the weekly team sync-up meeting.", Completed: true},
			{Title: "Code review", Description: "Review the pull request from the junior developer.", Completed: false},
		}

		for _, task := range tasks {
			_, err := DB.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", task.Title, task.Description, task.Completed)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func UpdateTaskDB(id int, task models.Task) error {
	_, err := DB.Exec("UPDATE tasks SET title = ?, description = ?, completed = ? WHERE id = ?", task.Title, task.Description, task.Completed, id)
	return err
}

func DeleteTaskDB(id int) error {
	_, err := DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

func CompleteTaskDB(id int) error {
    _, err := DB.Exec("UPDATE tasks SET completed = ? WHERE id = ?", true, id)
    return err
}
