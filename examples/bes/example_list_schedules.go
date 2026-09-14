/*
 * Copyright 2026 Baidu, Inc.
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

package besexamples

import (
	"fmt"

	bes "github.com/baidubce/bce-sdk-go/services/bes/v3"
)

func ListSchedules(clusterId string) {
	if besClient == nil {
		if err := Init(); err != nil {
			fmt.Printf("Failed to new bes client, err: %v.\n", err)
			return
		}
	}
	response, err := besClient.ListSchedules(&bes.ListSchedulesRequest{
		ClusterId: clusterId,
		PageNo:    1,
		PageSize:  20,
	})
	if err != nil {
		fmt.Printf("Failed to list bes schedules, err: %v.\n", err)
		return
	}
	fmt.Printf("totalCount: %d\n", response.TotalCount)
	for _, item := range response.Schedules {
		fmt.Printf("scheduleId: %s, name: %s, type: %s, cronExpr: %s, enabled: %v\n",
			item.ScheduleId, item.ScheduleName, item.ScheduleType, item.CronExpr, item.Enabled)
		fmt.Printf("  lastTriggerTime: %s, lastStatus: %s, executed: %d, failed: %d\n",
			item.LastTriggerTime, item.LastStatus, item.ExecutedCount, item.FailedCount)
		if item.LastErrorMessage != "" {
			fmt.Printf("  lastErrorMessage: %s\n", item.LastErrorMessage)
		}
	}
}
