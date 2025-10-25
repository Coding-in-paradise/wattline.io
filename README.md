# wattline.io

Wattline.io is an alternative to the original powerline.io game with the server re-written in Golang with the chief goal to decrease latency. 

Wattline has several features that distinguish it from the original powerline.io game, such as:

    The backend server is programmed in Golang. 

    A profanity filter to prevent people from having usernames that are vulgar. This is to make the game family-friendly. 

    Add other restrictions to usernames like no special characters, and no white spaces.

    Add a leaderboard for the top 25 players in each category: K/D, best score, best kill streak 

    No gameplay chat.

An auxiliary goal of this project is to use performance benchmarks to see the difference between deploying the server logic in TS/JS vs Golang. 

