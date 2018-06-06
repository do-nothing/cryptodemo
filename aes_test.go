package main

import (
	"testing"
	"fmt"
	"encoding/base64"
)

func TestAes(t *testing.T) {
	key := []byte("a123456789b23456")
	result, err := AesEncrypt([]byte("测试-->{'aaa':'bbb'}"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))

	result, err = base64.StdEncoding.DecodeString("xyUdDpguOmAfQoCFeKWJnQ+o/RxkH0IctvuIADUug/w=")
	origData, err := AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))

}
