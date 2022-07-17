// Calling GitHub API
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	NumRepos int    `json:"public_repos"` // from public_repos in JSON
}

// userInfo return information on github user
func userInfo(login string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com/users/"+login, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Can't get from the given URL")
	}
	defer resp.Body.Close()

	var user User
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&user); err != nil {
		log.Fatal(err)
	}
	return &user, nil
}

func main() {
	user, err := userInfo("aditsawant")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%#v\n", user)
}
