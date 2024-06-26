package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/ppfeister/gostuff/messages"
)

type Targets struct {
	Targets []Target `json:"standard"`
}

type Target struct {
	Name               string `json:"name"`
	Url                string `json:"request_url"`
	Method             string `json:"request_method"`
	Data               string `json:"request_data"`
	Validation_Method  string `json:"validation_method"`
	Validation_Type    string `json:"validation_type"`
	Validation_Message string `json:"message"`
	Response_Data      response_data
}

type response_data struct {
	Status   uint8
	Headers  string
	Data     string
	Validity bool
}

var manifest_data Targets

func FetchManifest() Targets {
	file, err := os.Open("resources/targets.json")
	if err != nil {
		log.Fatal(messages.Fetch(i18n.LocalizeConfig{
			MessageID: "Err_UnableToLoadManifest",
			TemplateData: map[string]interface{}{
				"Error": err.Error(),
			},
		}))
	}
	byteval, _ := io.ReadAll(file)
	file.Close()
	json.Unmarshal(byteval, &manifest_data)
	return manifest_data
}
