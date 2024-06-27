# GOkart: the Mario kart game analyzer

This project uses a a video feet from an capturecard to analyze an gather data about mario kart 8 deluxe games. 

## Setup
You need to install the C libraries `opencv` for gocv and `tesseract` for goseract on your operation system. Create a Database and create the .env out of the example. Run with `go run .` after connection the capture card.

### For mac setup this helped:
- export LIBRARY_PATH="/opt/homebrew/lib"
- export CPATH="/opt/homebrew/include"


## Features
The Current status of the features is:
- [x] Detect race track name
- [x] Detect Player names
- [x] Detect Player Placements
- [x] Detect Player lap times 
- [x] Detect live player positions
- [ ] Item detection
- [x] Database observer to persist data 
- [x] Logger observer to log event in console
- [ ] Websocket observer
- [ ] REST API

## API
TODOs for the API
- Games
  - [x] Index - /games - list of all games -> all games are shown with date, players and final placements
  - [x] Show - /games/:id - game detail
  - [x] Current - /games/current - current/active game details
  - [x] Player - PATCH /games/:game_id/:number/player - attach persons to player or adapt character
- Round
  - [x] Show - /rounds/:id - details to a round
- Person
  - [x] Index - /persons - list all persons
  - [x] Show - /persons/:id - show person detail
  - [x] Create - POST /persons - create person
  - [ ] Update - PATCH /person/:id - update person
  - [ ] Delete - DELETE /person/:id - delete person
- Character
  - [ ] Index - /characters - list all characters
  - [ ] Show - /characters/:id - show character detail
  - [ ] Persons - /characters/:id/persons - list of all persons which have character as default
- Track
  - [ ] Index - /tracks - list of all tracks
  - [ ] Show - /tracks/:id - stats and details to a track
- Stats
  - [ ] Leaderboard - /stats/leaderboard
  - ...
- Websocket

## API-Instructions:


### Patch player-names for any game
PATCH /games/:game_id/:number/player
- adjust :game_id and :number according to which player you was at the game /table players
- to set in data:
  - fallback_name
  - character_id
  - person_id
example:
curl --location --request PATCH 'http://localhost:8888/api/games/2/4/player' \
--header 'Content-Type: application/json' \
--data '{
  "fallback_name": "tanuki",
  "character_id": 1,
  "person_id": 1
}'


### Patch player-names for current game
PATCH /games/current/:number/player
- adjust :number according to which playernr. you are at the current game
- to set in data:
  - fallback_name
  - character_id
  - person_id

example:
curl --location --request PATCH 'http://localhost:8888/api/games/current/4/player' \
--header 'Content-Type: application/json' \
--data '{
  "fallback_name": "tanuki",
  "character_id": 1,
  "person_id": 1
}'

### Create Person
POST /persons/
- set "name" and "CharacterID" in json-Body

example:
{
    "name": 3,
    "CharacterID": 3
}


### Update Person
PATCH /persons/:id
- set "name" and "CharacterID" in json-Body

example:
{
    "name": 1,
    "CharacterID": 1
}