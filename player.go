package main

import (
  "github.com/PuerkitoBio/goquery"
)

type playerSpec struct {
  Name string `json:"name"`
  Team string `json:"team"`
  Report string `json:"report"`
  Date string `json:"date"`
}

func BuildPlayer(selection *goquery.Selection) playerSpec {
  name    := getName(selection)
  team    := getTeam(selection)
  date    := getDate(selection)
  report  := getReport(selection)

  return playerSpec{name, team, report, date}
}

// Private methods

func getName(selection *goquery.Selection) string {
  return selection.Find(".player a").First().Text()
}

func getTeam(selection *goquery.Selection) string {
  return selection.Find(".player a").Last().Text()
}

func getDate(selection *goquery.Selection) string {
  return selection.Find(".info .date").Text()
}

func getReport(selection *goquery.Selection) string {
  return selection.Find(".report p").Text()
}
