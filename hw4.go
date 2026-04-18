package main

import (
	"fmt"
	"strings"
)

func main() {
	//1
	/*var sender string = "Ali"
	var receiver string = "Vali"
	var amount int = 100000
	commission := amount * 1 / 100
	/total := amount + commission
	//fmt.Println("Чек Перевода:")
	//fmt.Printf("Отправитель: %s\n", sender)
	//fmt.Printf("Получатель: %s\n", receiver)
	//fmt.Printf("Сумма перевода: %d сум\n", amount)
	//fmt.Printf("Комиссия (1%%): %d сум\n", commission)
	//fmt.Printf("Итого: %d сум\n", total)*/

	////2
	/*var name string = "iPhone 15 Pro"
	var brand string = "Apple"
	var price int64 = 12_990_000
	var stock bool = true
	months := 12
	monthly := price / int64(months)
	fmt.Println("===== Alifshop =====")
	fmt.Printf("Товар: %s\n", name)
	fmt.Printf("Бренд: %s\n", brand)
	fmt.Printf("Цена: %d сум\n", price)
	fmt.Printf("В наличии: %t\n", stock)
	fmt.Printf("Рассрочка: %d мес → %d сум/мес\n", months, monthly)*/

	//3
	card := "4111-2222-3333-4444"
	result := MaskCard(card)
	fmt.Println("Карта:", result)

}

// 3
func MaskCard(number string) string {

	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")

	if len(number) != 16 {
		panic("ошибка в размере")
	}

	for _, ch := range number {
		if ch < '0' || ch > '9' {
			panic("не цифры")
		}
	}

	return number[:4] + " **** **** " + number[12:]
}
