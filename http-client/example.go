package main

import (
	"Projects/PublicGitRepo/http-client/gohttp"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	githubHTTPClient = getGithubClient()
)

func getGithubClient() gohttp.HTTPClient {

	client := gohttp.NewHTTPClient()

	//commonHeaders := make(http.Header)
	//commonHeaders.Set("Authorization", "Bearer ABC-123")
	//client.SetHeaders(commonHeaders)

	client.SetMaxIdleConnections(20)
	client.SetConnectionTimeout(2 * time.Second)
	client.SetRequestTimeout(50 * time.Millisecond)

	return client
}

//User struct
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {

	headers := make(http.Header)

	headers.Set("Authorization", "Bearer ABC-123")

	res, err := githubHTTPClient.Get("https://api.github.com" /*headers*/, nil)

	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)

	bytes, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(bytes))
}

func createUser(user User) {
	//res, err := githubHTTPClient.Post("https://api.github.com", nil, user)

}
