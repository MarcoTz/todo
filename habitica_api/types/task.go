package types 

import (
  "fmt"
) 

type TaskType string

const (
	Habits         TaskType = "habit"
	Dailys         TaskType = "daily"
	Todos          TaskType = "todo"
	Rewards        TaskType = "reward"
	CopmletedTodos TaskType = "completedTodo"
)

type Task struct {
  id string
  completed bool 
  text string 
  notes string 
}

func TaskFromJson(json map[string]interface{}) (*Task,error) {
  err_msg := "Trying to decode task"

  id,ok := json["_id"].(string)
  if !ok { 
    return nil,&DecodingErr{field_name:"id",message:err_msg}
  }

  completed, ok := json["completed"].(bool)
  if !ok { 
    return nil,&DecodingErr{field_name:"completed",message:err_msg}
  }

  text,ok := json["text"].(string)
  if !ok { 
    return nil,&DecodingErr{field_name:"text",message:err_msg}
  }

  notes,ok := json["notes"].(string)
  if !ok { 
    return nil,&DecodingErr{field_name:"notes",message:err_msg}
  }

  task := Task { id:id, completed:completed, text:text, notes:notes} 

  return &task,nil
}

func ShowTask(task Task) string{
  var completed_str string
  if task.completed { completed_str = "Y" } else {completed_str = "N"} 
  return fmt.Sprintf("Task: %s %s, notes: %s",completed_str,task.text,task.notes)
}
