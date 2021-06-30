// Copyright 2019 Baidu Inc. All rights reserved
// Use of this source code is governed by a CCE
// license that can be found in the LICENSE file.
/*
modification history
--------------------
2021/06/24 16:00:00, by pansiyuan02@baidu.com, create
*/
package types

import "time"

const (
	TaskTypeInstanceGroupReplicas TaskType = "InstanceGroupReplicas"
)

const (
	TaskPhasePending    TaskPhase = "Pending"
	TaskPhaseProcessing TaskPhase = "Processing"
	TaskPhaseDone       TaskPhase = "Done"
	TaskPhaseAborted    TaskPhase = "Aborted"
)

const (
	TaskProcessPhasePending    TaskProcessPhase = "Pending"
	TaskProcessPhaseProcessing TaskProcessPhase = "Processing"
	TaskProcessPhaseDone       TaskProcessPhase = "Done"
	TaskProcessPhaseAborted    TaskProcessPhase = "Aborted"
)

type TaskType string
type TaskPhase string
type TaskProcessPhase string

type Task struct {
	ID          string     `json:"id"`
	Type        TaskType   `json:"type"`
	Description string     `json:"description"`
	StartTime   time.Time  `json:"startTime"`
	FinishTime  *time.Time `json:"finishTime,omitempty"`
	Phase       TaskPhase  `json:"phase"`

	TaskProcesses []TaskProcess `json:"processes,omitempty"`
	ErrMessage    string        `json:"errMessage,omitempty"`
}

type TaskProcess struct {
	Name       string           `json:"name"`
	Phase      TaskProcessPhase `json:"phase,omitempty"`
	StartTime  *time.Time       `json:"startTime,omitempty"`
	FinishTime *time.Time       `json:"finishTime,omitempty"`

	Metrics      map[string]string `json:"metrics,omitempty"`
	SubProcesses []TaskProcess     `json:"subProcesses,omitempty"`
	ErrMessage   string            `json:"errMessage,omitempty"`
}
