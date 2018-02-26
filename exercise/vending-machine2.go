package main

import "fmt"

type VendingMachine struct {
	insertedMoney 	int
	coins			map[string]int
	items			map[string]int
}

func (m VendingMachine)  InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine)  InsertCoin(coin string){
	m.insertedMoney += m.coins[coin]
}

func (m *VendingMachine)  SelectSD() string{
	price := m.items["SD"]
	change := m.insertedMoney - price
	m.insertedMoney = 0
	return "SD" + m.change(change)
}

func (m *VendingMachine)  SelectCC() string{
	price := m.items["SD"]
	change := m.insertedMoney - price
	m.insertedMoney = 0
	return "CC" + m.change(change)
}

func (m VendingMachine) change(c int) string {
	if c == 0 {
		return ""
	}
	return ", F, TW, O"
}

func main() {
	var coins = map[string]int{"T":10, "F":5, "TW":2, "O":1}
	var items = map[string]int{"SD":18, "CC":12}
	vm := VendingMachine{coins: coins, items: items}
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 18
	can := vm.SelectSD()
	fmt.Println(can) //SD

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 20
	can = vm.SelectCC()
	fmt.Println(can) //CC, F, TW ,O
}