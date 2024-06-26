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
  - [ ] Index - /games - list of all games
  - [ ] Show - /games/:id - game detail
  - [ ] Current - /games/current - current/active game details
  - [ ] Player - PATCH /games/:id/player - attach persons to player or adapt character
- Round
  - [ ] Show - /rounds/:id - details to a round
- Person
  - [ ] Index - /persons - list all persons
  - [ ] Show - /persons/:id - show person detail
  - [ ] Create - POST /persons - create person
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
