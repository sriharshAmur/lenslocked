package main

import (
	"errors"
	"fmt"
)

func main() {
	err := CreateOrg()
	fmt.Println(err)
}

func Connect() error {
	return errors.New("connection failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		// return err
		return fmt.Errorf("while creating user: %w", err)
	}
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		// return err
		return fmt.Errorf("while creating org: %w", err)
	}
	return nil
}
