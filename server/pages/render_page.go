package pages

import ( 
  "fmt"
  "strings"
  "rooxo/todo/habitica_api/types"
)

const PageTemplate string = `
<!doctype html>
<html>
  <head><title>%s</title></head>
  <body>
    <ul>
     %s
    </ul>
  </body>
</html>`

func RenderIndex(tasks []types.Task) string{
  lis := make([]string,0)
  for _,task := range tasks{
    next_li := fmt.Sprintf("<li>%s</li>",types.ShowTask(task))
    lis = append(lis,next_li)
  }
  return fmt.Sprintf(PageTemplate,"Index", strings.Join(lis,"\n"))
}
