package services

import (
	"github.com/adepte-myao/test_provider/internal/models"
	"github.com/sirupsen/logrus"
)

type TaskServiceBase struct {
	logger   *logrus.Logger
	taskRepo TaskRepository
}

func NewTaskServiceBase(logger *logrus.Logger, taskRepo TaskRepository) *TaskServiceBase {
	return &TaskServiceBase{logger: logger, taskRepo: taskRepo}
}

func (service *TaskServiceBase) GetRandomTask(sectionId int32, certAreaId int32) (models.Task, error) {
	task, err := service.taskRepo.GetRandom(sectionId, certAreaId)
	if err != nil {
		service.logger.Error(err)
		return models.Task{}, err
	}
	return task, nil
}
