package player 
/*
import (

    "math/rand"
    "strconv"
)

var botColors = [6]string{"#ffe6f7", "#bf73ff", "#36badf","#08c96b","#f0dbb7","#beff11"}
var botColorDefault string = "purple"

func ColorsAreCloseForBots(color1 string, color2 string) bool {
    
    threshold := 50

	const r1 = parseInt(color1.replace("#","").substring(0,2),16)
	const g1 = parseInt(color1.replace("#","").substring(2,4),16)
    const b1 = parseInt(color1.replace("#","").substring(4,6),16)
	const r2 = parseInt(color2.replace("#","").substring(0,2),16)
	const g2 = parseInt(color2.replace("#","").substring(2,4),16)
    const b2 = parseInt(color2.replace("#","").substring(4,6),16)

    return ((r1 - r2) * (r1 - r2) + (b1 - b2) * (b1 - b2) + (g1 - g2) * (g1 - g2)) <= threshold * threshold

}

type PlayerBot struct {

    id                  string
    game                Game
    renderState         IPlayerState
    heading             Position
    hasInput            bool
    velocity            Vel 
    noTurnSteps         float64
}

func getNameForBot(bot *PlayerBot) string {

    return bot.renderState.name

}

func getSnakeForBot(bot *PlayerBot) string{
            
    return bot.renderState.snake

}

func randomColorForBot(state *Gamestate, bot *PlayerBot) string {

    randomNum := rand.Float64() * 255 * 255 * 255
    color := "#" + strconv.FormatInt(math.floor(randomNum), 16)
    
    for _, b := range state.Bots {

        if b.id == bot.id {
            continue
        }
        
        if ColorsAreClose(color, b.renderState.color) {

            return randomColor(state, bot)
        }

    }

    if ColorsAreClose(color, BG_COLOR) || ColorsAreClose(color, FOOD_COLOR) {
        
        return randomColor(state, bot)
    
    } else{

        return color
    }

}
        

func randomPositionForBot() Position {

    randPosition = Position{

        X: math.floor(rand.Float64() * GRID_SIZE),
        Y: math.floor(rand.Float54() * GRID_SIZE),
    }

    return randPosition

}

func generateBotName() string {

    for {

        num := rand.Intn(10000) // 0–9999
        guestUsername := fmt.Sprintf("guest%04d", num)

        if _, exists := users[username]; !exists {
            return guestUsername
        }
    
    }

}

func newBot(newGame *Game, state *Gamestate) *PlayerBot {

    var b PlayerBot 
    
    b.game := *newGame
    b.heading := randomPosition()
    b.hasInput := false 
    b.velocity := Vel{ Dxdt: 1, Dydt: 0 }
    b.id := generateGuestUsername() 
    b.renderState := IPlayerState{

        name: b.id,
        snake: b.heading,
        color: randomColor(state, b),
    }

    return b 
    
}

func handleInputForBot(direction string, b *PlayerBot){

    if b.hasInput {
        return
    }


    switch(direction) {

        case "Right":

            if b.velocity.Dydt != 0 {
                
                b.velocity = Vel{Dxdt: 1, Dydt: 0}
            }

        case "Left":

            if b.velocity.Dydt != 0 {

                b.velocity = Vel{Dxdt: -1, Dydt: 0}

            }

        case "Up":

            if b.velocity.Dxdt != 0 {

                b.velocity = Vel{Dxdt: 0, Dydt: -1}

            }

        case "Down":

            if b.velocity.Dxdt != 0 {

                b.velocity = Vel{Dxdt: 0, Dydt: -1}

            }

    }

}

func moveBot(b *PlayerBot){

    if !PositionEqual(b.heading, b.snake[len(b.snake) - 1]) {

        b.snake = append(b.snake, b.heading)
        b.renderState.snake = append(b.renderState.snake[:1], b.renderState.snake[1:])

    }

    b.heading.X := b.velocity.Dxdt
    b.heading.Y := b.velocity.Dydt

    b.hasInput = false

    if b.noTurnSteps < 5 {

        noTurnSteps++
    } else{

        i := math.floor(rand.Float64() * 4)
        directs := [5]string{"Left", "Right", "Up", "Down"}
        handleInput(directs[i], b)
    }
        
}

func checkHitBoundryForBot(b *PlayerBot) bool {

   if b.heading.X < 0 || b.heading.X > GRID_SIZE - 1 || b.heading.Y < 0 || b.heading.Y > GRID_SIZE - 1 {

        return true

   } else {

        return false

   }

}

func checkHitFoodsForBot(foods []Food, b *PlayerBot){

   for k:= 0; k < len(foods); {

       food := foods[k]

       if canEatFood(food) {

            b.snake = append(b.snake, b.heading)
            foods = append(foods[:1], foods[k+1:])
            return
        } else {
            k++
        }
    }

}

func checkHitBot(b *PlayerBot) bool {

    for _, cell := range b.snake {

        if PositionEqual(b.heading, cell) {
            return true
        }

    }

    return false

}

func convertToFoodsForBot(existingFoods []Food, b *PlayerBot) []Food{

    result = make([]Food, 0)

    for i := 0; i < len(b.renderState.snake); i++ {

        cell := b.renderState.snake[i]

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

func canEatFoodForBot(food Food, b *PlayerBot) bool {

    if math.Abs(food.Pos.X - b.Heading.X) <= 1 && math.Abs(food.Pos.Y - b.Heading.Y) <= 1 {

        return true
    }
    return false

}













*/
