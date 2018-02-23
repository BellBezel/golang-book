package main

import "fmt"

type VendingMachine struct {
	insert string
	Get string
	SelectSD string
}

func main() {
	//vm := NewVendingMachine(coins, items)
	vm := VendingMachine{}
	// Buy SD(soft drink) with exact change
	vm.InsertCoin("T")
	//vm.InsertCoin("F")
	//vm.InsertCoin("TW")
	//vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 18
	//can := vm.SelectSD()
	//fmt.Println(can) // SD
}

func (vm VendingMachine)  InsertCoin(coin string) {

}

func (vm VendingMachine)  GetInsertedMoney(coin string) int {
	return 
}

