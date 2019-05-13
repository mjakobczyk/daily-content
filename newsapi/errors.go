package newsapi

// ProductError informs about possible failures in requests.
type NewsAPIError string

const (
	// RequestFailedError when service call fails.
	RequestFailedError NewsAPIError = "HTTP_REQUEST_FAILED"

	// ReadDataFailedError when response contains errors.
	ReadDataFailedError NewsAPIError = "READ_DATA_FAILED"

	// UnmarshalDataFailedError when unmarshal data from response fails.
	UnmarshalDataFailedError NewsAPIError = "UNMARSHAL_DATA_FAILED"
)

// Error returns error as a string.
func (n NewsAPIError) Error() string {
	return string(n)
}

// String returns string value.
func (n NewsAPIError) String() string {
	return string(n)
}
