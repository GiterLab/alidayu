package alidayu

import (
	"fmt"

	"github.com/GiterLab/urllib"
	"github.com/buger/jsonparser"
)

func DoPost(m map[string]string) (success bool, response string) {
	if AppKey == "" || AppSecret == "" {
		return false, "AppKey or AppSecret is requierd!"
	}
	body := getRequestBody(m)
	fmt.Println(body)
	req := urllib.Post(URL)
	for v := range body {
		req.Param(v, body[v])
	}
	req.Header("Content-Type", "application/x-www-form-urlencoded")

	str, err := req.String()
	if err != nil {
		fmt.Println("err,", err)
		return false, str
	}

	if !check(str) {
		return false, str
	}

	return true, str

}

func check(str string) bool {
	data := []byte(str)
	err_code, _, _, err := jsonparser.Get(data, "alibaba_aliqin_fc_sms_num_send_response", "result", "err_code")
	if err != nil {
		return false
	}

	if string(err_code) == "0" {
		return true
	}

	return false
}
