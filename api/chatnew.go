package handler

import (
	"encoding/json"
	"fmt"
	"golangnext/goservice"
	"net/http"
)

func ChatNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		aiService := goservice.GetAIService()

		newChatResponse, err := aiService.CreateNewChat()
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonString, err := json.Marshal(newChatResponse)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		goservice.ResponseJson(w, jsonString, http.StatusOK)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
