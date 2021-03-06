// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdTriggersElement string

// Enum values for AdTriggersElement
const (
	AdTriggersElementSplice_insert                             AdTriggersElement = "SPLICE_INSERT"
	AdTriggersElementBreak                                     AdTriggersElement = "BREAK"
	AdTriggersElementProvider_advertisement                    AdTriggersElement = "PROVIDER_ADVERTISEMENT"
	AdTriggersElementDistributor_advertisement                 AdTriggersElement = "DISTRIBUTOR_ADVERTISEMENT"
	AdTriggersElementProvider_placement_opportunity            AdTriggersElement = "PROVIDER_PLACEMENT_OPPORTUNITY"
	AdTriggersElementDistributor_placement_opportunity         AdTriggersElement = "DISTRIBUTOR_PLACEMENT_OPPORTUNITY"
	AdTriggersElementProvider_overlay_placement_opportunity    AdTriggersElement = "PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY"
	AdTriggersElementDistributor_overlay_placement_opportunity AdTriggersElement = "DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY"
)

type PeriodTriggersElement string

// Enum values for PeriodTriggersElement
const (
	PeriodTriggersElementAds PeriodTriggersElement = "ADS"
)

type AdMarkers string

// Enum values for AdMarkers
const (
	AdMarkersNone            AdMarkers = "NONE"
	AdMarkersScte35_enhanced AdMarkers = "SCTE35_ENHANCED"
	AdMarkersPassthrough     AdMarkers = "PASSTHROUGH"
	AdMarkersDaterange       AdMarkers = "DATERANGE"
)

type AdsOnDeliveryRestrictions string

// Enum values for AdsOnDeliveryRestrictions
const (
	AdsOnDeliveryRestrictionsNone         AdsOnDeliveryRestrictions = "NONE"
	AdsOnDeliveryRestrictionsRestricted   AdsOnDeliveryRestrictions = "RESTRICTED"
	AdsOnDeliveryRestrictionsUnrestricted AdsOnDeliveryRestrictions = "UNRESTRICTED"
	AdsOnDeliveryRestrictionsBoth         AdsOnDeliveryRestrictions = "BOTH"
)

type EncryptionMethod string

// Enum values for EncryptionMethod
const (
	EncryptionMethodAes_128    EncryptionMethod = "AES_128"
	EncryptionMethodSample_aes EncryptionMethod = "SAMPLE_AES"
)

type ManifestLayout string

// Enum values for ManifestLayout
const (
	ManifestLayoutFull    ManifestLayout = "FULL"
	ManifestLayoutCompact ManifestLayout = "COMPACT"
)

type Origination string

// Enum values for Origination
const (
	OriginationAllow Origination = "ALLOW"
	OriginationDeny  Origination = "DENY"
)

type PlaylistType string

// Enum values for PlaylistType
const (
	PlaylistTypeNone  PlaylistType = "NONE"
	PlaylistTypeEvent PlaylistType = "EVENT"
	PlaylistTypeVod   PlaylistType = "VOD"
)

type Profile string

// Enum values for Profile
const (
	ProfileNone      Profile = "NONE"
	ProfileHbbtv_1_5 Profile = "HBBTV_1_5"
)

type SegmentTemplateFormat string

// Enum values for SegmentTemplateFormat
const (
	SegmentTemplateFormatNumber_with_timeline SegmentTemplateFormat = "NUMBER_WITH_TIMELINE"
	SegmentTemplateFormatTime_with_timeline   SegmentTemplateFormat = "TIME_WITH_TIMELINE"
	SegmentTemplateFormatNumber_with_duration SegmentTemplateFormat = "NUMBER_WITH_DURATION"
)

type Status string

// Enum values for Status
const (
	StatusIn_progress Status = "IN_PROGRESS"
	StatusSucceeded   Status = "SUCCEEDED"
	StatusFailed      Status = "FAILED"
)

type StreamOrder string

// Enum values for StreamOrder
const (
	StreamOrderOriginal                 StreamOrder = "ORIGINAL"
	StreamOrderVideo_bitrate_ascending  StreamOrder = "VIDEO_BITRATE_ASCENDING"
	StreamOrderVideo_bitrate_descending StreamOrder = "VIDEO_BITRATE_DESCENDING"
)
