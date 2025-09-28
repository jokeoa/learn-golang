package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type XMLToJSONAdapter struct {
	xmlProcessor XMLProcessor
}

func NewXMLToJSONAdapter(xmlProcessor XMLProcessor) *XMLToJSONAdapter {
	return &XMLToJSONAdapter{
		xmlProcessor: xmlProcessor,
	}
}

func (adapter *XMLToJSONAdapter) ProcessData(data string) (string, error) {
	if err := adapter.xmlProcessor.ValidateXML(data); err != nil {
		return "", err
	}

	var user User
	if err := xml.Unmarshal([]byte(data), &user); err != nil {
		return "", fmt.Errorf("%w: failed to parse XML: %v", ErrProcessingFailed, err)
	}

	result := ProcessingResult{
		OriginalFormat: "XML",
		ProcessedData:  fmt.Sprintf("XML adapted to JSON - User: ID=%d, Name=%s, Email=%s", user.ID, user.Name, user.Email),
		Success:        true,
	}

	processedJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("%w: failed to marshal JSON: %v", ErrProcessingFailed, err)
	}

	return string(processedJSON), nil
}

func (adapter *XMLToJSONAdapter) GetSupportedFormat() string {
	return "XML-to-JSON"
}
