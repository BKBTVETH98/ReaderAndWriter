package main

import (
	account "JsonExample/account"
	"fmt"
)

func main() {

	acc := &account.Account{}

	for {
		err := acc.ControlActions()
		if err != nil {
			fmt.Println(err)
		}
	}
}
