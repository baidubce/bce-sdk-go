package dev

type Interface interface {
	CreateDevInstance(*CreateDevInstanceArgs) (*CreateDevInstanceResult, error)
	ListDevInstance(*ListDevInstanceArgs) (*ListDevInstanceResult, error)
	QueryDevInstanceDetail(*QueryDevInstanceDetailArgs) (*QueryDevInstanceDetailResult, error)
	UpdateDevInstance(*CreateDevInstanceArgs) (*CreateDevInstanceResult, error)
	StartDevInstance(*StartDevInstanceArgs) (*StartDevInstanceResult, error)
	StopDevInstance(*StopDevInstanceArgs) (*StopDevInstanceResult, error)
	DeleteDevInstance(*DeleteDevInstanceArgs) (*DeleteDevInstanceResult, error)
	TimedStopDevInstance(*TimedStopDevInstanceArgs) (*TimedStopDevInstanceResult, error)
	CreateDevInstanceImagePackJob(*CreateDevInstanceImagePackJobArgs) (*CreateDevInstanceImagePackJobResult, error)
	DevInstanceImagePackJobDetail(*DevInstanceImagePackJobDetailArgs) (*DevInstanceImagePackJobDetailResult, error)
	ListDevInstanceEvent(*ListDevInstanceEventArgs) (*ListDevInstanceEventResult, error)
}
