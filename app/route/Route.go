package route

import (
	"github.com/gin-gonic/gin"
	"github.com/todo-app/app/handlers"
	repo "github.com/todo-app/app/repositories"
)

func Route(r *gin.Engine) *gin.Engine {
	version := r.Group("/v1")
	{
		task := version.Group("/tasks")
		{
			hTask := handlers.TaskHandler{
				TaskRepo: &repo.TaskRepoBoiler{},
			}

			task.POST("", hTask.Create)
			task.GET("", hTask.All)
			task.PATCH("")
			task.DELETE("", hTask.Delete)

			specific := task.Group("/:id")
			{
				specific.POST("")
				specific.GET("", hTask.Select)
				specific.PATCH("", hTask.UpdateById)
				specific.DELETE("", hTask.DeleteById)

				subtask := specific.Group("/subtasks")
				{
					subtask.POST("")
					subtask.GET("")
					subtask.DELETE("")
					subtask.PATCH("")
				}
			}
		}
	}
	return r
}
