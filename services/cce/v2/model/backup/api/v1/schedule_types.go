/*
Copyright 2020 the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"fmt"
	"time"
)

// ScheduleSpec defines the specification for a Velero schedule
type ScheduleSpec struct {
	// Template is the definition of the Backup to be run
	// on the provided schedule
	Template BackupSpec `json:"template"`

	// Schedule is a Cron expression defining when to run
	// the Backup.
	Schedule string `json:"schedule"`

	// UseOwnerReferencesBackup specifies whether to use
	// OwnerReferences on backups created by this Schedule.
	// +optional
	// +nullable
	UseOwnerReferencesInBackup *bool `json:"useOwnerReferencesInBackup,omitempty"`

	// Paused specifies whether the schedule is paused or not
	// +optional
	Paused bool `json:"paused,omitempty"`

	SkipImmediately *bool `json:"skipImmediately,omitempty"`
}

// SchedulePhase is a string representation of the lifecycle phase
// of a Velero schedule
// +kubebuilder:validation:Enum=New;Enabled;FailedValidation
type SchedulePhase string

const (
	// SchedulePhaseNew means the schedule has been created but not
	// yet processed by the ScheduleController
	SchedulePhaseNew SchedulePhase = "New"

	// SchedulePhaseEnabled means the schedule has been validated and
	// will now be triggering backups according to the schedule spec.
	SchedulePhaseEnabled SchedulePhase = "Enabled"

	// SchedulePhaseFailedValidation means the schedule has failed
	// the controller's validations and therefore will not trigger backups.
	SchedulePhaseFailedValidation SchedulePhase = "FailedValidation"
)

// ScheduleStatus captures the current state of a Velero schedule
type ScheduleStatus struct {
	// Phase is the current phase of the Schedule
	// +optional
	Phase SchedulePhase `json:"phase,omitempty"`

	// LastBackup is the last time a Backup was run for this
	// Schedule schedule
	// +optional
	// +nullable
	LastBackup *Time `json:"lastBackup,omitempty"`

	// LastSkipped is the last time a Schedule was skipped
	// +optional
	// +nullable
	LastSkipped *Time `json:"lastSkipped,omitempty"`

	// ValidationErrors is a slice of all validation errors (if
	// applicable)
	// +optional
	ValidationErrors []string `json:"validationErrors,omitempty"`
}

// Schedule is a Velero resource that represents a pre-scheduled or
// periodic Backup that should be run.
type Schedule struct {
	TypeMeta `json:",inline"`

	// +optional
	ObjectMeta `json:"metadata"`

	// +optional
	Spec ScheduleSpec `json:"spec,omitempty"`

	// +optional
	Status ScheduleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// ScheduleList is a list of Schedules.
type ScheduleList struct {
	TypeMeta `json:",inline"`

	// +optional
	ListMeta `json:"metadata,omitempty"`

	Items []Schedule `json:"items"`
}

// TimestampedName returns the default backup name format based on the schedule
func (s *Schedule) TimestampedName(timestamp time.Time) string {
	return fmt.Sprintf("%s-%s", s.Name, timestamp.Format("20060102150405"))
}
