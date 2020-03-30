package models

type HttpError struct {
	ErrorMessage    string   `json:"errorMessage"`
	FormErrorFields []string `json:"formErrorFields"`
}
