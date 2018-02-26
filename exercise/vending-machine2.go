package main

import "fmt"

type VendingMachine struct {
	insertedMoney int
}

func (m VendingMachine)  InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine)  InsertCoin(coin string){
	if coin == "T" {
		m.insertedMoney += 10
	}
	if coin == "F" {
		m.insertedMoney += 5
	}
}

func main() {
	vm := VendingMachine{}
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 0
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 10
}