package taskservice

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/ocsen-hoc-code/go-auj/models/config"
	"github.com/ocsen-hoc-code/go-auj/models/task"
	taskrepository "github.com/ocsen-hoc-code/go-auj/reponsitories/task"
)

type TaskService struct {
	reps      *taskrepository.TaskRepository
	jwtConfig *config.JWTConfig
}

func NewTaskService(taskRepsInject *taskrepository.TaskRepository, jwtConfigInject *config.JWTConfig) *TaskService {
	return &TaskService{reps: taskRepsInject, jwtConfig: jwtConfigInject}
}

func (serv TaskService) AddTask(ctx context.Context, t *task.Task) (*task.Task, error) {
	now := time.Now()
	t.ID = uuid.New().String()
	t.CreatedDate = now
	isLimitTask, maxToDo := serv.reps.ValidateTask(ctx, t.UserID, now)
	if isLimitTask {
		return t, errors.New("Users are limited to create only " + strconv.Itoa(maxToDo) + " task only per day!")
	}
	err := serv.reps.CreateTask(ctx, t)

	return t, err
}

func (serv TaskService) GetTasks(ctx context.Context, userId string, createdDate time.Time) ([]*task.Task, error) {
	tasks, err := serv.reps.FindTasks(ctx, sql.NullString{String: userId, Valid: true}, createdDate)
	return tasks, err
}
