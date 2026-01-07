package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB(filepath string) error {
	var err error
	db, err = sql.Open("sqlite3", filepath)
	if err != nil {
		return err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"description" TEXT,
		"completed" BOOLEAN
	);`

	_, err = db.Exec(createTableSQL)
	return err
}

func getTasksDB() ([]Task, error) {
	rows, err := db.Query("SELECT id, title, description, completed FROM tasks ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func createTaskDB(task Task) (int64, error) {
	result, err := db.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", task.Title, task.Description, task.Completed)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func seedTasksDB() error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		tasks := []Task{
			{Title: "Complete project proposal", Description: "Write and submit the project proposal for the new client.", Completed: false},
			{Title: "Team meeting", Description: "Attend the weekly team sync-up meeting.", Completed: true},
			{Title: "Code review", Description: "Review the pull request from the junior developer.", Completed: false},
		}

		for _, task := range tasks {
			_, err := db.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", task.Title, task.Description, task.Completed)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
