package main

import "fmt"

func main() {

	//var sender string = "Ali"
	//var receiver string = "Vali"
	//var amount int = 100000
	//commission := amount * 1 / 100
	//total := amount + commission
	//fmt.Println("Чек Перевода:")
	//fmt.Printf("Отправитель: %s\n", sender)
	//fmt.Printf("Получатель: %s\n", receiver)
	//fmt.Printf("Сумма перевода: %d сум\n", amount)
	//fmt.Printf("Комиссия (1%%): %d сум\n", commission)
	//fmt.Printf("Итого: %d сум\n", total)

	var name string = "iPhone 15 Pro"
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
	fmt.Printf("Рассрочка: %d мес → %d сум/мес\n", months, monthly)
}
