package models 

import ( 

    "github.com/coder/websocket"
)

type Direction string

const (
 
     Right   Direction = "Right"
     Left    Direction = "Left"
     Down    Direction = "Down"
     Up      Direction = "Up"
)

type User struct {
    
    Username        string
    JoinDate        string
    Rank            int 
    MaxScore        int
    KingsKilled     int
    NumOfKills      int
    NumOfDeaths     int
    Rating          int
    BestKillStreak  int
    PasswordHash    []byte

}

type IPlayerState struct {

    Name    string
    Snake   []Position
    Color   string
}

type Player struct {

    Id                  string
    ClientConnection    *websocket.Conn 
    PlayerName          string
    Heading             Position
    Snake               []Position
    Color               string
    Vel

}

type Bot struct {

    Id              string            
    BotName         string
    Heading         Position
    Snake           []Position
    Color           string
    Vel

}

type Vel struct {
       
    Dxdt float64
    Dydt float64
    
}

type Gamestate struct {

    Players     map[string]*Player
    Bots        map[string]*Bot
    Foods       []Food
    Gridsize    float64

}

type Position struct {

    X float64
    Y float64

}

type Food struct {

    Pos         Position
    BornTime    float64

}

type Userinput struct {

    PlayerID    string
    Direction   Direction
}
