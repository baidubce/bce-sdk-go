package inference

import (
	"github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
)

type Interface interface {
	CreateApp(args *api.CreateAppArgs, region string, extraInfo map[string]string) (*api.CreateAppResult, error)
	ListApp(args *api.ListAppArgs, region string, extraInfo map[string]string) (*api.ListAppResult, error)
	ListAppStats(args *api.ListAppStatsArgs, region string) (*api.ListAppStatsResult, error)
	AppDetails(args *api.AppDetailsArgs, region string) (*api.AppDetailsResult, error)
	UpdateApp(args *api.UpdateAppArgs, region string) (*api.UpdateAppResult, error)
	ScaleApp(args *api.ScaleAppArgs, region string) (*api.ScaleAppResult, error)
	PubAccess(args *api.PubAccessArgs, region string) (*api.PubAccessResult, error)
	ListChange(args *api.ListChangeArgs, region string) (*api.ListChangeResult, error)
	ChangeDetail(args *api.ChangeDetailArgs, region string) (*api.ChangeDetailResult, error)
	DeleteApp(args *api.DeleteAppArgs, region string) (*api.DeleteAppResult, error)
	ListPod(args *api.ListPodArgs, region string) (*api.ListPodResult, error)
	BlockPod(args *api.BlockPodArgs, region string) (*api.BlockPodResult, error)
	DeletePod(args *api.DeletePodArgs, region string) (*api.DeletePodResult, error)
	ListBriefResPool(args *api.ListBriefResPoolArgs, region string) (*api.ListBriefResPoolResult, error)
	ResPoolDetail(args *api.ResPoolDetailArgs, region string) (*api.ResPoolDetailResult, error)
}
