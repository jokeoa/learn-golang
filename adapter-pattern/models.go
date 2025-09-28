package main

type User struct {
	ID    int    `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

type ProcessingResult struct {
	OriginalFormat string `json:"original_format"`
	ProcessedData  string `json:"processed_data"`
	Success        bool   `json:"success"`
	ErrorMessage   string `json:"error_message,omitempty"`
}
