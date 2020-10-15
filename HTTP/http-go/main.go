package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	httpMethod := "GET"
	//url := "https://api.github.com"

	url := "http://localhost:8080"

	client := http.Client{}

	//res, err := client.Get(url)

	request, err := http.NewRequest(httpMethod, url, nil)

	request.Header.Set("Accept", "application/xml")

	res, err := client.Do(request)

	if err != nil {

		panic(err)
	}

	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {

		panic(err)
	}

	fmt.Println(res.StatusCode)

	fmt.Println(string(bytes))
}
