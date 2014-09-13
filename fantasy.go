package main

import (
  "fmt"
  "log"
  "net/http"
  "flag"

  "github.com/gorilla/mux"
  "github.com/PuerkitoBio/goquery"
)

var host = flag.String("host", "www.gofantasy.com", "the host to listen on")
var port = flag.Int("port", 8000, "the port to listen on")

func retrievePlayers() []playerSpec {
  players := []playerSpec{}

  Scrape().Find(".pb").Each(func(i int, s *goquery.Selection) {
    players = append(players, BuildPlayer(s))
  })

  return players
}

func main() {
  flag.Parse()
  if err := readCredentials(); err != nil {
    log.Fatalf("Error reading configuration, %v", err)
  }
  fmt.Printf("%v\n", yahooAuthClient.Credentials)
  r := mux.NewRouter()
  r.HandleFunc("/", rootHandler)
  r.HandleFunc("/authorize", authorizeHandler)
  r.HandleFunc("/oauth/callback", oauthCallbackHandler)
  r.HandleFunc("/players.json", playersHandler)
  http.Handle("/", r)

  fmt.Printf("%v\n", yahooAuthClient.TemporaryCredentialRequestURI)

  addr := fmt.Sprintf("%s:%v", *host, *port)
  log.Printf("Listening on http://%s\n", addr)
  listenErr := http.ListenAndServe(addr, nil)

  if listenErr != nil {
    panic(listenErr)
  }
}
