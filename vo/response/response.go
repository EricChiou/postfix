package response

// Response vo
type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Trace  string      `json:"trace,omitempty"`
}
