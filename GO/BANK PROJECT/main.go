package main

import (
	"fmt"

	"github.com/HayeongShim/learngo/accounts"
)

//-----------------BANK & DICTIONARY PROJECTS-----------------//

// 소문자 : private, 대문자 : public
// constructor가 없기 때문에, function으로 construct하거나 struct를 만들도록 해야한다.

func main() {
	account := accounts.NewAccount("harry")
	account.Deposit(1000)
	fmt.Println(account.GetBalance())

	err := account.Withdraw(5000) // 손수 에러를 체크해주어야 함. 장점이기도..?
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.GetBalance())

	account.ChangeOwner("ron")
	fmt.Println(account.GetOwner())

	fmt.Println(account.String())
}
