package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var mutex sync.Mutex
	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Counter:", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance:", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user 1", user1.name)
	user1.Change(-amount)
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user 2", user2.name)
	user2.Change(amount)
	time.Sleep(1 * time.Second)

	user2.Unlock()
	user1.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := &UserBalance{
		name:    "Marchel",
		Balance: 1000000,
	}
	user2 := &UserBalance{
		name:    "Matthew",
		Balance: 1000000,
	}
	go Transfer(user1, user2, 100000)
	go Transfer(user2, user1, 200000)
	time.Sleep(10 * time.Second)
	fmt.Println("User 1 :", user1.name, ", Balance :", user1.Balance)
	fmt.Println("User 2 :", user2.name, ", Balance :", user2.Balance)
}

// test github actions
