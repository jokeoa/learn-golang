package main

import (
	"fmt"
)

func main() {

	jsonData := `{
		"id": 1,
		"name": "John Doe",
		"email": "john.doe@example.com"
	}`

	xmlData := `<User>
		<id>2</id>
		<name>Jane Smith</name>
		<email>jane.smith@example.com</email>
	</User>`

	jsonProcessor := NewJSONDataProcessor()
	service := NewDataProcessingService(jsonProcessor)

	fmt.Println(" Processing JSON data directly:")
	service.ProcessUserData(jsonData)

	xmlProcessor := NewXMLDataProcessor()
	xmlAdapter := NewXMLToJSONAdapter(xmlProcessor)
	service.SetProcessor(xmlAdapter)

	fmt.Println(" Processing XML data via adapter:")
	service.ProcessUserData(xmlData)
}
