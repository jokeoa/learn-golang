package main

import (
	"fmt"
)

type DataProcessingService struct {
	processor DataProcessor
}

func NewDataProcessingService(processor DataProcessor) *DataProcessingService {
	return &DataProcessingService{
		processor: processor,
	}
}

func (service *DataProcessingService) ProcessUserData(data string) error {
	fmt.Printf("Processing data using %s processor...\n", service.processor.GetSupportedFormat())

	result, err := service.processor.ProcessData(data)
	if err != nil {
		fmt.Printf("Processing failed: %v\n", err)
		return err
	}

	fmt.Printf("Processing successful!\n")
	fmt.Printf("Result:\n%s\n", result)
	return nil
}

func (service *DataProcessingService) SetProcessor(processor DataProcessor) {
	service.processor = processor
}
