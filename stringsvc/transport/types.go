package transport

type UppercaseRequest struct {
	S string `json:"s,omitempty"`
}

type UppercaseResponse struct {
	V string `json:"v,omitempty"`
}

type CountRequest struct {
	S string `json:"s,omitempty"`
}

type CountResponse struct {
	N int `json:"v,omitempty"`
}
