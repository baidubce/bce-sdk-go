/*
 * Copyright 2025 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package bct

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/services/bct/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK        string
	SK        string
	Endpoint  string
	PageSize  int
	StartTime string
	EndTime   string
}

var BCT_CLIENT *Client

var config *Conf

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		fmt.Printf("config json file of ak/sk not given: %+v\n", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	config = &Conf{}
	decoder.Decode(config)
	BCT_CLIENT, _ = NewClientWithEndpoint(config.AK, config.SK, config.Endpoint)
	log.SetLogLevel(log.DEBUG)
}

func TestQueryEventsV2(t *testing.T) {
	start, err := time.Parse(time.RFC3339, config.StartTime)
	if err != nil {
		t.Errorf("parse start time failed: %s", err.Error())
	}
	end, err := time.Parse(time.RFC3339, config.EndTime)
	if err != nil {
		t.Errorf("parse end time failed: %s", err.Error())
	}
	args := &api.QueryEventsV2Request{
		StartTime: start,
		EndTime:   end,
		PageSize:  config.PageSize,
	}
	res, err := BCT_CLIENT.QueryEventsV2(args)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(json.Marshal(res))
}
