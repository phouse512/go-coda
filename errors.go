package coda

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	Message       string `json:"message"`
}

func buildError(resp *http.Response) error {
	var errResp ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		errResp = ErrorResponse{
			StatusCode:    resp.StatusCode,
			StatusMessage: "Unknown Error",
		}
	}

	switch errResp.StatusCode {
	case 400:
		return InvalidRequestError{err: errResp.StatusMessage}
	case 401:
		return InvalidTokenError{err: errResp.StatusMessage}
	case 403:
		return UnauthorizedError{err: errResp.StatusMessage}
	case 404:
		return ResourceNotFoundError{err: errResp.StatusMessage}
	case 410:
		return ResourceDeletedError{err: errResp.StatusMessage}
	case 429:
		return RateLimitError{err: errResp.StatusMessage}
	case 502:
		return InternalError{err: errResp.StatusMessage}
	default:
		return ApiError{err: errResp.StatusMessage}
	}
}

// Standard Bad Request 400 Error
type InvalidRequestError struct {
	err string
}

func (e InvalidRequestError) Error() string {
	return fmt.Sprintf("Invalid request: %s", e.err)
}

// Invalid API Token
type InvalidTokenError struct {
	err string
}

func (e InvalidTokenError) Error() string {
	return fmt.Sprintf("Invalid token: %s", e.err)
}

// Rate Limit Error
type RateLimitError struct {
	err string
}

func (e RateLimitError) Error() string {
	return fmt.Sprintf("Rate limited: %s", e.err)
}

type ApiError struct {
	err string
}

func (e ApiError) Error() string {
	return fmt.Sprintf("Coda API Error: %s", e.err)
}

type ResourceDeletedError struct {
	err string
}

func (e ResourceDeletedError) Error() string {
	return fmt.Sprintf("Resource deleted: %s", e.err)
}

type ResourceNotFoundError struct {
	err string
}

func (e ResourceNotFoundError) Error() string {
	return fmt.Sprintf("Resource not found: %s", e.err)
}

type UnauthorizedError struct {
	err string
}

func (e UnauthorizedError) Error() string {
	return fmt.Sprintf("Token does not have access: %s", e.err)
}

type InternalError struct {
	err string
}

func (e InternalError) Error() string {
	return fmt.Sprintf("Internal api error: %s", e.err)
}
