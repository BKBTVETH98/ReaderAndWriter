package account

import (
	"errors"
	"fmt"
)

type Account struct { //структура аккаунта
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

func (acc *Account) getUser() { //метод вывода структуры используя указатель на структуру

	fmt.Println(NewVault().Accounts)
}

func (acc *Account) makeAccount() { //метод создания аккаунта используя указатель на структуру

	fmt.Println()
	fmt.Print("Введи желаемый логин - ")
	fmt.Scan(&acc.Login)
	fmt.Print("Введи password - ")
	fmt.Scan(&acc.Pass)
	fmt.Println("Успешное  создание аккаунта")
	fmt.Println()

	vault := NewVault()
	vault.AddAccount(*acc)

	_, err := vault.ToBytes()

	if err != nil {
		fmt.Println("Не удалось преобразовать json")
		return
	}

}

func (acc *Account) deleteUser() {
	fmt.Print("Введите логин пользователя для удаления: ")
	var logForDelete string
	fmt.Scan(&logForDelete)
	NewVault().DeleteAccount(logForDelete)

	fmt.Println("Пользователь удален из структуры и файла")
	acc.getUser() //проверяем что структура очищена
}

func (acc *Account) ControlActions() error {

	fmt.Println()
	fmt.Println("Выберите действие: ")
	fmt.Println("1 - создать аккаунт")
	fmt.Println("2 - вывести структуру и прочитать из файла")
	fmt.Println("3 - Удалить пользователя из файла и структуры")

	control := map[int]any{
		1: acc.makeAccount,
		2: acc.getUser,
		3: acc.deleteUser,
	}

	var action int

	fmt.Scan(&action)

	if actionFunc, ok := control[action]; ok {
		actionFunc.(func())()
	} else {
		fmt.Println()
		fmt.Println("Неверный выбор действия")
		return errors.New("invalid action")
	}

	return nil
}
