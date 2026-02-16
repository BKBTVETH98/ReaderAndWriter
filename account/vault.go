package account

import (
	files "JsonExample/file"
	"encoding/json"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (vault *Vault) ToBytes() ([]byte, error) { //метод преобразования структуры в байты для записи в файл
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewVault() *Vault { // метод создания новой структуры и чтения из файла
	file, err := files.ReadFile("account.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red(err.Error())
	}

	return &vault
}

func (vault *Vault) AddAccount(acc Account) { //Метод добавления аккаунта в структуру и файл
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	files.WriteFile(data) //запись в файл
}

func (v *Vault) DeleteAccount(login string) { //переписать метод удаления аккаунта из структуры и файла
	for i, acc := range v.Accounts {
		if acc.Login == login {
			v.Accounts = append(v.Accounts[:i], v.Accounts[i+1:]...)
			v.UpdatedAt = time.Now()
			data, err := v.ToBytes()
			if err != nil {
				color.Red("Не удалось преобразовать")
			}
			files.WriteFile(data)
			return
		}
	}
	color.Red("Пользователь с логином %s не найден", login)
}
