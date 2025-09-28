package main

import "errors"

var (
	ErrInvalidData      = errors.New("invalid data format")
	ErrProcessingFailed = errors.New("data processing failed")
)

type DataProcessor interface {
	ProcessData(data string) (string, error)
	GetSupportedFormat() string
}

type XMLProcessor interface {
	ProcessXML(xmlData string) (string, error)
	ValidateXML(xmlData string) error
}
