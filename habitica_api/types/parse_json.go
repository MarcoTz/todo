package types

import ( 
  "fmt" 
) 

func ParseResponse(json map[string]interface{}) ([]Task,error) {
  success,ok := json["success"].(bool)
  if !ok { 
    return []Task{}, &DecodingErr{field_name:"success",message:"Trying to parse response"}
  } 
  if !success{
    return []Task{}, &BadRequestErr{contents:fmt.Sprintf("%s",json)}
  }


  data,ok := json["data"].([]interface{})
  if !ok {
    return []Task{}, &DecodingErr{field_name:"data",message:"Trying to parse response"}
  }


  tasks := make([]Task,0)
  for _,data_row := range data{
    row_map,ok := data_row.(map[string]interface{})
    if !ok {
      return []Task{}, &DecodingErr{field_name:"data",message:"Trying to parse response"}
    }


    ty,ok := row_map["type"].(string)
    if !ok {
      return []Task{}, &DecodingErr{field_name:"type",message:"Trying to parse response"}
    }
    if ty != "todo"{
      continue
    }
 
    task,err := TaskFromJson(row_map)
    if err != nil { return []Task{}, err }  

    tasks = append(tasks,*task)

  }

  return tasks,nil
}
