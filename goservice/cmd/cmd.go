package main

import (
	"flag"
	"fmt"
	"golangnext/goservice"
	"os"
)

func main() {
	cmd := flag.String("cmd", "", "service to run")
	name := flag.String("name", "", "name assistant")
	pathFile := flag.String("file", "", "path file")
	fileID := flag.String("file-id", "", "file id")
	assistantID := flag.String("assistant-id", "", "assistatn id")
	instructions := flag.String("instructions", "", "instructions assistant")
	flag.Parse()

	aiService := goservice.GetAIService()

	switch *cmd {
	case "create-assistant":
		assistant, err := aiService.CreateAssistant(*name, *instructions)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Assistant ID created: " + assistant)

	case "create-file":
		fileUpload, err := aiService.CreateFile(*pathFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("File uploaded ID: " + fileUpload)

	case "create-assistant-file":
		assistantFile, err := aiService.CreateAssistantFile(*assistantID, *fileID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Assistant file created: " + assistantFile)
	default:
		fmt.Println("Command not found")

	}
}
