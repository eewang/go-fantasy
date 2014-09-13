package main

import (
  "log"
  
  "github.com/PuerkitoBio/goquery"
)

var rotoworld_url = "http://www.rotoworld.com/playernews/nfl/football-player-news"

func Scrape() *goquery.Document {
  doc, err := goquery.NewDocument(rotoworld_url)
  if err != nil {
    log.Fatal(err) 
  }

  return doc
}
