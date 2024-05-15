package model

// TODO: need to fix fiels of json
type SubjectResponse struct {
	Name        string `json:"Name"`
	Cabinet     string `json:"Cabinet"`
	Teacher     string `json:"Teacher"`
	Description string `json:"Description"`
	StartTime   string `json:"StartTime"`
	EndTime     string `json:"EndTime"`
}
