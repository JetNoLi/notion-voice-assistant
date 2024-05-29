package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadEnv() map[string]string {
	rawEnvFile, err := os.ReadFile(".env")

	if err != nil {
		log.Fatal(err.Error())
		panic(0)
	}

	fmt.Println("Env File Read Successfully")

	envFileLines := strings.Split(string(rawEnvFile), "\n")
	envVars := make(map[string]string)

	for _, line := range envFileLines {
		// Ignore Comments
		if line[0] == '#' {
			continue
		}
		lineValues := strings.Split(line, "=")
		key, value := lineValues[0], lineValues[1]
		envVars[key] = value
	}

	return envVars
}

var config map[string]string = ReadEnv()

var NotionApiKey string = config["NOTION_API_KEY"]
var NotionMainDbId string = config["NOTION_MAIN_DB_ID"]
var WhisperApiUrl string = config["WHISPER_API_URL"]
var OpenAiApiUrl string = config["OPEN_AI_API_URL"]
var OpenAiApiKey string = config["OPEN_AI_API_KEY"]
var Port string = config["PORT"]
