package main

import (
  "net/http"
  "log"
  "fmt"
  "sync"
  "flag"
  "net/url"
  "io"
  "io/ioutil"
  "encoding/json"

  "github.com/garyburd/go-oauth/oauth"
)

const (
  indexHTML = `<html><body>Hello World!</body></html>`
)

var credPath = flag.String("config", "config.json", "Path to config file for Yahoo creds")

var (
  secretsMutex sync.Mutex
  secrets = map[string]string{}
)

func readCredentials() error {
  b, err := ioutil.ReadFile(*credPath)
  if err != nil {
    return err
  }
  return json.Unmarshal(b, &yahooAuthClient.Credentials)
}

func putCredentials(cred *oauth.Credentials) {
  secretsMutex.Lock()
  defer secretsMutex.Unlock()
  secrets[cred.Token] = cred.Secret  
}

func authorizeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("%v\n", r.Host)
  callback := "http://" + r.Host + "/oauth/callback"
  tempCred, err := yahooAuthClient.RequestTemporaryCredentials(http.DefaultClient, callback, nil)
  fmt.Printf("%v\n", tempCred)
  if err != nil {
    http.Error(w, "Error getting temp cred, "+err.Error(), 500)
    return
  }
  putCredentials(tempCred)

  params := url.Values{
    "oauth_callback": {callback},
    "oauth_consumer_key": {yahooAuthClient.Credentials.Token},
  }
  fmt.Printf("%v\n", params)
  http.Redirect(w, r, yahooAuthClient.AuthorizationURL(tempCred, params), 302)
}

func oauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
  code := r.FormValue("code")
  log.Printf("%v\n", code)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html;charset=UTF-8")
  io.WriteString(w, indexHTML)
}

func playersHandler(w http.ResponseWriter, r *http.Request) {
  output := map[string]interface{}{
    "players": retrievePlayers(),
  }

  outputBytes, err := json.Marshal(output)
  if err != nil {
    log.Printf("Error marshalling JSON: %s\n", err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "Application/json;charset=UTF-8")
  w.Write(outputBytes)
}
