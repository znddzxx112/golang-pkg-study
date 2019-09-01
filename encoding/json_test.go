package encodingt

import (
	"encoding/json"
	"testing"
)

const (
	CODE_SUCESS     = 0
	CODE_ERR_COMMON = -1
)

type ResponseResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type ResponseCallbackResult struct {
	ResponseResult
	Num int `json:"num"`
}

func TestMarshal(t *testing.T) {
	callbackResult := ResponseCallbackResult{
		ResponseResult{
			CODE_ERR_COMMON,
			"err msg",
			"",
		},
		1,
	}
	t.Log(callbackResult.Code)
	t.Log(callbackResult.Num)
	MarshalBytes, MarshalErr := json.Marshal(callbackResult)
	if MarshalErr != nil {
		t.Fatal(MarshalErr)
	}
	t.Log(string(MarshalBytes))
}

func TestUnmarshal(t *testing.T) {
	marshaBytes := `{"code":-1,"msg":"err msg","data":"","num":1}`

	var responseResult ResponseResult
	if UnmarshalErr := json.Unmarshal([]byte(marshaBytes), &responseResult); UnmarshalErr != nil {
		t.Fatal(UnmarshalErr)
	}
	t.Log(responseResult.Code)

	var responseCallbackResult ResponseCallbackResult
	if UnmarshalErr := json.Unmarshal([]byte(marshaBytes), &responseCallbackResult); UnmarshalErr != nil {
		t.Fatal(UnmarshalErr)
	}
	t.Log(responseCallbackResult.Code)
	t.Log(responseCallbackResult.Num)
}
