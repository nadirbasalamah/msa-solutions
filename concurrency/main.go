package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// define the BASE_URL of the API
const BASE_URL = "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&limit=100"

// GraduationData represents graduation data from the API
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

// fetchGraduationData returns a channel that contains graduation data
func fetchGraduationData() <-chan GraduationData {
	// create a channel
	result := make(chan GraduationData)

	// create a HTTP client
	var client = &http.Client{}

	// create a graduation data
	// this data will be decoded
	var data GraduationData

	// launch a goroutine
	go func() {
		// create a new HTTP request
		request, err := http.NewRequest(http.MethodGet, BASE_URL, nil)
		if err != nil {
			log.Fatalln("error when creating a request: ", err)
		}

		// send the request
		response, err := client.Do(request)
		if err != nil {
			log.Fatalln("error when sending a request: ", err)
		}

		// close the response body if decoding process is finished
		defer response.Body.Close()

		// decode the response body into the "data" variable
		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			log.Fatalln("error when decoding a response: ", err)
		}

		// assign the graduation data to the channel
		result <- data

		// close the channel
		close(result)
	}()

	// return the channel that contains graduation data
	return result
}

func main() {
	// fetch the graduation data
	graduationChannel := fetchGraduationData()

	// create a receiver to receive the graduation data
	receiver := make(chan GraduationData)

	// receive the graduation data from the
	// graduationChannel
	go receive(graduationChannel, receiver)

	// create a new CSV file
	csvFile, err := os.Create("graduation.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// create a new CSV writer
	// to write the content to the CSV
	csvwriter := csv.NewWriter(csvFile)

	// iterate through the received value from the receiver channel
	for received := range receiver {
		// iterate through the graduation data
		for _, graduation := range received.Result.Records {
			var data []string = []string{
				strconv.Itoa(graduation.ID),
				graduation.Sex,
				graduation.TypeOfCourse,
				graduation.NoOfGraduates,
				graduation.Year,
			}

			// write the graduation data to the CSV file
			_ = csvwriter.Write(data)
		}
	}

	// close the writer and the CSV file
	// to avoid resources leaks
	csvwriter.Flush()
	csvFile.Close()
}

// receive receives a value from the channel
func receive(channel <-chan GraduationData, receiver chan GraduationData) {
	// create a WaitGroup
	var wg sync.WaitGroup

	// add one WaitGroup
	wg.Add(1)

	// launch a goroutine
	// to retrieve a value from the channel
	go func() {
		// iterate through every value inside the channel
		for n := range channel {
			// assign the value to the receiver channel
			receiver <- n
		}
		// the operation is done
		wg.Done()
	}()

	// wait until the operation inside the goroutine is finished
	wg.Wait()

	// close the receiver channel
	close(receiver)
}
