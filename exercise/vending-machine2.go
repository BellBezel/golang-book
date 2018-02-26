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
	m.insertedMoney += m.coins[coin]
}

func (m *VendingMachine)  SelectSD() string{
	return "SD"
}

func main() {
	var coins = map[string]int{"T":10, "F":5, "TW":2, "O":1}
	vm := VendingMachine{coins: coins}
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 18
	can := vm.SelectSD()
	fmt.Println(can) //SD
}