package service_code

type ServiceCodeResponse struct {
	ID                    string `json:"id"`
	Code                  string `json:"code"`
	ServiceTimeRangeStart int    `json:"service_time_range_start"`
	ServiceTimeRangeEnd   int    `json:"service_time_range_end"`
}

type ServiceCodeListResponse []ServiceCodeResponse
