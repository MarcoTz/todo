package habitica_api

import (
	"fmt"
	"net/http"
	"os"
  "io"
  "encoding/json"
)

const BASE_URL string = "https://habitica.com/api/v3/"
const API_KEY_ENV string = "HABITICA_API_KEY"
const USER_ID_ENV string = "HABITICA_USER_ID"

type RequestMethod string

const (
	Get  RequestMethod = http.MethodGet
	Post RequestMethod = http.MethodPost
)

type ApiHandler struct {
	api_key string
  user_id string
}

type ApiRequest struct {
	Endpoint string
	Method   RequestMethod
}

func SetupApi() (*ApiHandler, error) {
  api_key := os.Getenv(API_KEY_ENV) 
  if api_key == "" { return nil,&MissingVarErr{variable:API_KEY_ENV} }
  user_id := os.Getenv(USER_ID_ENV)
  if user_id == "" { return nil,&MissingVarErr{variable:USER_ID_ENV} } 
  return &ApiHandler{api_key:api_key,user_id:user_id}, nil
}

func (handler *ApiHandler) PerformRequest(request ApiRequest) (map[string]interface{}, error) {
	request_url := BASE_URL + request.Endpoint
	req, err := http.NewRequest(string(request.Method), request_url, nil)
	if err != nil {
		return nil, &CreateRequestErr{message: fmt.Sprintf("%s", err)}
	}

  req.Header.Add("Content-Type","application/json")
  req.Header.Add("x-api-user",handler.user_id)
  req.Header.Add("x-api-key",handler.api_key)
  req.Header.Add("x-client", fmt.Sprintf("%s-todo",handler.user_id))

	http_res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, &PerformRequestErr{message: fmt.Sprintf("%s", err)}
	}

  buf,err := io.ReadAll(http_res.Body)
  if err != nil { return nil,err }

  if http_res.StatusCode != 200 {
    return nil,&ResponseErr{code:http_res.StatusCode,message:string(buf)}
  }

  var json_res map[string]interface{}
  err = json.Unmarshal(buf, &json_res)

	return json_res, nil
}
