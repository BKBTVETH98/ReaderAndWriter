package files

import (
	"fmt"
	"os"
)

func WriteFile(content []byte) {

	file, err := os.Create("account.json")

	if err != nil {
		fmt.Printf("Ошибка - %s", err)
	}
	defer file.Close()
	_, err = file.Write(content)

	if err != nil {
		fmt.Printf("Ошибка - %s", err)
		return
	}

	fmt.Println("Запись - ок")
	fmt.Println()

}

func ReadFile(name string) ([]byte, error) {

	readFile, err := os.ReadFile(name)

	if err != nil {
		fmt.Println("Ошибка чтения ", err)
		return nil, err
	}

	return readFile, nil
}
