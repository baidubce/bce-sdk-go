package api

const (
	HEADER_REGION = "x-Region"

	URI_PREFIX_V2       = "/api/bes/cluster"
	URI_V2              = URI_PREFIX_V2 + "/v2"
	URI_INSTANCE        = URI_V2 + "/instance"
	URI_MIGRATE_V1      = URI_PREFIX_V2 + "/migrate/v1"
	URI_LOG_VIEW        = URI_PREFIX_V2 + "/es_log_view_new"
	URI_PLUGIN          = URI_PREFIX_V2 + "/plugin"
	URI_DEFAULT_PLUGIN  = URI_PREFIX_V2 + "/default_plugin"
	URI_INSPECT         = URI_PREFIX_V2 + "/inspect"
	URI_SCHEDULE        = URI_PREFIX_V2 + "/schedule"
	URI_AUTO_RENEW_RULE = URI_PREFIX_V2 + "/auto_renew_rule"
	URI_RENEW           = URI_PREFIX_V2 + "/renew"
)
