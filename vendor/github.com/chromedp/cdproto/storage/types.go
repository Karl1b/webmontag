package storage

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"fmt"

	"github.com/chromedp/cdproto/cdp"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// SerializedStorageKey [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-SerializedStorageKey
type SerializedStorageKey string

// String returns the SerializedStorageKey as string value.
func (t SerializedStorageKey) String() string {
	return string(t)
}

// Type enum of possible storage types.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-StorageType
type Type string

// String returns the Type as string value.
func (t Type) String() string {
	return string(t)
}

// Type values.
const (
	TypeAppcache       Type = "appcache"
	TypeCookies        Type = "cookies"
	TypeFileSystems    Type = "file_systems"
	TypeIndexeddb      Type = "indexeddb"
	TypeLocalStorage   Type = "local_storage"
	TypeShaderCache    Type = "shader_cache"
	TypeWebsql         Type = "websql"
	TypeServiceWorkers Type = "service_workers"
	TypeCacheStorage   Type = "cache_storage"
	TypeInterestGroups Type = "interest_groups"
	TypeSharedStorage  Type = "shared_storage"
	TypeStorageBuckets Type = "storage_buckets"
	TypeAll            Type = "all"
	TypeOther          Type = "other"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t Type) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t Type) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *Type) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch Type(v) {
	case TypeAppcache:
		*t = TypeAppcache
	case TypeCookies:
		*t = TypeCookies
	case TypeFileSystems:
		*t = TypeFileSystems
	case TypeIndexeddb:
		*t = TypeIndexeddb
	case TypeLocalStorage:
		*t = TypeLocalStorage
	case TypeShaderCache:
		*t = TypeShaderCache
	case TypeWebsql:
		*t = TypeWebsql
	case TypeServiceWorkers:
		*t = TypeServiceWorkers
	case TypeCacheStorage:
		*t = TypeCacheStorage
	case TypeInterestGroups:
		*t = TypeInterestGroups
	case TypeSharedStorage:
		*t = TypeSharedStorage
	case TypeStorageBuckets:
		*t = TypeStorageBuckets
	case TypeAll:
		*t = TypeAll
	case TypeOther:
		*t = TypeOther

	default:
		in.AddError(fmt.Errorf("unknown Type value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *Type) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// UsageForType usage for a storage type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-UsageForType
type UsageForType struct {
	StorageType Type    `json:"storageType"` // Name of storage type.
	Usage       float64 `json:"usage"`       // Storage usage (bytes).
}

// TrustTokens pair of issuer origin and number of available (signed, but not
// used) Trust Tokens from that issuer.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-TrustTokens
type TrustTokens struct {
	IssuerOrigin string  `json:"issuerOrigin"`
	Count        float64 `json:"count"`
}

// InterestGroupAccessType enum of interest group access types.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-InterestGroupAccessType
type InterestGroupAccessType string

// String returns the InterestGroupAccessType as string value.
func (t InterestGroupAccessType) String() string {
	return string(t)
}

// InterestGroupAccessType values.
const (
	InterestGroupAccessTypeJoin             InterestGroupAccessType = "join"
	InterestGroupAccessTypeLeave            InterestGroupAccessType = "leave"
	InterestGroupAccessTypeUpdate           InterestGroupAccessType = "update"
	InterestGroupAccessTypeLoaded           InterestGroupAccessType = "loaded"
	InterestGroupAccessTypeBid              InterestGroupAccessType = "bid"
	InterestGroupAccessTypeWin              InterestGroupAccessType = "win"
	InterestGroupAccessTypeAdditionalBid    InterestGroupAccessType = "additionalBid"
	InterestGroupAccessTypeAdditionalBidWin InterestGroupAccessType = "additionalBidWin"
	InterestGroupAccessTypeClear            InterestGroupAccessType = "clear"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t InterestGroupAccessType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t InterestGroupAccessType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *InterestGroupAccessType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch InterestGroupAccessType(v) {
	case InterestGroupAccessTypeJoin:
		*t = InterestGroupAccessTypeJoin
	case InterestGroupAccessTypeLeave:
		*t = InterestGroupAccessTypeLeave
	case InterestGroupAccessTypeUpdate:
		*t = InterestGroupAccessTypeUpdate
	case InterestGroupAccessTypeLoaded:
		*t = InterestGroupAccessTypeLoaded
	case InterestGroupAccessTypeBid:
		*t = InterestGroupAccessTypeBid
	case InterestGroupAccessTypeWin:
		*t = InterestGroupAccessTypeWin
	case InterestGroupAccessTypeAdditionalBid:
		*t = InterestGroupAccessTypeAdditionalBid
	case InterestGroupAccessTypeAdditionalBidWin:
		*t = InterestGroupAccessTypeAdditionalBidWin
	case InterestGroupAccessTypeClear:
		*t = InterestGroupAccessTypeClear

	default:
		in.AddError(fmt.Errorf("unknown InterestGroupAccessType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *InterestGroupAccessType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// InterestGroupAd ad advertising element inside an interest group.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-InterestGroupAd
type InterestGroupAd struct {
	RenderURL string `json:"renderURL"`
	Metadata  string `json:"metadata,omitempty"`
}

// InterestGroupDetails the full details of an interest group.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-InterestGroupDetails
type InterestGroupDetails struct {
	OwnerOrigin               string              `json:"ownerOrigin"`
	Name                      string              `json:"name"`
	ExpirationTime            *cdp.TimeSinceEpoch `json:"expirationTime"`
	JoiningOrigin             string              `json:"joiningOrigin"`
	BiddingLogicURL           string              `json:"biddingLogicURL,omitempty"`
	BiddingWasmHelperURL      string              `json:"biddingWasmHelperURL,omitempty"`
	UpdateURL                 string              `json:"updateURL,omitempty"`
	TrustedBiddingSignalsURL  string              `json:"trustedBiddingSignalsURL,omitempty"`
	TrustedBiddingSignalsKeys []string            `json:"trustedBiddingSignalsKeys"`
	UserBiddingSignals        string              `json:"userBiddingSignals,omitempty"`
	Ads                       []*InterestGroupAd  `json:"ads"`
	AdComponents              []*InterestGroupAd  `json:"adComponents"`
}

// SharedStorageAccessType enum of shared storage access types.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-SharedStorageAccessType
type SharedStorageAccessType string

// String returns the SharedStorageAccessType as string value.
func (t SharedStorageAccessType) String() string {
	return string(t)
}

// SharedStorageAccessType values.
const (
	SharedStorageAccessTypeDocumentAddModule      SharedStorageAccessType = "documentAddModule"
	SharedStorageAccessTypeDocumentSelectURL      SharedStorageAccessType = "documentSelectURL"
	SharedStorageAccessTypeDocumentRun            SharedStorageAccessType = "documentRun"
	SharedStorageAccessTypeDocumentSet            SharedStorageAccessType = "documentSet"
	SharedStorageAccessTypeDocumentAppend         SharedStorageAccessType = "documentAppend"
	SharedStorageAccessTypeDocumentDelete         SharedStorageAccessType = "documentDelete"
	SharedStorageAccessTypeDocumentClear          SharedStorageAccessType = "documentClear"
	SharedStorageAccessTypeWorkletSet             SharedStorageAccessType = "workletSet"
	SharedStorageAccessTypeWorkletAppend          SharedStorageAccessType = "workletAppend"
	SharedStorageAccessTypeWorkletDelete          SharedStorageAccessType = "workletDelete"
	SharedStorageAccessTypeWorkletClear           SharedStorageAccessType = "workletClear"
	SharedStorageAccessTypeWorkletGet             SharedStorageAccessType = "workletGet"
	SharedStorageAccessTypeWorkletKeys            SharedStorageAccessType = "workletKeys"
	SharedStorageAccessTypeWorkletEntries         SharedStorageAccessType = "workletEntries"
	SharedStorageAccessTypeWorkletLength          SharedStorageAccessType = "workletLength"
	SharedStorageAccessTypeWorkletRemainingBudget SharedStorageAccessType = "workletRemainingBudget"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SharedStorageAccessType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SharedStorageAccessType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SharedStorageAccessType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch SharedStorageAccessType(v) {
	case SharedStorageAccessTypeDocumentAddModule:
		*t = SharedStorageAccessTypeDocumentAddModule
	case SharedStorageAccessTypeDocumentSelectURL:
		*t = SharedStorageAccessTypeDocumentSelectURL
	case SharedStorageAccessTypeDocumentRun:
		*t = SharedStorageAccessTypeDocumentRun
	case SharedStorageAccessTypeDocumentSet:
		*t = SharedStorageAccessTypeDocumentSet
	case SharedStorageAccessTypeDocumentAppend:
		*t = SharedStorageAccessTypeDocumentAppend
	case SharedStorageAccessTypeDocumentDelete:
		*t = SharedStorageAccessTypeDocumentDelete
	case SharedStorageAccessTypeDocumentClear:
		*t = SharedStorageAccessTypeDocumentClear
	case SharedStorageAccessTypeWorkletSet:
		*t = SharedStorageAccessTypeWorkletSet
	case SharedStorageAccessTypeWorkletAppend:
		*t = SharedStorageAccessTypeWorkletAppend
	case SharedStorageAccessTypeWorkletDelete:
		*t = SharedStorageAccessTypeWorkletDelete
	case SharedStorageAccessTypeWorkletClear:
		*t = SharedStorageAccessTypeWorkletClear
	case SharedStorageAccessTypeWorkletGet:
		*t = SharedStorageAccessTypeWorkletGet
	case SharedStorageAccessTypeWorkletKeys:
		*t = SharedStorageAccessTypeWorkletKeys
	case SharedStorageAccessTypeWorkletEntries:
		*t = SharedStorageAccessTypeWorkletEntries
	case SharedStorageAccessTypeWorkletLength:
		*t = SharedStorageAccessTypeWorkletLength
	case SharedStorageAccessTypeWorkletRemainingBudget:
		*t = SharedStorageAccessTypeWorkletRemainingBudget

	default:
		in.AddError(fmt.Errorf("unknown SharedStorageAccessType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SharedStorageAccessType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SharedStorageEntry struct for a single key-value pair in an origin's
// shared storage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-SharedStorageEntry
type SharedStorageEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SharedStorageMetadata details for an origin's shared storage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-SharedStorageMetadata
type SharedStorageMetadata struct {
	CreationTime    *cdp.TimeSinceEpoch `json:"creationTime"`
	Length          int64               `json:"length"`
	RemainingBudget float64             `json:"remainingBudget"`
}

// SharedStorageReportingMetadata pair of reporting metadata details for a
// candidate URL for selectURL().
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-SharedStorageReportingMetadata
type SharedStorageReportingMetadata struct {
	EventType    string `json:"eventType"`
	ReportingURL string `json:"reportingUrl"`
}

// SharedStorageURLWithMetadata bundles a candidate URL with its reporting
// metadata.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-SharedStorageUrlWithMetadata
type SharedStorageURLWithMetadata struct {
	URL               string                            `json:"url"`               // Spec of candidate URL.
	ReportingMetadata []*SharedStorageReportingMetadata `json:"reportingMetadata"` // Any associated reporting metadata.
}

// SharedStorageAccessParams bundles the parameters for shared storage access
// events whose presence/absence can vary according to SharedStorageAccessType.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-SharedStorageAccessParams
type SharedStorageAccessParams struct {
	ScriptSourceURL  string                          `json:"scriptSourceUrl,omitempty"`  // Spec of the module script URL. Present only for SharedStorageAccessType.documentAddModule.
	OperationName    string                          `json:"operationName,omitempty"`    // Name of the registered operation to be run. Present only for SharedStorageAccessType.documentRun and SharedStorageAccessType.documentSelectURL.
	SerializedData   string                          `json:"serializedData,omitempty"`   // The operation's serialized data in bytes (converted to a string). Present only for SharedStorageAccessType.documentRun and SharedStorageAccessType.documentSelectURL.
	UrlsWithMetadata []*SharedStorageURLWithMetadata `json:"urlsWithMetadata,omitempty"` // Array of candidate URLs' specs, along with any associated metadata. Present only for SharedStorageAccessType.documentSelectURL.
	Key              string                          `json:"key,omitempty"`              // Key for a specific entry in an origin's shared storage. Present only for SharedStorageAccessType.documentSet, SharedStorageAccessType.documentAppend, SharedStorageAccessType.documentDelete, SharedStorageAccessType.workletSet, SharedStorageAccessType.workletAppend, SharedStorageAccessType.workletDelete, and SharedStorageAccessType.workletGet.
	Value            string                          `json:"value,omitempty"`            // Value for a specific entry in an origin's shared storage. Present only for SharedStorageAccessType.documentSet, SharedStorageAccessType.documentAppend, SharedStorageAccessType.workletSet, and SharedStorageAccessType.workletAppend.
	IgnoreIfPresent  bool                            `json:"ignoreIfPresent,omitempty"`  // Whether or not to set an entry for a key if that key is already present. Present only for SharedStorageAccessType.documentSet and SharedStorageAccessType.workletSet.
}

// BucketsDurability [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-StorageBucketsDurability
type BucketsDurability string

// String returns the BucketsDurability as string value.
func (t BucketsDurability) String() string {
	return string(t)
}

// BucketsDurability values.
const (
	BucketsDurabilityRelaxed BucketsDurability = "relaxed"
	BucketsDurabilityStrict  BucketsDurability = "strict"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t BucketsDurability) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t BucketsDurability) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *BucketsDurability) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch BucketsDurability(v) {
	case BucketsDurabilityRelaxed:
		*t = BucketsDurabilityRelaxed
	case BucketsDurabilityStrict:
		*t = BucketsDurabilityStrict

	default:
		in.AddError(fmt.Errorf("unknown BucketsDurability value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *BucketsDurability) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Bucket [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-StorageBucket
type Bucket struct {
	StorageKey SerializedStorageKey `json:"storageKey"`
	Name       string               `json:"name,omitempty"` // If not specified, it is the default bucket of the storageKey.
}

// BucketInfo [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-StorageBucketInfo
type BucketInfo struct {
	Bucket     *Bucket             `json:"bucket"`
	ID         string              `json:"id"`
	Expiration *cdp.TimeSinceEpoch `json:"expiration"`
	Quota      float64             `json:"quota"` // Storage quota (bytes).
	Persistent bool                `json:"persistent"`
	Durability BucketsDurability   `json:"durability"`
}

// AttributionReportingSourceType [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-AttributionReportingSourceType
type AttributionReportingSourceType string

// String returns the AttributionReportingSourceType as string value.
func (t AttributionReportingSourceType) String() string {
	return string(t)
}

// AttributionReportingSourceType values.
const (
	AttributionReportingSourceTypeNavigation AttributionReportingSourceType = "navigation"
	AttributionReportingSourceTypeEvent      AttributionReportingSourceType = "event"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AttributionReportingSourceType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AttributionReportingSourceType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AttributionReportingSourceType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch AttributionReportingSourceType(v) {
	case AttributionReportingSourceTypeNavigation:
		*t = AttributionReportingSourceTypeNavigation
	case AttributionReportingSourceTypeEvent:
		*t = AttributionReportingSourceTypeEvent

	default:
		in.AddError(fmt.Errorf("unknown AttributionReportingSourceType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AttributionReportingSourceType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// UnsignedInt64asBase10 [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-UnsignedInt64AsBase10
type UnsignedInt64asBase10 string

// String returns the UnsignedInt64asBase10 as string value.
func (t UnsignedInt64asBase10) String() string {
	return string(t)
}

// UnsignedInt128asBase16 [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-UnsignedInt128AsBase16
type UnsignedInt128asBase16 string

// String returns the UnsignedInt128asBase16 as string value.
func (t UnsignedInt128asBase16) String() string {
	return string(t)
}

// SignedInt64asBase10 [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-SignedInt64AsBase10
type SignedInt64asBase10 string

// String returns the SignedInt64asBase10 as string value.
func (t SignedInt64asBase10) String() string {
	return string(t)
}

// AttributionReportingFilterDataEntry [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-AttributionReportingFilterDataEntry
type AttributionReportingFilterDataEntry struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// AttributionReportingAggregationKeysEntry [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-AttributionReportingAggregationKeysEntry
type AttributionReportingAggregationKeysEntry struct {
	Key   string                 `json:"key"`
	Value UnsignedInt128asBase16 `json:"value"`
}

// AttributionReportingEventReportWindows [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-AttributionReportingEventReportWindows
type AttributionReportingEventReportWindows struct {
	Start int64   `json:"start"` // duration in seconds
	Ends  []int64 `json:"ends"`  // duration in seconds
}

// AttributionReportingSourceRegistration [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-AttributionReportingSourceRegistration
type AttributionReportingSourceRegistration struct {
	Time                     *cdp.TimeSinceEpoch                         `json:"time"`
	Expiry                   int64                                       `json:"expiry"` // duration in seconds
	EventReportWindows       *AttributionReportingEventReportWindows     `json:"eventReportWindows"`
	AggregatableReportWindow int64                                       `json:"aggregatableReportWindow"` // duration in seconds
	Type                     AttributionReportingSourceType              `json:"type"`
	SourceOrigin             string                                      `json:"sourceOrigin"`
	ReportingOrigin          string                                      `json:"reportingOrigin"`
	DestinationSites         []string                                    `json:"destinationSites"`
	EventID                  UnsignedInt64asBase10                       `json:"eventId"`
	Priority                 SignedInt64asBase10                         `json:"priority"`
	FilterData               []*AttributionReportingFilterDataEntry      `json:"filterData"`
	AggregationKeys          []*AttributionReportingAggregationKeysEntry `json:"aggregationKeys"`
	DebugKey                 UnsignedInt64asBase10                       `json:"debugKey,omitempty"`
}

// AttributionReportingSourceRegistrationResult [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-AttributionReportingSourceRegistrationResult
type AttributionReportingSourceRegistrationResult string

// String returns the AttributionReportingSourceRegistrationResult as string value.
func (t AttributionReportingSourceRegistrationResult) String() string {
	return string(t)
}

// AttributionReportingSourceRegistrationResult values.
const (
	AttributionReportingSourceRegistrationResultSuccess                               AttributionReportingSourceRegistrationResult = "success"
	AttributionReportingSourceRegistrationResultInternalError                         AttributionReportingSourceRegistrationResult = "internalError"
	AttributionReportingSourceRegistrationResultInsufficientSourceCapacity            AttributionReportingSourceRegistrationResult = "insufficientSourceCapacity"
	AttributionReportingSourceRegistrationResultInsufficientUniqueDestinationCapacity AttributionReportingSourceRegistrationResult = "insufficientUniqueDestinationCapacity"
	AttributionReportingSourceRegistrationResultExcessiveReportingOrigins             AttributionReportingSourceRegistrationResult = "excessiveReportingOrigins"
	AttributionReportingSourceRegistrationResultProhibitedByBrowserPolicy             AttributionReportingSourceRegistrationResult = "prohibitedByBrowserPolicy"
	AttributionReportingSourceRegistrationResultSuccessNoised                         AttributionReportingSourceRegistrationResult = "successNoised"
	AttributionReportingSourceRegistrationResultDestinationReportingLimitReached      AttributionReportingSourceRegistrationResult = "destinationReportingLimitReached"
	AttributionReportingSourceRegistrationResultDestinationGlobalLimitReached         AttributionReportingSourceRegistrationResult = "destinationGlobalLimitReached"
	AttributionReportingSourceRegistrationResultDestinationBothLimitsReached          AttributionReportingSourceRegistrationResult = "destinationBothLimitsReached"
	AttributionReportingSourceRegistrationResultReportingOriginsPerSiteLimitReached   AttributionReportingSourceRegistrationResult = "reportingOriginsPerSiteLimitReached"
	AttributionReportingSourceRegistrationResultExceedsMaxChannelCapacity             AttributionReportingSourceRegistrationResult = "exceedsMaxChannelCapacity"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AttributionReportingSourceRegistrationResult) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AttributionReportingSourceRegistrationResult) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AttributionReportingSourceRegistrationResult) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch AttributionReportingSourceRegistrationResult(v) {
	case AttributionReportingSourceRegistrationResultSuccess:
		*t = AttributionReportingSourceRegistrationResultSuccess
	case AttributionReportingSourceRegistrationResultInternalError:
		*t = AttributionReportingSourceRegistrationResultInternalError
	case AttributionReportingSourceRegistrationResultInsufficientSourceCapacity:
		*t = AttributionReportingSourceRegistrationResultInsufficientSourceCapacity
	case AttributionReportingSourceRegistrationResultInsufficientUniqueDestinationCapacity:
		*t = AttributionReportingSourceRegistrationResultInsufficientUniqueDestinationCapacity
	case AttributionReportingSourceRegistrationResultExcessiveReportingOrigins:
		*t = AttributionReportingSourceRegistrationResultExcessiveReportingOrigins
	case AttributionReportingSourceRegistrationResultProhibitedByBrowserPolicy:
		*t = AttributionReportingSourceRegistrationResultProhibitedByBrowserPolicy
	case AttributionReportingSourceRegistrationResultSuccessNoised:
		*t = AttributionReportingSourceRegistrationResultSuccessNoised
	case AttributionReportingSourceRegistrationResultDestinationReportingLimitReached:
		*t = AttributionReportingSourceRegistrationResultDestinationReportingLimitReached
	case AttributionReportingSourceRegistrationResultDestinationGlobalLimitReached:
		*t = AttributionReportingSourceRegistrationResultDestinationGlobalLimitReached
	case AttributionReportingSourceRegistrationResultDestinationBothLimitsReached:
		*t = AttributionReportingSourceRegistrationResultDestinationBothLimitsReached
	case AttributionReportingSourceRegistrationResultReportingOriginsPerSiteLimitReached:
		*t = AttributionReportingSourceRegistrationResultReportingOriginsPerSiteLimitReached
	case AttributionReportingSourceRegistrationResultExceedsMaxChannelCapacity:
		*t = AttributionReportingSourceRegistrationResultExceedsMaxChannelCapacity

	default:
		in.AddError(fmt.Errorf("unknown AttributionReportingSourceRegistrationResult value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AttributionReportingSourceRegistrationResult) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}