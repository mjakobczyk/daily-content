package env

// APIError informs about possible failures in requests.
type APIError string

const (
	// RequestFailedError when service call fails.
	RequestFailedError APIError = "HTTP_REQUEST_FAILED"

	// ReadDataFailedError when response contains errors.
	ReadDataFailedError APIError = "READ_DATA_FAILED"

	// UnmarshalDataFailedError when unmarshal data from response fails.
	UnmarshalDataFailedError APIError = "UNMARSHAL_DATA_FAILED"
)

// Error returns error value as string.
func (e APIError) Error() string {
	return string(e)
}

// String returns string value.
func (e APIError) String() string {
	return string(e)
}
