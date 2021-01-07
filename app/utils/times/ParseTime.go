package times

import (
	utils "github.com/todo-app/app/utils/erros"
	"time"
)

func parse(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	utils.PanicError(err)
	return t
}
