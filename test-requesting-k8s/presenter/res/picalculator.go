package res

type CalculatePi struct {
	PiValue  float64 `json:"piValue"`
	HostName string  `json:"hostName"`
}

type TestAPIResponse struct {
	Time  int64 `json:"timeInMilliseconds"`
	Hits  int   `json:"hits"`
	Fails int   `json:"fails"`
}
