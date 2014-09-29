package main 

import (
  // "fmt"
  // "log"
  // "crypto/rand"
  // "encoding/base64"

  "github.com/garyburd/go-oauth/oauth"
)

var yahoo_auth_url = "https://api.login.yahoo.com/oauth/v2/request_auth"
var yahoo_request_token_url = "https://api.login.yahoo.com/oauth/v2/get_request_token"
var yahoo_get_token_url = "https://api.login.yahoo.com/oauth/v2/get_token"

var test_url = "http://fantasysports.yahooapis.com/fantasy/v2/game/nfl"

var yahooAuthClient = oauth.Client{
  TemporaryCredentialRequestURI:  yahoo_request_token_url,
  ResourceOwnerAuthorizationURI:  yahoo_auth_url,
  TokenRequestURI:                yahoo_get_token_url,
}
