package main

import "fmt"

type VendingMachine struct {
	insertedMoney 	int
	coins			map[string]int
}

func (m VendingMachine)  InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine)  InsertCoin(coin string){
	/*if coin == "T" {
		m.insertedMoney += 10
	}
	if coin == "F" {
		m.insertedMoney += 5
	}
	if coin == "TW" {
		m.insertedMoney += 2
	}
	if coin == "O" {
		m.insertedMoney += 1
	}*/
	m.insertedMoney += m.coins[coin]
}

func main() {
	var coins = map[string]int{"T":10, "F":5, "TW":2, "O":1}
	vm := VendingMachine{coins: coins}
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 0
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 10
}