package main

import "fmt"

func main() {
	var sender string = "Ali"
	var receiver string = "Vali"
	var amount int = 100000
	commission := amount * 1 / 100
	total := amount + commission
	fmt.Println("Чек Перевода:")
	fmt.Printf("Отправитель: %s\n", sender)
	fmt.Printf("Получатель: %s\n", receiver)
	fmt.Printf("Сумма перевода: %d сум\n", amount)
	fmt.Printf("Комиссия (1%%): %d сум\n", commission)
	fmt.Printf("Итого: %d сум\n", total)
	
}
