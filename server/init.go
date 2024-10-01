package server 

import (
  "net/http"
  "rooxo/todo/server/pages"
  "rooxo/todo/habitica_api"
)

func StartServer(handler *habitica_api.ApiHandler) error {
  http.HandleFunc("/", pages.HandleIndex(handler))
  return http.ListenAndServe(":8080",nil)
}
