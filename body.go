package alidayu

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func getRequestBody(m map[string]string) (body map[string]string) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	signString := AppSecret
	for _, k := range keys {
		signString += k + m[k]
	}
	signString += AppSecret

	signByte := md5.Sum([]byte(signString))
	sign := strings.ToUpper(fmt.Sprintf("%x", signByte))
	m["sign"] = sign

	return m
}
