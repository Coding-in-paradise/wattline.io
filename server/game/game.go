package game 

import (
   
    "github.com/coder/websocket"
    "wattline/server/models"
    "net/http"
    "wattline/server/constants"
    "log"
    "fmt"
    "errors"
    "encoding/json"
    "time"
    "math/rand"
    "context"
    "math"
)

var colors = []string{"#ffe6f7", "#bf73ff", "#36badf","#08c96b","#f0dbb7","#beff11"}

type Message struct {
    
    Event       string            `json:"event"`
    Data        json.RawMessage   `json:"data"`
}

type Game struct {

    State       models.Gamestate
    Server      *http.Server 
    Clients     map[string]*websocket.Conn
    
}

func NewGame() *Game {

    http.HandleFunc("/ws", createClientWebsocket)

    return &Game{

        State:           models.Gamestate {

            Players:     make(map[string]*models.Player),
            Foods:       []models.Food {
                
                {

                     Pos: models.Position{X: 1, Y: 1}, 
                     BornTime: float64(time.Now().Unix()), 
                },
        
            },

            Gridsize:    constants.GRID_SIZE,

        },

        Server:     &http.Server{}, 
        Clients:    make(map[string]*websocket.Conn),
    }

}

func createClientWebsocket(w http.ResponseWriter, r *http.Request) {
    
    remoteAddr := r.RemoteAddr

    client, err := websocket.Accept(w, r, &websocket.AcceptOptions{
        InsecureSkipVerify: true,
        //originpatterns: []string{"https://wattline.com"},
    })

    if err != nil {
        log.Printf("WebSocket accept failed: %v", err)
        return
    }

    log.Println("WebSocket connection established from: ", remoteAddr)

    newData := models.Userinput{}
    newGame := NewGame()
    newPlayer := models.Player{}

    go handleOnConnect(client, &newData, newGame, &newPlayer, remoteAddr)
}

func StartGameInterval(game *Game){

    fmt.Println("startGameInterval")

    ticker := time.NewTicker(time.Second / constants.FRAME_RATE)

    defer ticker.Stop()

    for {
        <-ticker.C
        gameLoop(game)
    }

}

func handleOnConnect(client *websocket.Conn, datainput *models.Userinput, game *Game, player *models.Player, remoteAddr string) {

    defer client.Close(websocket.StatusNormalClosure, "closing")
	
    log.Println("Client connected: ", remoteAddr)

    for {
        ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
        
        defer cancel()
        
        var msg Message

         _, data, err := client.Read(ctx)

        msg.Data = json.RawMessage(data)
        
        if err != nil {
            log.Fatal("read:", err)
        }
        
        var playerName string

        playerID := remoteAddr + playerName

        switch msg.Event {

            case "joinGame":
                
                if err := json.Unmarshal(msg.Data, &playerID); err != nil {
                    log.Printf("Failed to parse joinGame payload: %v", err)
                    continue
                }
          
                addPlayer(client, game, playerName, playerID)

            case "disconnection":
                
                log.Println("Client disconnected: ", remoteAddr)
                
                delete(game.Clients, playerID)
                delete(game.State.Players, playerID)

            case "turn":

                handlePlayerInput(client, datainput, game, player)

            default:
                log.Printf("Unknown event: %s", msg.Event)
        }

    }

}

func handlePlayerInput(client *websocket.Conn, data *models.Userinput, game *Game, player *models.Player) error {

	if _, ok := game.State.Players[data.PlayerID]; !ok {

        err := errors.New("Player not found")
        return err
    }
		
    switch(data.Direction) {
		
        case "Right":
			
            if player.Vel.Dydt != 0 {
				player.Vel = models.Vel{ Dxdt: 1, Dydt: 0 }
			}
			break
		
        case "Left":
			if player.Vel.Dydt != 0 {
				player.Vel = models.Vel{ Dxdt: -1, Dydt: 0 }
			}
			break 
		
        case "Up":
			if player.Vel.Dxdt != 0 {
				player.Vel = models.Vel{ Dxdt: 0, Dydt: -1 }
			}
			break 
		
        case "Down":
			if player.Vel.Dxdt != 0 {
				player.Vel = models.Vel{ Dxdt: 0, Dydt: 1 }
			}
			break 
		}

    return nil
}

func addPlayer(clientConnection *websocket.Conn, game *Game, name, playerID string) {

        log.Printf("Add player ID: %s\n", playerID)
		
        if _, ok := game.State.Players[playerID]; !ok {

			colorPicked := colors[rand.Intn(len(colors))]

            var newHeading = models.Position{

                X: 10,
                Y: 10,

            }

            var newVel = models.Vel{

                Dxdt: 1,
                Dydt: 0,
            }

            newSnake := make([]models.Position, 0)
            
            newSnake = append(newSnake, models.Position{X: 10, Y: 10}) 
			
            var newPlayer = models.Player{
				Id:                 playerID,
				PlayerName:         name,
                ClientConnection:   clientConnection,
				Heading:            newHeading,
				Vel:                newVel,
				Snake:              newSnake,
				Color:              colorPicked,
            }
			
            game.State.Players[playerID] = &newPlayer
		}
		
        game.Clients[playerID] = clientConnection

}

func generateFood(game *Game) {

        var foodPosition = models.Position{

            X: float64(rand.Intn(constants.GRID_SIZE)),
            Y: float64(rand.Intn(constants.GRID_SIZE)),
        }

        var newFood = models.Food{

            Pos: foodPosition,
            BornTime: float64(time.Now().Unix()),
        }

        for _, player := range game.State.Players {

            for _, cell := range player.Snake {

                if cell.X == newFood.Pos.X && cell.Y == newFood.Pos.Y {
                    generateFood(game)
                    return
                }

            }

        }

        for _, food := range game.State.Foods {

            if food.Pos.X == newFood.Pos.X && food.Pos.Y == newFood.Pos.Y {
                
                generateFood(game)
                return
            }

        }
        
        game.State.Foods = append(game.State.Foods, newFood)
		
}

func checkHitBoundry(player *models.Player) bool {
	
    if player.Heading.X < 0 || player.Heading.X > constants.GRID_SIZE - 1 || player.Heading.Y < 0 || player.Heading.Y > constants.GRID_SIZE - 1 {
		return true
	}
		return false
}

func checkHitFood(player *models.Player, game *Game) {

    for i:= 0; i < len(game.State.Foods); i++ {

        food := game.State.Foods[i]
        
        if math.Abs(food.Pos.X - player.Heading.X) <= 1 && math.Abs(food.Pos.Y - player.Heading.Y) <= 1 {

            player.Snake = append(player.Snake, player.Heading)
            player.Heading.X += player.Vel.Dxdt
            player.Heading.Y += player.Vel.Dydt
            game.State.Foods = append(game.State.Foods[:i], game.State.Foods[i+1:]...)
            continue
        }

    }
}

func PositionEqual(player *models.Player, game *Game) bool {

    for _, other := range game.State.Players {

        for _, cell := range other.Snake {

            if player.Heading.X == cell.X && player.Heading.Y == cell.Y {
                return true
            }

        }

    }

    return false

}

func convertToFoods(player *models.Player, game *Game){

    for i:= 0; i < len(player.Snake); i++{

        cell := player.Snake[i]

        hasFoodHere := false

        for _, food := range game.State.Foods {

            if food.Pos.X == cell.X && food.Pos.Y == cell.Y {
                hasFoodHere = true
                break
            }
        }
            
        if !hasFoodHere {

            var foodPos = models.Position{

                X: cell.X,
                Y: cell.Y,
            }
                
            var newFood = models.Food{

                Pos: foodPos,
                BornTime: float64(time.Now().Unix()),
            }
                
            game.State.Foods = append(game.State.Foods, newFood)

        }

    }

}

func gameLoop(game *Game){

        for playerID, player := range game.State.Players {

            player.Heading.X += player.Vel.Dxdt
            player.Heading.Y += player.Vel.Dydt

            if checkHitBoundry(player) || PositionEqual(player, game) {

                convertToFoods(player, game)
                delete(game.State.Players, playerID)             
                continue
            }

            checkHitFood(player, game)
            player.Snake = append(player.Snake, player.Heading)
            player.Snake = player.Snake[1:]

        }

        currentTime := time.Now()

	    // Convert the current time to seconds since the Unix epoch
	    seconds := currentTime.Unix()

        for i := 0; i < len(game.State.Foods); i++{

            food := game.State.Foods[i]

            if ( float64(seconds) - food.BornTime) > (30 * 1000) {

                game.State.Foods = append(game.State.Foods[:i], game.State.Foods[i+1:]...)
                i-- 
            }

        }

        if len(game.State.Foods) < 10 {

            generateFood(game)

        }

}

