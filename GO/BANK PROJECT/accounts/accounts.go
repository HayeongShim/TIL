package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw! You are poor")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // 복사본을 리턴하는 것이 아니라, 우리가 만든 것을 리턴!
}

// method
// (account Account) <-- 이 부분을 receiver라고 부른다.
// receiver 작성시, struct의 첫 글자를 따서 소문자로 지어야 한다.
// func (a Account) 라고 하는 순간, '복사본'이 생기므로 func (a *Account)로 해야함. "pointer receiver"
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Get balance of your account
func (a Account) GetBalance() int {
	return a.balance
}

// Withdraw from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Change the owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Get owner of the account
func (a Account) GetOwner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.GetOwner(), "'s account.\nhas: ", a.GetBalance())
}
