package helper

import "strings"


// Response is used to shape return json
type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Error interface{} `json:"error"`
	Data interface{}	`json:"data"`
}

//EmptyObj used to send data that should not be null on json
type EmptyObj struct {

}

//BuildResponse method to inject data value send out
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status: status, 
		Message: message,
		Error: nil,
		Data: data,
	}

	return res
}

//Error builder to send error to the frontend.
func BuildErrorResponse(message string, err string, data interface{}) Response {
	
	splittedError := strings.Split(err,"\n")
	
	res := Response{
		Status: false,
		Message: message,
		Error: splittedError,
		Data: data,
	}

	return res
}
