package handlers

import "github.com/adepte-myao/test_provider/internal/models"

type TaskService interface {
	GetRandomTask(sectionId int32, certAreaId int32) (models.Task, error)
}
