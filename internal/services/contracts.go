package services

import "github.com/adepte-myao/test_provider/internal/models"

type TaskRepository interface {
	GetRandom(sectionID int32, cerAreaID int32) (models.Task, error)
}
