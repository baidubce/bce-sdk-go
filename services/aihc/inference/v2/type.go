package api

type Interface interface {
	CreateService(args *ServiceConf, clientToken string) (*CreateServiceResult, error)
	ListService(args *ListServiceArgs) (*ListServiceResult, error)
	ListServiceStats(args *ListServiceStatsArgs) (*ListServiceStatsResult, error)
	ServiceDetails(args *ServiceDetailsArgs) (*ServiceDetailsResult, error)
	UpdateService(args *UpdateServiceArgs) (*UpdateServiceResult, error)
	ScaleService(args *ScaleServiceArgs) (*ScaleServiceResult, error)
	PubAccess(args *PubAccessArgs) (*PubAccessResult, error)
	ListChange(args *ListChangeArgs) (*ListChangeResult, error)
	ChangeDetail(args *ChangeDetailArgs) (*ChangeDetailResult, error)
	DeleteService(args *DeleteServiceArgs) (*DeleteServiceResult, error)
	ListPod(args *ListPodArgs) (*ListPodResult, error)
	BlockPod(args *BlockPodArgs) (*BlockPodResult, error)
	DeletePod(args *DeletePodArgs) (*DeletePodResult, error)
	ListPodGroups(args *ListPodGroupsArgs) (*ListPodGroupsResult, error)
}
