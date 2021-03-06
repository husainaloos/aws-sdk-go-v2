// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AppAttributesKeys string

// Enum values for AppAttributesKeys
const (
	AppAttributesKeysDocumentroot        AppAttributesKeys = "DocumentRoot"
	AppAttributesKeysRailsenv            AppAttributesKeys = "RailsEnv"
	AppAttributesKeysAutobundleondeploy  AppAttributesKeys = "AutoBundleOnDeploy"
	AppAttributesKeysAwsflowrubysettings AppAttributesKeys = "AwsFlowRubySettings"
)

type AppType string

// Enum values for AppType
const (
	AppTypeAwsFlowRuby AppType = "aws-flow-ruby"
	AppTypeJava        AppType = "java"
	AppTypeRails       AppType = "rails"
	AppTypePhp         AppType = "php"
	AppTypeNodejs      AppType = "nodejs"
	AppTypeStatic      AppType = "static"
	AppTypeOther       AppType = "other"
)

type Architecture string

// Enum values for Architecture
const (
	ArchitectureX86_64 Architecture = "x86_64"
	ArchitectureI386   Architecture = "i386"
)

type AutoScalingType string

// Enum values for AutoScalingType
const (
	AutoScalingTypeLoad  AutoScalingType = "load"
	AutoScalingTypeTimer AutoScalingType = "timer"
)

type CloudWatchLogsEncoding string

// Enum values for CloudWatchLogsEncoding
const (
	CloudWatchLogsEncodingAscii           CloudWatchLogsEncoding = "ascii"
	CloudWatchLogsEncodingBig5            CloudWatchLogsEncoding = "big5"
	CloudWatchLogsEncodingBig5hkscs       CloudWatchLogsEncoding = "big5hkscs"
	CloudWatchLogsEncodingCp037           CloudWatchLogsEncoding = "cp037"
	CloudWatchLogsEncodingCp424           CloudWatchLogsEncoding = "cp424"
	CloudWatchLogsEncodingCp437           CloudWatchLogsEncoding = "cp437"
	CloudWatchLogsEncodingCp500           CloudWatchLogsEncoding = "cp500"
	CloudWatchLogsEncodingCp720           CloudWatchLogsEncoding = "cp720"
	CloudWatchLogsEncodingCp737           CloudWatchLogsEncoding = "cp737"
	CloudWatchLogsEncodingCp775           CloudWatchLogsEncoding = "cp775"
	CloudWatchLogsEncodingCp850           CloudWatchLogsEncoding = "cp850"
	CloudWatchLogsEncodingCp852           CloudWatchLogsEncoding = "cp852"
	CloudWatchLogsEncodingCp855           CloudWatchLogsEncoding = "cp855"
	CloudWatchLogsEncodingCp856           CloudWatchLogsEncoding = "cp856"
	CloudWatchLogsEncodingCp857           CloudWatchLogsEncoding = "cp857"
	CloudWatchLogsEncodingCp858           CloudWatchLogsEncoding = "cp858"
	CloudWatchLogsEncodingCp860           CloudWatchLogsEncoding = "cp860"
	CloudWatchLogsEncodingCp861           CloudWatchLogsEncoding = "cp861"
	CloudWatchLogsEncodingCp862           CloudWatchLogsEncoding = "cp862"
	CloudWatchLogsEncodingCp863           CloudWatchLogsEncoding = "cp863"
	CloudWatchLogsEncodingCp864           CloudWatchLogsEncoding = "cp864"
	CloudWatchLogsEncodingCp865           CloudWatchLogsEncoding = "cp865"
	CloudWatchLogsEncodingCp866           CloudWatchLogsEncoding = "cp866"
	CloudWatchLogsEncodingCp869           CloudWatchLogsEncoding = "cp869"
	CloudWatchLogsEncodingCp874           CloudWatchLogsEncoding = "cp874"
	CloudWatchLogsEncodingCp875           CloudWatchLogsEncoding = "cp875"
	CloudWatchLogsEncodingCp932           CloudWatchLogsEncoding = "cp932"
	CloudWatchLogsEncodingCp949           CloudWatchLogsEncoding = "cp949"
	CloudWatchLogsEncodingCp950           CloudWatchLogsEncoding = "cp950"
	CloudWatchLogsEncodingCp1006          CloudWatchLogsEncoding = "cp1006"
	CloudWatchLogsEncodingCp1026          CloudWatchLogsEncoding = "cp1026"
	CloudWatchLogsEncodingCp1140          CloudWatchLogsEncoding = "cp1140"
	CloudWatchLogsEncodingCp1250          CloudWatchLogsEncoding = "cp1250"
	CloudWatchLogsEncodingCp1251          CloudWatchLogsEncoding = "cp1251"
	CloudWatchLogsEncodingCp1252          CloudWatchLogsEncoding = "cp1252"
	CloudWatchLogsEncodingCp1253          CloudWatchLogsEncoding = "cp1253"
	CloudWatchLogsEncodingCp1254          CloudWatchLogsEncoding = "cp1254"
	CloudWatchLogsEncodingCp1255          CloudWatchLogsEncoding = "cp1255"
	CloudWatchLogsEncodingCp1256          CloudWatchLogsEncoding = "cp1256"
	CloudWatchLogsEncodingCp1257          CloudWatchLogsEncoding = "cp1257"
	CloudWatchLogsEncodingCp1258          CloudWatchLogsEncoding = "cp1258"
	CloudWatchLogsEncodingEuc_jp          CloudWatchLogsEncoding = "euc_jp"
	CloudWatchLogsEncodingEuc_jis_2004    CloudWatchLogsEncoding = "euc_jis_2004"
	CloudWatchLogsEncodingEuc_jisx0213    CloudWatchLogsEncoding = "euc_jisx0213"
	CloudWatchLogsEncodingEuc_kr          CloudWatchLogsEncoding = "euc_kr"
	CloudWatchLogsEncodingGb2312          CloudWatchLogsEncoding = "gb2312"
	CloudWatchLogsEncodingGbk             CloudWatchLogsEncoding = "gbk"
	CloudWatchLogsEncodingGb18030         CloudWatchLogsEncoding = "gb18030"
	CloudWatchLogsEncodingHz              CloudWatchLogsEncoding = "hz"
	CloudWatchLogsEncodingIso2022_jp      CloudWatchLogsEncoding = "iso2022_jp"
	CloudWatchLogsEncodingIso2022_jp_1    CloudWatchLogsEncoding = "iso2022_jp_1"
	CloudWatchLogsEncodingIso2022_jp_2    CloudWatchLogsEncoding = "iso2022_jp_2"
	CloudWatchLogsEncodingIso2022_jp_2004 CloudWatchLogsEncoding = "iso2022_jp_2004"
	CloudWatchLogsEncodingIso2022_jp_3    CloudWatchLogsEncoding = "iso2022_jp_3"
	CloudWatchLogsEncodingIso2022_jp_ext  CloudWatchLogsEncoding = "iso2022_jp_ext"
	CloudWatchLogsEncodingIso2022_kr      CloudWatchLogsEncoding = "iso2022_kr"
	CloudWatchLogsEncodingLatin_1         CloudWatchLogsEncoding = "latin_1"
	CloudWatchLogsEncodingIso8859_2       CloudWatchLogsEncoding = "iso8859_2"
	CloudWatchLogsEncodingIso8859_3       CloudWatchLogsEncoding = "iso8859_3"
	CloudWatchLogsEncodingIso8859_4       CloudWatchLogsEncoding = "iso8859_4"
	CloudWatchLogsEncodingIso8859_5       CloudWatchLogsEncoding = "iso8859_5"
	CloudWatchLogsEncodingIso8859_6       CloudWatchLogsEncoding = "iso8859_6"
	CloudWatchLogsEncodingIso8859_7       CloudWatchLogsEncoding = "iso8859_7"
	CloudWatchLogsEncodingIso8859_8       CloudWatchLogsEncoding = "iso8859_8"
	CloudWatchLogsEncodingIso8859_9       CloudWatchLogsEncoding = "iso8859_9"
	CloudWatchLogsEncodingIso8859_10      CloudWatchLogsEncoding = "iso8859_10"
	CloudWatchLogsEncodingIso8859_13      CloudWatchLogsEncoding = "iso8859_13"
	CloudWatchLogsEncodingIso8859_14      CloudWatchLogsEncoding = "iso8859_14"
	CloudWatchLogsEncodingIso8859_15      CloudWatchLogsEncoding = "iso8859_15"
	CloudWatchLogsEncodingIso8859_16      CloudWatchLogsEncoding = "iso8859_16"
	CloudWatchLogsEncodingJohab           CloudWatchLogsEncoding = "johab"
	CloudWatchLogsEncodingKoi8_r          CloudWatchLogsEncoding = "koi8_r"
	CloudWatchLogsEncodingKoi8_u          CloudWatchLogsEncoding = "koi8_u"
	CloudWatchLogsEncodingMac_cyrillic    CloudWatchLogsEncoding = "mac_cyrillic"
	CloudWatchLogsEncodingMac_greek       CloudWatchLogsEncoding = "mac_greek"
	CloudWatchLogsEncodingMac_iceland     CloudWatchLogsEncoding = "mac_iceland"
	CloudWatchLogsEncodingMac_latin2      CloudWatchLogsEncoding = "mac_latin2"
	CloudWatchLogsEncodingMac_roman       CloudWatchLogsEncoding = "mac_roman"
	CloudWatchLogsEncodingMac_turkish     CloudWatchLogsEncoding = "mac_turkish"
	CloudWatchLogsEncodingPtcp154         CloudWatchLogsEncoding = "ptcp154"
	CloudWatchLogsEncodingShift_jis       CloudWatchLogsEncoding = "shift_jis"
	CloudWatchLogsEncodingShift_jis_2004  CloudWatchLogsEncoding = "shift_jis_2004"
	CloudWatchLogsEncodingShift_jisx0213  CloudWatchLogsEncoding = "shift_jisx0213"
	CloudWatchLogsEncodingUtf_32          CloudWatchLogsEncoding = "utf_32"
	CloudWatchLogsEncodingUtf_32_be       CloudWatchLogsEncoding = "utf_32_be"
	CloudWatchLogsEncodingUtf_32_le       CloudWatchLogsEncoding = "utf_32_le"
	CloudWatchLogsEncodingUtf_16          CloudWatchLogsEncoding = "utf_16"
	CloudWatchLogsEncodingUtf_16_be       CloudWatchLogsEncoding = "utf_16_be"
	CloudWatchLogsEncodingUtf_16_le       CloudWatchLogsEncoding = "utf_16_le"
	CloudWatchLogsEncodingUtf_7           CloudWatchLogsEncoding = "utf_7"
	CloudWatchLogsEncodingUtf_8           CloudWatchLogsEncoding = "utf_8"
	CloudWatchLogsEncodingUtf_8_sig       CloudWatchLogsEncoding = "utf_8_sig"
)

type CloudWatchLogsInitialPosition string

// Enum values for CloudWatchLogsInitialPosition
const (
	CloudWatchLogsInitialPositionStart_of_file CloudWatchLogsInitialPosition = "start_of_file"
	CloudWatchLogsInitialPositionEnd_of_file   CloudWatchLogsInitialPosition = "end_of_file"
)

type CloudWatchLogsTimeZone string

// Enum values for CloudWatchLogsTimeZone
const (
	CloudWatchLogsTimeZoneLocal CloudWatchLogsTimeZone = "LOCAL"
	CloudWatchLogsTimeZoneUtc   CloudWatchLogsTimeZone = "UTC"
)

type DeploymentCommandName string

// Enum values for DeploymentCommandName
const (
	DeploymentCommandNameInstall_dependencies    DeploymentCommandName = "install_dependencies"
	DeploymentCommandNameUpdate_dependencies     DeploymentCommandName = "update_dependencies"
	DeploymentCommandNameUpdate_custom_cookbooks DeploymentCommandName = "update_custom_cookbooks"
	DeploymentCommandNameExecute_recipes         DeploymentCommandName = "execute_recipes"
	DeploymentCommandNameConfigure               DeploymentCommandName = "configure"
	DeploymentCommandNameSetup                   DeploymentCommandName = "setup"
	DeploymentCommandNameDeploy                  DeploymentCommandName = "deploy"
	DeploymentCommandNameRollback                DeploymentCommandName = "rollback"
	DeploymentCommandNameStart                   DeploymentCommandName = "start"
	DeploymentCommandNameStop                    DeploymentCommandName = "stop"
	DeploymentCommandNameRestart                 DeploymentCommandName = "restart"
	DeploymentCommandNameUndeploy                DeploymentCommandName = "undeploy"
)

type LayerAttributesKeys string

// Enum values for LayerAttributesKeys
const (
	LayerAttributesKeysEcsclusterarn               LayerAttributesKeys = "EcsClusterArn"
	LayerAttributesKeysEnablehaproxystats          LayerAttributesKeys = "EnableHaproxyStats"
	LayerAttributesKeysHaproxystatsurl             LayerAttributesKeys = "HaproxyStatsUrl"
	LayerAttributesKeysHaproxystatsuser            LayerAttributesKeys = "HaproxyStatsUser"
	LayerAttributesKeysHaproxystatspassword        LayerAttributesKeys = "HaproxyStatsPassword"
	LayerAttributesKeysHaproxyhealthcheckurl       LayerAttributesKeys = "HaproxyHealthCheckUrl"
	LayerAttributesKeysHaproxyhealthcheckmethod    LayerAttributesKeys = "HaproxyHealthCheckMethod"
	LayerAttributesKeysMysqlrootpassword           LayerAttributesKeys = "MysqlRootPassword"
	LayerAttributesKeysMysqlrootpasswordubiquitous LayerAttributesKeys = "MysqlRootPasswordUbiquitous"
	LayerAttributesKeysGangliaurl                  LayerAttributesKeys = "GangliaUrl"
	LayerAttributesKeysGangliauser                 LayerAttributesKeys = "GangliaUser"
	LayerAttributesKeysGangliapassword             LayerAttributesKeys = "GangliaPassword"
	LayerAttributesKeysMemcachedmemory             LayerAttributesKeys = "MemcachedMemory"
	LayerAttributesKeysNodejsversion               LayerAttributesKeys = "NodejsVersion"
	LayerAttributesKeysRubyversion                 LayerAttributesKeys = "RubyVersion"
	LayerAttributesKeysRubygemsversion             LayerAttributesKeys = "RubygemsVersion"
	LayerAttributesKeysManagebundler               LayerAttributesKeys = "ManageBundler"
	LayerAttributesKeysBundlerversion              LayerAttributesKeys = "BundlerVersion"
	LayerAttributesKeysRailsstack                  LayerAttributesKeys = "RailsStack"
	LayerAttributesKeysPassengerversion            LayerAttributesKeys = "PassengerVersion"
	LayerAttributesKeysJvm                         LayerAttributesKeys = "Jvm"
	LayerAttributesKeysJvmversion                  LayerAttributesKeys = "JvmVersion"
	LayerAttributesKeysJvmoptions                  LayerAttributesKeys = "JvmOptions"
	LayerAttributesKeysJavaappserver               LayerAttributesKeys = "JavaAppServer"
	LayerAttributesKeysJavaappserverversion        LayerAttributesKeys = "JavaAppServerVersion"
)

type LayerType string

// Enum values for LayerType
const (
	LayerTypeAwsFlowRuby      LayerType = "aws-flow-ruby"
	LayerTypeEcsCluster       LayerType = "ecs-cluster"
	LayerTypeJavaApp          LayerType = "java-app"
	LayerTypeLb               LayerType = "lb"
	LayerTypeWeb              LayerType = "web"
	LayerTypePhpApp           LayerType = "php-app"
	LayerTypeRailsApp         LayerType = "rails-app"
	LayerTypeNodejsApp        LayerType = "nodejs-app"
	LayerTypeMemcached        LayerType = "memcached"
	LayerTypeDbMaster         LayerType = "db-master"
	LayerTypeMonitoringMaster LayerType = "monitoring-master"
	LayerTypeCustom           LayerType = "custom"
)

type RootDeviceType string

// Enum values for RootDeviceType
const (
	RootDeviceTypeEbs           RootDeviceType = "ebs"
	RootDeviceTypeInstanceStore RootDeviceType = "instance-store"
)

type SourceType string

// Enum values for SourceType
const (
	SourceTypeGit     SourceType = "git"
	SourceTypeSvn     SourceType = "svn"
	SourceTypeArchive SourceType = "archive"
	SourceTypeS3      SourceType = "s3"
)

type StackAttributesKeys string

// Enum values for StackAttributesKeys
const (
	StackAttributesKeysColor StackAttributesKeys = "Color"
)

type VirtualizationType string

// Enum values for VirtualizationType
const (
	VirtualizationTypeParavirtual VirtualizationType = "paravirtual"
	VirtualizationTypeHvm         VirtualizationType = "hvm"
)

type VolumeType string

// Enum values for VolumeType
const (
	VolumeTypeGp2      VolumeType = "gp2"
	VolumeTypeIo1      VolumeType = "io1"
	VolumeTypeStandard VolumeType = "standard"
)
