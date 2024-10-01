package habitica_api

import (
  "time"
  "rooxo/todo/habitica_api/types"
)


type GetTasksUser struct {
  Task_type *types.TaskType
  Due_date  *time.Time
}

func (g *GetTasksUser) ToRequest() ApiRequest{
  return ApiRequest {Endpoint:"tasks/user",Method:Get} 
}
func (handler *ApiHandler) GetTasks() ([]types.Task,error) {
  request := (&GetTasksUser{}).ToRequest()
  json,err := handler.PerformRequest(request)
  if err != nil { return []types.Task{},err } 

  tasks,err := types.ParseResponse(json)
  if err != nil { return []types.Task{},err }

  return tasks,nil
}
