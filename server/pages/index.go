package pages

import (
	"fmt"
	"net/http"
	"rooxo/todo/habitica_api"
)

func HandleIndex(api_handler *habitica_api.ApiHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := api_handler.GetTasks()
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		fmt.Fprint(w, RenderIndex(tasks))
	}
}
