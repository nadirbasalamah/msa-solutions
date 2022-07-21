package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const BASE_URL = "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&limit=5"

type GraduationData struct {
	Help    string `json:"help"`
	Success bool   `json:"success"`
	Result  struct {
		ResourceID string `json:"resource_id"`
		Fields     []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"fields"`
		Records []struct {
			ID            int    `json:"_id"`
			Sex           string `json:"sex"`
			NoOfGraduates string `json:"no_of_graduates"`
			TypeOfCourse  string `json:"type_of_course"`
			Year          string `json:"year"`
		} `json:"records"`
		Links struct {
			Start string `json:"start"`
			Next  string `json:"next"`
		} `json:"_links"`
		Limit int `json:"limit"`
		Total int `json:"total"`
	} `json:"result"`
}

func fetchGraduationData() (GraduationData, error) {
	var err error
	var client = &http.Client{}
	var data GraduationData

	request, err := http.NewRequest(http.MethodGet, BASE_URL, nil)
	if err != nil {
		return GraduationData{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return GraduationData{}, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return GraduationData{}, err
	}

	return data, nil
}

func main() {
	var graduationData, err = fetchGraduationData()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	graduations := graduationData.Result.Records

	csvFile, err := os.Create("graduation.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, graduation := range graduations {
		var data []string = []string{
			strconv.Itoa(graduation.ID),
			graduation.Sex,
			graduation.TypeOfCourse,
			graduation.NoOfGraduates,
			graduation.Year,
		}

		_ = csvwriter.Write(data)
	}
	csvwriter.Flush()
	csvFile.Close()
}
