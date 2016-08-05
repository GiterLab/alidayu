package alidayu

import (
	"fmt"

	"github.com/GiterLab/urllib"
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
	return true, str

}
