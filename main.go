package main 

import ( 
  "fmt"
  "rooxo/todo/server"
  "rooxo/todo/habitica_api"
)

func main(){
  handler,err := habitica_api.SetupApi()
  if err!=nil { 
    fmt.Printf("Error loading config:\n\t%s\n",err) 
    return 
  }

  err = server.StartServer(handler)
  if err != nil { fmt.Printf("Server exited with error %s\n",err) } 
}
