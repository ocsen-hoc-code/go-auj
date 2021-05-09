package user_test

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/models/service"
	"github.com/ocsen-hoc-code/go-auj/models/task"
	util_test "github.com/ocsen-hoc-code/go-auj/tests/utils"
	"github.com/stretchr/testify/assert"
)

func getToken(r *gin.Engine) string {
	values := map[string]string{"username": util_test.USERNAME, "password": util_test.PASSWORD}
	jsonData, _ := json.Marshal(values)
	_, body := util_test.CreateRequest(r, util_test.POST, "/login", "", jsonData)
	return body["data"].(string)
}

func PostTask(r *gin.Engine, content, token string) (int, task.Task) {
	values := map[string]string{"content": content}
	jsonData, _ := json.Marshal(values)
	status, body := util_test.CreateRequest(r, util_test.POST, "/tasks", token, jsonData)
	task := task.Task{}
	dataByte, _ := json.Marshal(body["data"])
	json.Unmarshal(dataByte, &task)
	return status, task
}

func TestTask(t *testing.T) {
	serv := util_test.InstallTest()
	defer func(s *service.Service) {
		serv.Config.Database.Close()
	}(serv)

	t.Run("TestPostTask", func(t *testing.T) {
		token := getToken(serv.Server)
		content := "Minh Dep Trai"
		status, task := PostTask(serv.Server, content, token)
		assert.Equal(t, 201, status, "Expect 201")
		assert.Equal(t, util_test.USERID, task.UserID, "Expect "+util_test.USERID)
		assert.Equal(t, content, task.Content, "Expect "+content)
	})

	t.Run("TestPostTaskWithTokenExpired", func(t *testing.T) {
		content := "Minh Dep Trai"
		values := map[string]string{"content": content}
		jsonData, _ := json.Marshal(values)
		status, body := util_test.CreateRequest(serv.Server, util_test.POST, "/tasks", util_test.OLD_JWT, jsonData)
		assert.Equal(t, 401, status, "Expect 401")
		assert.Equal(t, "Token is invalid!", body["message"], "Expect Message: Token is invalid!")
	})

	// t.Run("TestPostTaskLimit", func(t *testing.T) {
	// 	token := getToken(serv.Server)
	// 	content := "Minh Dep Trai"
	// 	for i := 0; i < 5; i++ {
	// 		status, task := PostTask(serv.Server, content, token)
	// 		assert.Equal(t, 201, status, "Expect 201")
	// 		assert.Equal(t, util_test.USERID, task.UserID, "Expect "+util_test.USERID)
	// 		assert.Equal(t, content, task.Content, "Expect "+content)
	// 	}
	// 	msgError := "Users are limited to create only 5 task only per day!"
	// 	values := map[string]string{"content": content}
	// 	jsonData, _ := json.Marshal(values)
	// 	status, body := util_test.CreateRequest(serv.Server, util_test.POST, "/tasks", token, jsonData)
	// 	assert.Equal(t, 400, status, "Expect 400")
	// 	assert.Equal(t, msgError, body["error"], "Expect error: "+msgError)
	// })

	t.Run("TestGetTask", func(t *testing.T) {
		token := getToken(serv.Server)
		content := "Minh Dep Trai"
		status, body := util_test.CreateRequest(serv.Server, util_test.GET, "/tasks", token, nil)
		var tasks []task.Task
		dataByte, _ := json.Marshal(body["data"])
		json.Unmarshal(dataByte, &tasks)
		assert.Equal(t, 200, status, "Expect 200")
		assert.Equal(t, util_test.USERID, tasks[0].UserID, "Expect "+util_test.USERID)
		assert.Equal(t, content, tasks[0].Content, "Expect "+content)
	})

	t.Run("TestGetTaskWithCreatedTime", func(t *testing.T) {
		token := getToken(serv.Server)
		content := "Minh Dep Trai"
		status, body := util_test.CreateRequest(serv.Server, util_test.GET, "/tasks", token, nil)
		var tasks []task.Task
		dataByte, _ := json.Marshal(body["data"])
		json.Unmarshal(dataByte, &tasks)
		assert.Equal(t, 200, status, "Expect 200")

		assert.Equal(t, util_test.USERID, tasks[0].UserID, "Expect "+util_test.USERID)
		assert.Equal(t, content, tasks[0].Content, "Expect "+content)
	})

	t.Run("TestGetTaskWithTokenExpired", func(t *testing.T) {
		status, body := util_test.CreateRequest(serv.Server, util_test.GET, "/tasks", util_test.OLD_JWT, nil)
		assert.Equal(t, 401, status, "Expect 401")
		assert.Equal(t, "Token is invalid!", body["message"], "Expect Message: Token is invalid!")
	})
}
