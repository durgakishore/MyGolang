package main

import (
	"Projects/PublicGitRepo/http-client/gohttp"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	githubHTTPClient = getGithubClient()
)

func getGithubClient() gohttp.HTTPClient {

	client := gohttp.NewHTTPClient()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)

	return client
}
func main() {

	headers := make(http.Header)

	//headers.Set("Authorization", "Bearer ABC-123")

	res, err := githubHTTPClient.Get("https://api.github.com", headers)

	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)

	bytes, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(bytes))
}
