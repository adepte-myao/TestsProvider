package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adepte-myao/test_provider/internal/models"
	"github.com/sirupsen/logrus"
)

type GetRandomTopicTaskHandler struct {
	logger   *logrus.Logger
	taskServ TaskService
}

func NewGetRandomTopicTaskHandler(logger *logrus.Logger, taskService TaskService) *GetRandomTopicTaskHandler {
	return &GetRandomTopicTaskHandler{
		logger:   logger,
		taskServ: taskService,
	}
}

func (handler *GetRandomTopicTaskHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	handler.logger.Info("Get random topic task request received")

	var dto models.GetRandomTopicTaskDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		handler.logger.Error("Get random topic task: ", err.Error())
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := handler.taskServ.GetRandomTask(dto.SectionID, dto.CertAreaId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(rw).Encode(task)
}
