package account

import (
	"errors"
	"fmt"
	"os"
)

type Account struct { //структура аккаунта
	login string
	pass  string
}

func (acc *Account) getUser() { //метод вывода структуры используя указатель на структуру
	fmt.Println()
	fmt.Println("Вывод структуры")
	fmt.Printf("login: %s\n", acc.login)
	fmt.Println()
	readFile()
	fmt.Println()
}

func (acc *Account) makeAccount() { //метод создания аккаунта используя указатель на структуру

	fmt.Println()
	fmt.Print("Введи желаемый логин - ")
	fmt.Scan(&acc.login)
	fmt.Print("Введи password - ")
	fmt.Scan(&acc.pass)
	fmt.Println("Успешное  создание аккаунта")
	fmt.Println()
	acc.writeFile()

}

func (acc *Account) deleteUser() {
	fmt.Println("Удаление пользователя из файла и структуры")
	acc.login = ""
	acc.pass = ""
	acc.writeFile()

	*acc = Account{} //очищаем структуру
	fmt.Println("Пользователь удален из структуры и файла")
	acc.getUser() //проверяем что структура очищена
}

func (acc *Account) writeFile() {

	file, err := os.Create("account.txt")

	if err != nil {
		fmt.Printf("Ошибка - %s", err)
	}

	_, err = file.WriteString("login: " + acc.login + "\n" + "pass: " + acc.pass + "\n")

	if err != nil {
		fmt.Printf("Ошибка - %s", err)
		return
	}

	defer file.Close()
	fmt.Println("Запись - ок")
	fmt.Println()

}

func readFile() {

	readFile, err := os.ReadFile("account.txt")

	if err != nil {
		fmt.Println("Ошибка чтения ", err)
		return
	}

	fmt.Printf("Файл содержит \n%s", readFile)
	fmt.Println()
}

func (acc *Account) ControlActions() error {

	fmt.Println()
	fmt.Println("Выберите действие: ")
	fmt.Println("1 - создать аккаунт")
	fmt.Println("2 - вывести структуру и прочитать из файла")
	fmt.Println("3 - Удалить пользователя из файла и структуры")

	control := map[int]func(){
		1: acc.makeAccount,
		2: acc.getUser,
		3: acc.deleteUser,
	}
	var action int

	fmt.Scan(&action)

	if actionFunc, ok := control[action]; ok {
		actionFunc()
	} else {
		fmt.Println()
		fmt.Println("Неверный выбор действия")
		return errors.New("invalid action")
	}

	return nil
}
