package route

import (
	"github.com/gin-gonic/gin"
	"github.com/todo-app/app/handlers"
	repo "github.com/todo-app/app/repositories"
)

func Route(r *gin.Engine) *gin.Engine {
	version := r.Group("/v1")
	{
		task := version.Group("/task")
		{
			hTask := handlers.TaskHandler{
				TaskRepo: &repo.TaskRepo{},
			}

			task.POST("", hTask.Create)
			task.GET("", hTask.All)
			task.PATCH("")
			task.DELETE("")

			specific := task.Group("/:id")
			{
				specific.POST("")
				specific.GET("", hTask.Select)
				specific.PATCH("")
				specific.DELETE("", hTask.Delete)

				subtask := specific.Group("/subtask")
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
