# go-fortnite

fortnite tracker API client

## Install

```bash
go get -u github.com/usk81/go-fortnite
```

## Usage

``` go
// create new instance
s, _ := New(nil, "your access token")

// Get Fornite BR Player Stats
//   e.g.
//   Platform: gamepad (nintendo switch)
//   nickname: Shady Grove
result, err = s.BRPlayerStats("gamepad", "Shady Grove")

// Get Fornite Match History
result, err = s.MatchHistory("player account id")

// Get Current Fortnite Store
result, err = s.Store()

// Get Current Active Challenges
//   note: This API is inactive, maybe.
result, err = s.Challenges()
```
