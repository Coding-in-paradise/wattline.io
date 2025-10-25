package player 
/*
import (

    "math/rand"
    "strconv"
)

var playerColors = [6]string{"#ffe6f7", "#bf73ff", "#36badf","#08c96b","#f0dbb7","#beff11"}
var playerColorDefault string = "purple"

func ColorsAreCloseForPlayers(color1, color2 string) bool {
    
    threshold := 50

	const r1 = parseInt(color1.replace("#","").substring(0,2),16)
	const g1 = parseInt(color1.replace("#","").substring(2,4),16)
    const b1 = parseInt(color1.replace("#","").substring(4,6),16)
	const r2 = parseInt(color2.replace("#","").substring(0,2),16)
	const g2 = parseInt(color2.replace("#","").substring(2,4),16)
    const b2 = parseInt(color2.replace("#","").substring(4,6),16)

    return ((r1 - r2) * (r1 - r2) + (b1 - b2) * (b1 - b2) + (g1 - g2) * (g1 - g2)) <= threshold * threshold

}

type Player struct {

    id                  string
    game                Game
    renderState         IPlayerState
    heading             Position
    hasInput            bool
    velocity            Vel 
    socket              *websocket.Conn
    noTurnSteps         float64
}

func getNameForPlayer(p *Player) string {

    return p.renderState.name

}

func getSnakeForPlayer(p *Player) string{
            
    return p.renderState.snake

}

func randomColorForPlayer(state *Gamestate, p *Player) string {

    randomNum := rand.Float64() * 255 * 255 * 255
    color := "#" + strconv.FormatInt(math.floor(randomNum), 16)
    
    for _, player := range state.players {

        if player.id == p.id{ 
            continue
        }
    }

    if ColorsAreClose(color, BG_COLOR) || ColorsAreClose(color, FOOD_COLOR) {
        
        return randomColor(state, p)
    
    } else{

        return color
    }
        
}

func randomPositionForPlayer(){

    randPosition = Position{

        X: math.floor(rand.Float64() * GRID_SIZE),
        Y: math.floor(rand.Float54() * GRID_SIZE),
    }

    return randPosition

}

func generateGuestUsername() string {

    for {

        num := rand.Intn(10000) // 0–9999
        guestUsername := fmt.Sprintf("guest%04d", num)

        if _, exists := users[username]; !exists {
            return guestUsername
        }
    
    }

}

func newPlayer(newGame *Game, clientConnection *websocket.Conn, playerName string, r *http.Request, isGuest bool) *Player {

    remoteAddr := r.RemoteAddr

    var p Player
    
    p.game := *newGame
    p.socket := clientConnection
    p.heading := randomPosition()
    p.hasInput := false 
    p.velocity := Vel{ Dxdt: 1, Dydt: 0 }

    var playername string

    if isGuest {

        p.id := generateGuestUsername() + remoteAddr
        playername = generateGuestUsername()    
    
    } else {

        p.id := playerName + remoteAddr
        playername = playerName
    }

    p.renderState := IPlayerState{

        name: playername,
        snake: p.heading,
        color: randomColor(state, p),
    }
    
}

func handleInputForPlayer(direction string, p *Player){

    if p.hasInput {
        return
    }


    switch(direction) {

        case "Right":

            if p.velocity.Dydt != 0 {
                
                p.velocity = Vel{Dxdt: 1, Dydt: 0}
            }

        case "Left":

            if p.velocity.Dydt != 0 {

                p.velocity = Vel{Dxdt: -1, Dydt: 0}

            }

        case "Up":

            if p.velocity.Dxdt != 0 {

                p.velocity = Vel{Dxdt: 0, Dydt: -1}

            }

        case "Down":

            if p.velocity.Dxdt != 0 {

                p.velocity = Vel{Dxdt: 0, Dydt: -1}

            }

    }

}

func movePlayer(p *Player){

    if !PositionEqual(p.heading, p.snake[len(p.snake) - 1]) {

        p.snake = append(p.snake, p.heading)
        p.renderState.snake = append(p.renderState.snake[:1], p.renderState.snake[1:])

    }

    p.heading.X := p.velocity.Dxdt
    p.heading.Y := p.velocity.Dydt

    p.hasInput = false

}

func checkHitBoundryForPlayer(p *Player) bool {

   if p.heading.X < 0 || p.heading.X > GRID_SIZE - 1 || p.heading.Y < 0 || p.heading.Y > GRID_SIZE - 1 {

        return true

   } else {

        return false

   }

}

func checkHitFoodsForPlayer(foods []Food, p *Player){

   for k:= 0; k < len(foods); {

       food := foods[k]

       if canEatFood(food) {

            p.snake = append(p.snake, p.heading)
            foods = append(foods[:1], foods[k+1:])
            return
        } else {
            k++
        }
    }

}

func checkHitPlayer(p *Player) {

    for _, cell := range p.snake {

        if PositionEqual(p.heading, cell) {
            return true
        }

    }

    return false

}

func convertToFoodsForPlayer(existingFoods []Food, p *Player){

    result = make([]Food, 0)

    for i := 0; i < len(p.renderState.snake); i++ {

        cell := p.renderState.snake[i]

        hasFoodHere := false

        for _, food := range existingFoods {

            if food.Pos.X == cell.X && food.Pos.Y == cell.Y {

                hasFoodHere = true
            }

        }

        if !hasFoodHere {

            newPos := Position{

                X: cell.X,
                Y: cell.Y,

            }

            newFood := Food{

                Pos: newPos,
                BornTime: time.Now(),

            }

            result = append(result, newFood) 

        }


    }

    return result
}

func canEatFoodForPlayer(food Food, p *Player){

    if math.Abs(food.Pos.X - p.Heading.X) <= 1 && math.Abs(food.Pos.Y - p.Heading.Y) <= 1 {

        return true
    }
    return false

}
*/
