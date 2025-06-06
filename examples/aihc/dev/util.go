package dev

import (
	"encoding/json"
	"log"
	"os"

	api "github.com/baidubce/bce-sdk-go/services/aihc/dev"
)

func ReadJson(fileName string) (*api.CreateDevInstanceArgs, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	var args api.CreateDevInstanceArgs
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&args)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &args, nil
}

func ToJSON(o interface{}) string {
	bs, _ := json.MarshalIndent(o, "", "\t")
	return string(bs)
}
