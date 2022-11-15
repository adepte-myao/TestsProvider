package storage

import (
	"database/sql"

	"github.com/adepte-myao/test_provider/internal/models"
)

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (repo *PostgresTaskRepository) GetRandom(sectionID int32, cerAreaID int32) (models.Task, error) {
	task := models.Task{
		Options: make([]string, 0),
	}

	tx, _ := repo.db.Begin()
	err := tx.QueryRow("SELECT id, question, answer "+
		"FROM tasks LIMIT 1 OFFSET (random() * (SELECT COUNT(*) FROM tasks));",
	).Scan(&task.ID, &task.Question, &task.Answer)
	if err != nil {
		return models.Task{}, nil
	}

	rows, err := tx.Query("SELECT answer_option FROM options WHERE question_id = $1", task.ID)
	if err != nil {
		return models.Task{}, nil
	}

	var option string
	for rows.Next() {
		err = rows.Scan(&option)
		if err != nil {
			return task, nil
		}
		task.Options = append(task.Options, option)
	}

	return task, nil
}
