package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
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
	fmt.Println("Lock U1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock U2", user2.Name)
	user2.Change(amount)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Budi",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Refa",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(5 * time.Second)
	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)
	// Deadlocked karena pada Goroutine pertama user 1 sudah dilock sedangkan pada goroutine kedua user1 dilock setelah user2,
	// sehingga goroutine kedua menunggu goroutine pertama melakukan unlock pada user1 yang mana tidak akan pernah terjadi
	// begitu juga pada goroutine pertama dimana menunggu goroutine ke 2 melakukan unlock pada user2
}
