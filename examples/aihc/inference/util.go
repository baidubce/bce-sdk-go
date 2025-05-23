package inference

import (
	"encoding/json"
	"fmt"
	"os"

	api "github.com/baidubce/bce-sdk-go/services/aihc/inference/v2"
)

func ReadJson(fileName string) (*api.ServiceConf, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	var serviceConf api.ServiceConf
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&serviceConf)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	return &serviceConf, nil
}

func ToJSON(o interface{}) string {
	bs, _ := json.MarshalIndent(o, "", "\t")
	return string(bs)
}
