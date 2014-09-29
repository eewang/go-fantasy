package main

import (
  // "fmt"
)
// Stub for the response from the Yahoo! Fantasy Sports API

type Roster struct {
  name string
}

func wide_receivers() []string {
  wrs := []string{
    "Alshon Jeffrey",
    "Jordy Nelson",
    "Calvin Johnson",
    "Roddy White",
  }
  return wrs
}

func CreateRoster(name string) Roster {
  return Roster{name}
}

func playerMatches(roster string, news string) bool {
  return roster == news
}

//
// Returns a slice of players from the roster that match the scraped player news
//
func rosterMatch() []playerSpec {
  matched_players := make([]playerSpec, 0)

  for _, player := range retrievePlayers() {
    for _, receiver := range wide_receivers() {
      if playerMatches(receiver, player.Name) {
        matched_players = append(matched_players, player)
      }
    }
  }

  return matched_players
}