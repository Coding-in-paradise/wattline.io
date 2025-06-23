package main

import (
	"math/rand"
)

func makeid(length int) string {
	
    result := ""
	
    const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	const charactersLength = len(characters)

	for i := 0; i < length; i++ {
        
        index := math.floor(rand.Float64() * charactersLength)
        result += characters[index]
	
    }

	return result
}
