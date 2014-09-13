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

// func randString() string {
//   size := 12

//   rb := make([]byte, size)
//   _, err := rand.Read(rb)

//   if err != nil {
//     fmt.Println(err)
//   }

//   rs := base64.URLEncoding.EncodeToString(rb)

//   return rs
// }

// func tokenRequest() string {
//   params := map[string]string{
//     "oauth_consumer_key": yahoo_key,
//     "oauth_nonce": randString(),
//     "oauth_signature_method": "plaintext",
//     "oauth_signature": yahoo_secret,
//     "oauth_timestamp": "1410580587",
//     "oauth_version": "1.0",
//     "oauth_callback": "http://127.0.0.1:8001/oauth/callback",
//   }

//   request := yahoo_token_url + "?"

//   for k, v := range params {
//     request +="&"
//     request += k
//     request += "="
//     request += v
//   }

//   fmt.Printf("%v\n", request)

//   return request
// }
