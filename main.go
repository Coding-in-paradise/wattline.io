package main 

import (
    "log"
    "net/http"
    "wattline/server/game"
)

func main() {
    
    log.SetFlags(0)
    
    newgame := game.NewGame()

    go game.StartGameInterval(newgame)
    
    log.Println("Starting server on :8080")
    
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
} 
   /*
    http.HandleFunc("/ws", wsHandler)
    http.HandleFunc("/login", auth.loginHandler)
    http.HandleFunc("/register", auth.registerHandler)
    http.HandleFunc("/leaderboard", auth.leaderboardHandler)
    
    if err := auth.LoadBadWords("~/workspace/github.com/Coding-in-paradise/wattline.io/server/src/auth/badwords.txt"); err != nil {
        log.Fatalf("Failed to load bad words: %v", err)
    }

    err := run()

    if err != nil {
        log.Fatal(err)
    }
}

func run() error {
    
    game := Game{}

    game.startGameInterval()

    log.Fatal(http.ListenAndServe(":8080", nil))
}
    */
