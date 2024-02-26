package model

type AlarmType string

const (
	AlarmTypeInstance AlarmType = "INSTANCE"
	AlarmTypeService  AlarmType = "SERVICE"
)

type AlarmLevel string

const (
	AlarmLevelNotice   AlarmLevel = "NOTICE"
	AlarmLevelWarning  AlarmLevel = "WARNING"
	AlarmLevelCritical AlarmLevel = "CRITICAL"
	AlarmLevelMajor    AlarmLevel = "MAJOR"
	AlarmLevelCustom   AlarmLevel = "CUSTOM"
)

var AlarmLevelMap = map[AlarmLevel]bool{
	AlarmLevelNotice:   true,
	AlarmLevelWarning:  true,
	AlarmLevelCritical: true,
	AlarmLevelMajor:    true,
	AlarmLevelCustom:   true,
}

type TargetType string

const (
	TargetTypeAllInstances   TargetType = "TARGET_TYPE_ALL_INSTANCES"
	TargetTypeInstanceGroup  TargetType = "TARGET_TYPE_INSTANCE_GROUP"
	TargetTypeMultiInstances TargetType = "TARGET_TYPE_MULTI_INSTANCES"
	TargetTypeInstanceTags   TargetType = "TARGET_TYPE_INSTANCE_TAGS"
)

var TargetTypeMap = map[TargetType]bool{
	TargetTypeAllInstances:   true,
	TargetTypeInstanceGroup:  true,
	TargetTypeMultiInstances: true,
	TargetTypeInstanceTags:   true,
}

type KV struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
