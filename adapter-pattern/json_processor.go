package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type JSONDataProcessor struct {
	name string
}

func NewJSONDataProcessor() *JSONDataProcessor {
	return &JSONDataProcessor{
		name: "JSON Data Processor",
	}
}

func (j *JSONDataProcessor) ProcessData(data string) (string, error) {
	if !j.isValidJSON(data) {
		return "", ErrInvalidData
	}

	var user User
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return "", fmt.Errorf("%w: %v", ErrProcessingFailed, err)
	}

	result := ProcessingResult{
		OriginalFormat: j.GetSupportedFormat(),
		ProcessedData:  fmt.Sprintf("Processed user: ID=%d, Name=%s, Email=%s", user.ID, user.Name, user.Email),
		Success:        true,
	}

	processedJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrProcessingFailed, err)
	}

	return string(processedJSON), nil
}

func (j *JSONDataProcessor) GetSupportedFormat() string {
	return "JSON"
}

func (j *JSONDataProcessor) isValidJSON(data string) bool {
	data = strings.TrimSpace(data)
	return (strings.HasPrefix(data, "{") && strings.HasSuffix(data, "}")) ||
		(strings.HasPrefix(data, "[") && strings.HasSuffix(data, "]"))
}
