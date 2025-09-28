package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type XMLDataProcessor struct {
	name string
}

func NewXMLDataProcessor() *XMLDataProcessor {
	return &XMLDataProcessor{
		name: "XML Data Processor",
	}
}

func (x *XMLDataProcessor) ProcessXML(xmlData string) (string, error) {
	if err := x.ValidateXML(xmlData); err != nil {
		return "", err
	}

	var user User
	if err := xml.Unmarshal([]byte(xmlData), &user); err != nil {
		return "", fmt.Errorf("%w: %v", ErrProcessingFailed, err)
	}

	result := fmt.Sprintf("XML Processed user: ID=%d, Name=%s, Email=%s", user.ID, user.Name, user.Email)
	return result, nil
}

func (x *XMLDataProcessor) ValidateXML(xmlData string) error {
	xmlData = strings.TrimSpace(xmlData)
	if !strings.HasPrefix(xmlData, "<") || !strings.HasSuffix(xmlData, ">") {
		return ErrInvalidData
	}
	return nil
}

func (x *XMLDataProcessor) GetName() string {
	return x.name
}
