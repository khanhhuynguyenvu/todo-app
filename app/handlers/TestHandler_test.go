package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/todo-app/app/models"
	"github.com/volatiletech/null/v8"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type FakeRepo struct{}

func (f *FakeRepo) Create(task *models.Task) (*models.Task, error) {
	task.ID = 1
	return task, nil
}

func (f *FakeRepo) SelectAll() ([]*models.Task, error) {
	task1 := models.Task{
		ID:      1,
		Title:   null.StringFrom("Title"),
		Content: "Content",
	}
	task2 := models.Task{
		ID:      2,
		Title:   null.StringFrom("Title"),
		Content: "Content",
	}
	task3 := models.Task{
		ID:      3,
		Title:   null.StringFrom("Title"),
		Content: "Content",
	}
	tasks := []*models.Task{&task1, &task2, &task3}
	return tasks, nil
}

func (f *FakeRepo) Select(id string) (*models.Task, error) {
	task := models.Task{
		ID:      10,
		Title:   null.StringFrom("Title"),
		Content: "Content Select",
	}
	fmt.Println(task)
	return &task, nil
}

func (f *FakeRepo) UpdateById(id string, task *models.Task) (*models.Task, error) {
	return nil, nil
}

func (f *FakeRepo) DeleteById(id string) (*models.Task, error) {
	return nil, nil
}

func (f *FakeRepo) Delete() error {
	return nil
}

func TestCreate(t *testing.T) {
	payload := `{
    "title": "note today",
    "content": "note message today",
    "created_at":"2021-05-21T15:47:05Z"
	}`
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("CREATE", "/v1/tasks", strings.NewReader(payload))
	hTask := TaskHandler{
		&FakeRepo{},
	}
	hTask.Create(c)
	assert.Nil(t, c.Errors.Last())
	assert.Equal(t, 201, w.Code)
	assert.NotNil(t, w.Body)
	task := &models.Task{}
	err := json.Unmarshal([]byte(w.Body.String()), task)
	assert.Nil(t, err)
	assert.Greater(t, task.ID, 0, "Must be above 0")
	assert.Equal(t, task.Title.String, "note today")
}

func TestSelect(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/v1/tasks/10", nil)
	hTask := TaskHandler{
		&FakeRepo{},
	}
	hTask.Select(c)
	var task models.Task
	err := json.Unmarshal([]byte(w.Body.String()), &task)
	assert.Nil(t, c.Errors.Last())
	assert.Nil(t, err)
	assert.NotNil(t, task)
	fmt.Println(w.Body.String())
	assert.Equal(t, 10, task.ID)
	assert.Equal(t, "Title", task.Title.String)
	assert.Equal(t, "Content Select", task.Content)
}

func TestSelectAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/v1/tasks", nil)
	hTask := TaskHandler{
		&FakeRepo{},
	}
	hTask.All(c)
	var tasks []*models.Task
	err := json.Unmarshal([]byte(w.Body.String()), &tasks)
	assert.Nil(t, err)
	assert.Nil(t, c.Errors.Last())
	assert.IsType(t, &models.Task{}, tasks[0])
	assert.Equal(t, 3, len(tasks))
	assert.Equal(t, 1, tasks[0].ID)
	assert.Equal(t, 2, tasks[1].ID)
	assert.Equal(t, 3, tasks[2].ID)
}
