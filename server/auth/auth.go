package auth 

import (
 
  "golang.org/x/crypto/bcrypt"
   crand "crypto/rand"
  "encoding/hex"
  "strings"
  "os"
  "encoding/json"
  "fmt"
  "math/rand"
)

func hashPassword(password string) ([]byte, error) {
    return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func checkPassword(hash []byte, password string) error {
    return bcrypt.CompareHashAndPassword(hash, []byte(password))
}

func generateGuestUsernameForAuth(Users map[string]User) string {

    for {

        num := rand.Intn(10000) // 0–9999
        guestUsername := fmt.Sprintf("guest%04d", num)

        if _, exists := Users[guestUsername]; !exists {
            return guestUsername
        }
    
    }

}

func generateSessionToken() (error, string) {
    
    b := make([]byte, 32)

    _, err := crand.Read(b)

    if err != nil {

        return err, ""
    
    }
    
    return nil, hex.EncodeToString(b)
}

func LoadBadWords() (error, []string) {

	file, err := os.ReadFile("~/workspace/github.com/Coding-in-paradise/wattline.io/server/src/auth/badwords.txt")
	
    if err != nil {
		return err, nil
	}

	var badWords []string
	
    if err := json.Unmarshal(file, &badWords); err != nil {
		return err, nil
	}

	return nil, badWords
}

func isValidUsername(username string) (error, bool) {

    if username == "" {
        return nil, false
    }

    numberString := "1234567890"
    specialCharacterString := "`~!@#$%^&*()_+-=[]\\{}|;':\"/.,<>?"

    for i := 0; i < len(specialCharacterString); i++{

        if strings.Contains(username, string(specialCharacterString[i])) {
            return nil, false
        }

    }

    for i := 0; i < len(numberString); i++{
        
        if strings.Contains(username, string(numberString[i])) {
            return nil, false
        }
        
    }

    err, badWords := LoadBadWords()

    if err != nil {

        return err, false

    }

    for i := 0; i < len(badWords); i++{

        if strings.Contains(strings.ToLower(username), badWords[i]){

            return nil, false
        }
    }

    return nil, true
}
