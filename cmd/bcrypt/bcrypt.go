package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// for i, args := range os.Args {
	// 	fmt.Println(i, args)
	// }

	// fmt.Println("Length of the args", len(os.Args))

	args := os.Args
	if len(args) != 3 && len(args) != 4 {
		fmt.Println("Please enter a command: hash 'some password here' or compare 'some password here' 'some hash here' ")
		return
	}

	command := args[1]
	switch command {
	case "hash":
		{
			password := args[2]
			hashedPassword := hash(password)
			fmt.Printf("The hash is %q\n", hashedPassword)
		}
	case "compare":
		{
			if len(args) != 4 {
				fmt.Println("Please enter a valid command. Eg: compare 'some password here' 'some hash here'")
				return
			}
			password := args[2]
			hash := args[3]
			validPassword := compareHash(password, hash)
			if validPassword {
				fmt.Println("The password is Valid")
			} else {
				fmt.Println("The password is Not Valid")
			}
		}
	default:
		{
			fmt.Println("Please enter a valid command: hash or compare")
		}

	}
}

func hash(password string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error hashing: %v\n", password)
	}
	return string(hashedBytes)
}

func compareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
