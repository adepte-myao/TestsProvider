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
	defer tx.Rollback()

	err := tx.QueryRow(`WITH suit_tasks AS (
		SELECT tasks.id, tasks.question, tasks.answer
			FROM tasks
			JOIN tests t on tasks.test_id = t.id
			JOIN cert_area ca on t.cert_area_id = ca.id
		WHERE section_id = $1 AND ca.id = $2
		)
		SELECT suit_tasks.id, suit_tasks.question, suit_tasks.answer
			FROM suit_tasks
		LIMIT 1 OFFSET (random() * (SELECT COUNT(*) FROM suit_tasks));`,
		sectionID, cerAreaID,
	).Scan(&task.ID, &task.Question, &task.Answer)
	if err != nil {
		return models.Task{}, err
	}

	rows, err := tx.Query("SELECT answer_option FROM options WHERE question_id = $1", task.ID)
	if err != nil {
		return models.Task{}, err
	}

	var option string
	for rows.Next() {
		err = rows.Scan(&option)
		if err != nil {
			return task, err
		}
		task.Options = append(task.Options, option)
	}
	tx.Commit()

	return task, nil
}
