package main

type Direction string

const (

    Right   Direction = "Right"
    Left    Direction = "Left"
    Down    Direction = "Down"
    Up      Direction = "Up"
    
)

type player struct {

    id      string
    name    string
    vel
    heading position
    snake   []position
   // socket Socket
    color   string

}

type vel struct {
       
    x float64
    y float64
    
}

type gamestate struct {

    players     map[string]player
    foods       Food[]
    gridsize    float64

}

type position struct {

    x float64
    y float64

}

type Food struct {

    pos      position
    bornTime float64

}

type userinput struct {

    playerName string

    Direction Direction
    
}
