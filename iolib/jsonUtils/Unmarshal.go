package jsonUtils

import "encoding/json"

//json转bean, 这里的interface{}一般指struct的指针
func Unmarshal(jsonStr string,bean interface{}) error{

	s, err := Discard(jsonStr)
	if err!=nil {
		return err
	}
	err = json.Unmarshal([]byte(s), bean)
	return err
}