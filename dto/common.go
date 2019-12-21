package dto

type ListMetadata struct {
	Count  int64 `json:"count"`
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
	Total  int64 `json:"total"`
}

type ListResp struct {
	Metadata ListMetadata  `json:"metadata"`
	Results  []interface{} `json:"results"`
}

type ListRequest struct {
	Limit  int64
	Offset int64
}
