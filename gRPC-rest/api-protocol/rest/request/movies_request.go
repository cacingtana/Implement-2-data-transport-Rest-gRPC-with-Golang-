package request

import (
	"encoding/json"
	"fmt"
	"grpc-rest/pkg/models"
	"net/http"
)

var OMDBKEY = "faf7e5bb"

type RequestMovies struct {
}

func (s *RequestMovies) GetAllMovies(page string, search string) models.Info {
	value1 := fmt.Sprintf("%s", OMDBKEY)
	value2 := fmt.Sprintf("%s", search)
	value3 := fmt.Sprintf("%s", page)
	req, err := http.NewRequest("GET", "http://www.omdbapi.com/?apikey="+value1+"&s="+value2+"&page="+value3, nil)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err)
	}
	defer resp.Body.Close()
	var data models.Info
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}
	return data
}

func (s *RequestMovies) GetMoviesById(id string) models.Movies {
	value1 := fmt.Sprintf("%s", OMDBKEY)
	value2 := fmt.Sprintf("%s", id)
	req, err := http.NewRequest("GET", "http://www.omdbapi.com/?apikey="+value1+"&i="+value2, nil)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err)
	}
	defer resp.Body.Close()
	var data models.Movies
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}

	return data
}
