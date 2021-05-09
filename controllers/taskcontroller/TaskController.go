package taskcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/models/config"
	"github.com/ocsen-hoc-code/go-auj/models/task"
	"github.com/ocsen-hoc-code/go-auj/models/user"
	"github.com/ocsen-hoc-code/go-auj/services/taskservice"
)

type TaskController struct {
	taskServ  *taskservice.TaskService
	jwtConfig *config.JWTConfig
}

func NewTaskController(taskServInject *taskservice.TaskService, jwtConfigInject *config.JWTConfig) *TaskController {
	return &TaskController{taskServ: taskServInject, jwtConfig: jwtConfigInject}
}

func (ctrl *TaskController) Create(c *gin.Context) {
	t := &task.Task{}
	u := c.MustGet("user").(*user.User)
	t.UserID = u.ID
	if err := c.ShouldBindJSON(t); nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rs, err := ctrl.taskServ.AddTask(c.Request.Context(), t)

	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": rs})
}

func (ctrl *TaskController) View(c *gin.Context) {
	taskFilter := &task.TaskFilter{}
	if err := c.ShouldBindQuery(&taskFilter); nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := c.MustGet("user").(*user.User)

	tasks, err := ctrl.taskServ.GetTasks(c.Request.Context(), u.ID, taskFilter.CreatedDate)

	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}
