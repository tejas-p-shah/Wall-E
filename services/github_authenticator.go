package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Utils

func getGithubClientID() string {
	githubClientID, exists := os.LookupEnv("CLIENT_ID")
	if !exists {
		log.Fatal("Could not Found CLIENT_ID for Github in .env file")
	}

	return githubClientID
}

func getGithubClientSecret() string {
	githubClientSecret, exists := os.LookupEnv("CLIENT_SECRET")
	if !exists {
		log.Fatal("Could not Found CLIENT_SECRET for Github in .env file")
	}

	return githubClientSecret
}

func getGithubAccessToken(code string) string {
	clientID := getGithubClientID()
	clientSecret := getGithubClientSecret()

	requestBodyMap := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
	}
	requestJSON, _ := json.Marshal(requestBodyMap)

	req, reqerr := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)
	if reqerr != nil {
		log.Panic("Request Connection Failed")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request Failed")
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	var ghresp githubAccessTokenResponse
	json.Unmarshal(respbody, &ghresp)

	return ghresp.AccessToken
}

func getGithubData(accessToken string) string {
	req, reqerr := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if reqerr != nil {
		log.Panic("API Request creation failed")
	}

	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request Failed")
	}

	respbody, _ := ioutil.ReadAll(resp.Body)
	return string(respbody)
}

// Handler functions

func loggedinHandler(w http.ResponseWriter, r *http.Request, githubData string) {
	if githubData == "" {
		fmt.Fprintf(w, "UNAUTHORIZED!!")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var prettyJson bytes.Buffer

	parserr := json.Indent(&prettyJson, []byte(githubData), "", "\t")
	if parserr != nil {
		log.Panic("JSON parse error!")
	}

	fmt.Fprintf(w, string(prettyJson.Bytes()))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<a href="/login/github/">LOGIN</a>`)
}

func gitgubLoginHandler(w http.ResponseWriter, r *http.Request) {
	githubClientID := getGithubClientID()
	// port := getPort()

	redirectURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		githubClientID,
		"http://localhost:8099/login/github/callback",
	)

	http.Redirect(w, r, redirectURL, 301)
}

func githubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	githubAccessToken := getGithubAccessToken(code)

	githubData := getGithubData(githubAccessToken)

	loggedinHandler(w, r, githubData)

}

func Init_Github_Authenticator() {

	http.HandleFunc("/login/github/", gitgubLoginHandler)

	http.HandleFunc("/login/github/callback", githubCallbackHandler)

	http.HandleFunc("/loggedin", func(w http.ResponseWriter, r *http.Request) {
		loggedinHandler(w, r, "")
	})

}
