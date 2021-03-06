// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CloudWatchEventSource string

// Enum values for CloudWatchEventSource
const (
	CloudWatchEventSourceEc2         CloudWatchEventSource = "EC2"
	CloudWatchEventSourceCode_deploy CloudWatchEventSource = "CODE_DEPLOY"
	CloudWatchEventSourceHealth      CloudWatchEventSource = "HEALTH"
)

type ConfigurationEventResourceType string

// Enum values for ConfigurationEventResourceType
const (
	ConfigurationEventResourceTypeCloudwatch_alarm ConfigurationEventResourceType = "CLOUDWATCH_ALARM"
	ConfigurationEventResourceTypeCloudformation   ConfigurationEventResourceType = "CLOUDFORMATION"
	ConfigurationEventResourceTypeSsm_association  ConfigurationEventResourceType = "SSM_ASSOCIATION"
)

type ConfigurationEventStatus string

// Enum values for ConfigurationEventStatus
const (
	ConfigurationEventStatusInfo  ConfigurationEventStatus = "INFO"
	ConfigurationEventStatusWarn  ConfigurationEventStatus = "WARN"
	ConfigurationEventStatusError ConfigurationEventStatus = "ERROR"
)

type FeedbackKey string

// Enum values for FeedbackKey
const (
	FeedbackKeyInsights_feedback FeedbackKey = "INSIGHTS_FEEDBACK"
)

type FeedbackValue string

// Enum values for FeedbackValue
const (
	FeedbackValueNot_specified FeedbackValue = "NOT_SPECIFIED"
	FeedbackValueUseful        FeedbackValue = "USEFUL"
	FeedbackValueNot_useful    FeedbackValue = "NOT_USEFUL"
)

type LogFilter string

// Enum values for LogFilter
const (
	LogFilterError LogFilter = "ERROR"
	LogFilterWarn  LogFilter = "WARN"
	LogFilterInfo  LogFilter = "INFO"
)

type SeverityLevel string

// Enum values for SeverityLevel
const (
	SeverityLevelLow    SeverityLevel = "Low"
	SeverityLevelMedium SeverityLevel = "Medium"
	SeverityLevelHigh   SeverityLevel = "High"
)

type Status string

// Enum values for Status
const (
	StatusIgnore   Status = "IGNORE"
	StatusResolved Status = "RESOLVED"
	StatusPending  Status = "PENDING"
)

type Tier string

// Enum values for Tier
const (
	TierDefault        Tier = "DEFAULT"
	TierDot_net_core   Tier = "DOT_NET_CORE"
	TierDot_net_worker Tier = "DOT_NET_WORKER"
	TierDot_net_web    Tier = "DOT_NET_WEB"
	TierSql_server     Tier = "SQL_SERVER"
)
