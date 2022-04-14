package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type JsonSructure struct {
	ConferenceName     string `json:"conferenceName"`
	ConferenceLocation string `json:"conferenceLocation"`
	AvailableTickets   uint   `json:"availableTickets"`
}

func WriteToFile(data *JsonSructure) error {
    jsonFilePath, _ := filepath.Abs("configs/conferences.json")

    // first read json file
    jsonFile := *ReadConfigFile()

    // remove the recod being updated
    for index, obj := range jsonFile {
        if data.ConferenceName == obj.ConferenceName {
            jsonFile = append(jsonFile[:index], jsonFile[index+1:]...)
        }
    }

    // create a new record
    newRecod := JsonSructure{
        ConferenceName: data.ConferenceName,
        ConferenceLocation: data.ConferenceLocation,
        AvailableTickets: data.AvailableTickets,
    }

    jsonFile = append(jsonFile, newRecod)

    file, _ := json.MarshalIndent(&jsonFile, "", " ")

    err := ioutil.WriteFile(jsonFilePath, file, 0644)

    if err != nil {
        return err
    }

    return nil
}

func ReadConfigFile() *[]JsonSructure {
    jsonFilePath, _ := filepath.Abs("configs/conferences.json")
	file, err := os.ReadFile(jsonFilePath)
	jsonFile := []JsonSructure{}

	if err != nil {
		fmt.Println(err)
	}

	_ = json.Unmarshal([]byte(file), &jsonFile)

	return &jsonFile
}
