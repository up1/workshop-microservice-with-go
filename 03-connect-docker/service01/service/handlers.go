package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Result string `json:"result"`
}

func CallService02(w http.ResponseWriter, r *http.Request) {

	// Call service 2
	url := "http://service02:9002/hello"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
