package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	balance float64
}

func NewAccount(initialBalance float64) *BankAccount {
	return &BankAccount{balance: initialBalance}
}

func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount

}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount > b.balance {
		return errors.New("недостаточно средств")
	}
	b.balance -= amount
	return nil
}

func main() {
	acc := NewAccount(100)
	fmt.Println("Изначальный баланс:", acc.GetBalance())

	acc.Deposit(10)
	acc.Withdraw(20)

	fmt.Println("Внесли 10 и сняли 20, итоговый баланс :", acc.GetBalance())

}
