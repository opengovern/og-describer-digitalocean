//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

package provider

type Metadata struct{}

type FileResponse struct {
	Data    []FileDescription `json:"data"`
	LastID  string            `json:"last_id"`
	HasMore bool              `json:"has_more"`
}

type FileDescription struct {
	ID        string `json:"id"`
	FileName  string `json:"file_name"`
	CreatedAt int32  `json:"created_at"`
	Bytes     int32  `json:"bytes"`
	Object    string `json:"object"`
	Purpose   string `json:"purpose"`
}

type ModelsResponse struct {
	Data []ModelsDescription `json:"data"`
}

type Permission struct {
	CreatedAt          int64       `json:"created"`
	ID                 string      `json:"id"`
	Object             string      `json:"object"`
	AllowCreateEngine  bool        `json:"allow_create_engine"`
	AllowSampling      bool        `json:"allow_sampling"`
	AllowLogprobs      bool        `json:"allow_logprobs"`
	AllowSearchIndices bool        `json:"allow_search_indices"`
	AllowView          bool        `json:"allow_view"`
	AllowFineTuning    bool        `json:"allow_fine_tuning"`
	Organization       string      `json:"organization"`
	Group              interface{} `json:"group"`
	IsBlocking         bool        `json:"is_blocking"`
}

type ModelsDescription struct {
	ID         string       `json:"id"`
	CreatedAt  int32        `json:"created_at"`
	Object     string       `json:"object"`
	OwnedBy    string       `json:"owned_by"`
	Permission []Permission `json:"permission"`
	Root       string       `json:"root"`
	//Parent     string       `json:"parent"`
}

type AssistantObjectToolsInner struct {
	AssistantToolsCode       *AssistantToolsCode
	AssistantToolsFileSearch *AssistantToolsFileSearch
	AssistantToolsFunction   *AssistantToolsFunction
}

type AssistantToolsCode struct {
	Type string `json:"type"`
}

type AssistantToolsFileSearch struct {
	Type       string                              `json:"type"`
	FileSearch *AssistantToolsFileSearchFileSearch `json:"file_search,omitempty"`
}

type AssistantToolsFileSearchFileSearch struct {
	MaxNumResults  *int32                    `json:"max_num_results,omitempty"`
	RankingOptions *FileSearchRankingOptions `json:"ranking_options,omitempty"`
}

type FileSearchRankingOptions struct {
	Ranker         *string `json:"ranker,omitempty"`
	ScoreThreshold float32 `json:"score_threshold"`
}

type AssistantToolsFunction struct {
	Type     string         `json:"type"`
	Function FunctionObject `json:"function"`
}

type FunctionObject struct {
	Description *string                `json:"description,omitempty"`
	Name        string                 `json:"name"`
	Parameters  map[string]interface{} `json:"parameters,omitempty"`
	Strict      NullableBool           `json:"strict,omitempty"`
}

type NullableBool struct {
	value *bool
	isSet bool
}

type AssistantObjectToolResources struct {
	CodeInterpreter *AssistantObjectToolResourcesCodeInterpreter `json:"code_interpreter,omitempty"`
	FileSearch      *AssistantObjectToolResourcesFileSearch      `json:"file_search,omitempty"`
}

type AssistantObjectToolResourcesFileSearch struct {
	VectorStoreIds []string `json:"vector_store_ids,omitempty"`
}

type AssistantObjectToolResourcesCodeInterpreter struct {
	FileIds []string `json:"file_ids,omitempty"`
}

type ResponseFormatJsonObject struct {
	Type string `json:"type"`
}

type ResponseFormatJsonSchema struct {
	Type       string                             `json:"type"`
	JsonSchema ResponseFormatJsonSchemaJsonSchema `json:"json_schema"`
}

type ResponseFormatJsonSchemaJsonSchema struct {
	Description *string                `json:"description,omitempty"`
	Name        string                 `json:"name"`
	Schema      map[string]interface{} `json:"schema,omitempty"`
	Strict      NullableBool           `json:"strict,omitempty"`
}

type ResponseFormatText struct {
	Type string `json:"type"`
}

type AssistantResponse struct {
	Data    []AssistantDescription `json:"data"`
	LastID  string                 `json:"last_id"`
	HasMore bool                   `json:"has_more"`
}

type AssistantDescription struct {
	//Object         string
	//CreatedAt      time.Time
	ID            string                 `json:"id"`
	Name          *string                `json:"name"`
	Description   *string                `json:"description"`
	Model         string                 `json:"model"`
	Instructions  *string                `json:"instructions"`
	Tools         []interface{}          `json:"tools"`
	ToolResources *interface{}           `json:"tool_resources"`
	Metadata      map[string]interface{} `json:"metadata"`
	Temperature   *float32               `json:"temperature"`
	TopP          *float32               `json:"top_p"`
	//ResponseFormat *interface{}           `json:"response_format"`
}

type VectorStoreObjectFileCounts struct {
	InProgress int32 `json:"in_progress"`
	Completed  int32 `json:"completed"`
	Failed     int32 `json:"failed"`
	Cancelled  int32 `json:"cancelled"`
	Total      int32 `json:"total"`
}

type VectorStoreExpirationAfter struct {
	Anchor string `json:"anchor"`
	Days   int32  `json:"days"`
}

type VectorStoreResponse struct {
	Data    []VectorStoreDescription `json:"data"`
	LastID  string                   `json:"last_id"`
	HasMore bool                     `json:"has_more"`
}

type VectorStoreDescription struct {
	ID           string                      `json:"id"`
	Object       string                      `json:"object"`
	CreatedAt    int32                       `json:"created_at"`
	Name         string                      `json:"name"`
	UsageBytes   int32                       `json:"usage_bytes"`
	FileCounts   VectorStoreObjectFileCounts `json:"file_counts"`
	Status       string                      `json:"status"`
	ExpiresAfter *VectorStoreExpirationAfter `json:"expires_after"`
	ExpiresAt    *int32                      `json:"expires_at"`
	LastActiveAt *int32                      `json:"last_active_at"`
	Metadata     map[string]interface{}      `json:"metadata"`
}

type NullableInt32 struct {
	value *int32
	isSet bool
}

type ProjectResponse struct {
	Data    []ProjectDescription `json:"data"`
	LastID  string               `json:"last_id"`
	HasMore bool                 `json:"has_more"`
}

type ProjectDescription struct {
	ID         string        `json:"id"`
	Object     string        `json:"object"`
	Name       string        `json:"name"`
	CreatedAt  int32         `json:"created_at"`
	ArchivedAt NullableInt32 `json:"archived_at"`
	Status     string        `json:"status"`
}

type ProjectUserResponse struct {
	Data    []ProjectUser `json:"data"`
	LastID  string        `json:"last_id"`
	HasMore bool          `json:"has_more"`
}

type ProjectUser struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Object  string `json:"object"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	AddedAt int32  `json:"added_at"`
}

type ProjectUserDescription struct {
	UserID    string `json:"user_id"`
	ProjectID string `json:"project_id"`
	Object    string
	Name      string
	Email     string
	Role      string
}

type ProjectServiceAccountResponse struct {
	Data    []ProjectServiceAccount `json:"data"`
	LastID  string                  `json:"last_id"`
	HasMore bool                    `json:"has_more"`
}

type ProjectServiceAccount struct {
	Object    string `json:"object"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	CreatedAt int32  `json:"created_at"`
}

type ProjectServiceAccountDescription struct {
	ProjectServiceAccount
	ProjectID string `json:"project_id"`
}

type ProjectApiKeyOwner struct {
	Type           *string                `json:"type,omitempty"`
	User           *ProjectUser           `json:"user,omitempty"`
	ServiceAccount *ProjectServiceAccount `json:"service_account,omitempty"`
}

type ProjectApiKeyResponse struct {
	Data    []ProjectApiKey `json:"data"`
	LastID  string          `json:"last_id"`
	HasMore bool            `json:"has_more"`
}

type ProjectApiKey struct {
	Object        string             `json:"object"`
	ID            string             `json:"id"`
	Name          string             `json:"name"`
	RedactedValue string             `json:"redacted_value"`
	CreatedAt     int32              `json:"created_at"`
	Owner         ProjectApiKeyOwner `json:"owner"`
}

type ProjectApiKeyDescription struct {
	ProjectApiKey
	ProjectID string `json:"project_id"`
}

type ProjectRateLimitResponse struct {
	Data    []ProjectRateLimit `json:"data"`
	LastID  string             `json:"last_id"`
	HasMore bool               `json:"has_more"`
}

type ProjectRateLimit struct {
	Object                      string `json:"object"`
	ID                          string `json:"id"`
	Model                       string `json:"model"`
	MaxRequestsPer1Minute       int32  `json:"max_requests_per_1_minute"`
	MaxTokensPer1Minute         int32  `json:"max_tokens_per_1_minute"`
	MaxImagesPer1Minute         *int32 `json:"max_images_per_1_minute"`
	MaxAudioMegabytesPer1Minute *int32 `json:"max_audio_megabytes_per_1_minute"`
	MaxRequestsPer1Day          *int32 `json:"max_requests_per_1_day"`
	Batch1DayMaxInputTokens     *int32 `json:"batch_1_day_max_input_tokens"`
}

type ProjectRateLimitDescription struct {
	ProjectRateLimit
	ProjectID string `json:"project_id"`
}
