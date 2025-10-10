package bct

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/bct"
	"github.com/baidubce/bce-sdk-go/services/bct/api"
	"time"
)

func ExampleQueryEventsV2() {
	// init ak, sk, query
	ak, sk := "ak", "sk"
	client, _ := bct.NewClient(ak, sk)
	start, _ := time.Parse(time.RFC3339, "2025-09-02T16:00:00Z")
	end, _ := time.Parse(time.RFC3339, "2025-09-03T16:00:00Z")
	filters := api.FieldFilter{
		Field: "eventSource",
		Value: "iam",
	}
	args := &api.QueryEventsV2Request{
		StartTime: start,
		EndTime:   end,
		PageSize:  10,
		Filters:   []api.FieldFilter{filters},
	}

	res, err := client.QueryEventsV2(args)
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(res)
	fmt.Println(string(jsonBytes))
}
