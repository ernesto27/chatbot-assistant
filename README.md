
# Chatbot assistant

## Setup - requiremente

You need a apikey from openAI 

https://platform.openai.com/api-keys

Create assistant,  associate files, documents via CLI

On terminal run
```sh
export OPENAI_API_KEY=[yourapikey]
```

Create assistant 
```sh
go run goservice/cmd/cmd.go --cmd create-assistant --name myassistant --instructions "your system instruction"
```

Upload file 
```sh
go run goservice/cmd/cmd.go --cmd create-file --file /yourpathfile
```

Associate file to assistant
Upload file 
```sh
go run goservice/cmd/cmd.go --cmd create-assistant-file --assistant-id yourAssistantID  --file-id fileID
```

You can also create the assistan on the openAI dashboard 

https://platform.openai.com/assistants/


## Run API service 

On terminal run
```sh
export OPENAI_API_KEY=[yourapikey]
export ASSISTANT_ID=[assistantID]
```
Run go api service

```sh
go run .
```

## Run front

```sh
npm i 
npm run dev
```


## Vercel configuration 

On your vercel project setup, go to settings => Environment variables and add this values.

ASSISTANT_ID => yourassistanid

OPENAI_API_KEY => yourapikey

NEXT_PUBLIC_ENV => production









