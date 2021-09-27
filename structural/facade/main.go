package main

import (
	"fmt"
	"log"
)

// Facade is a structural design pattern that
// provides a simplified interface to a library, a framework,
// or any other complex set of classes.
// 将复杂的内部子系统，封装出简单的接口给client调用
// 例子：电话下单，客服就是你的facade
// 例子：信用卡支付：
// Check account
//Check security PIN
//Credit/debit balance
//Make ledger entry
//Send notification
func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
