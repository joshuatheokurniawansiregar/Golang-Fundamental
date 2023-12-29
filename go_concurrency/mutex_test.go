package go_concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Looping(counter *int) {
	var mut sync.Mutex
	for i := 1; i <= 100; i++ {
		mut.Lock()
		*counter = *counter + 1
		mut.Unlock()
	}
}

// Mutext menangaani race condition
func TestMutext(t *testing.T) {
	counter := 0
	var mut sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for i := 1; i <= 100; i++ {
				mut.Lock()
				counter = counter + 1
				mut.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(counter)
}

type BankAccount struct {
	RWMut   sync.RWMutex
	Balance int
}

func (bankAccount *BankAccount) AddBalance(balance int) {
	bankAccount.RWMut.Lock()
	bankAccount.Balance += 1
	bankAccount.RWMut.Unlock()
}

func (bankAccount *BankAccount) GetBalance() int {
	bankAccount.RWMut.RLock()
	var bank_account int = bankAccount.Balance
	bankAccount.RWMut.RUnlock()
	return bank_account
}

func TestRWMutext(t *testing.T) {
	var account BankAccount
	for i := 1; i <= 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance: ", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	sync.Cond
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) DecreasesUserBalance(amount int) {
	user.Balance = user.Balance - amount
}

func (user *UserBalance) IncerasesUserBalance(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(sender *UserBalance, receiver *UserBalance, amount int) {
	sender.Lock()
	fmt.Println("Lock user: ", sender.Name)
	sender.DecreasesUserBalance(amount)

	receiver.Lock()
	fmt.Println("Lock user: ", receiver.Name)
	receiver.IncerasesUserBalance(amount)
	receiver.Unlock()
	receiver.Unlock()
}

// sync.Cond untuk mengatasi deadlock karena locker
var sourceUserBalanceGlobal *UserBalance = &UserBalance{}

// var destinationUserBalance *UserBalance = &UserBalance{}
var cond *sync.Cond = sync.NewCond(&sourceUserBalanceGlobal.Mutex)

// var condReciever *sync.Cond = sync.NewCond(&sourceUserBalance.Mutex)
var group *sync.WaitGroup = &sync.WaitGroup{}

func (userBalance *UserBalance) ChangeUserBalance(money int) {
	userBalance.Balance = userBalance.Balance + money
}

func TransferWithSyncCondition(sourceUserBalance *UserBalance, destinationUserBalance *UserBalance, value int) {
	beforeSourceBalanceTransferBalance := sourceUserBalance.Balance
	// beforeDestinationBalanceTransferBalance := sourceUserBalance.Balance
	cond.L.Lock()
	sourceUserBalance.ChangeUserBalance(-value)
	var sourceBalance int = sourceUserBalance.Balance

	if sourceBalance == beforeSourceBalanceTransferBalance {
		cond.Wait()
	}
	destinationUserBalance.ChangeUserBalance(value)
	// var destinationBalance int = destinationUserBalance.Balance
	// if destinationBalance == beforeDestinationBalanceTransferBalance {
	// 	cond.Wait()
	// }
	cond.L.Unlock()

	group.Done()
}
func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Joshua Theo Kurniawan",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Budi",
		Balance: 500000,
	}
	go Transfer(&user1, &user2, 200000)
	go Transfer(&user2, &user1, 100000)
	time.Sleep(5 * time.Second)
	fmt.Println("User: ", user1.Name, "Balance: ", user1.Balance)
	fmt.Println("User: ", user2.Name, "Balance: ", user2.Balance)

}

func TestSyncCond(t *testing.T) {
	user1 := UserBalance{
		Name:    "Joshua Theo Kurniawan",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Budi",
		Balance: 500000,
	}
	group.Add(3)
	go TransferWithSyncCondition(&user1, &user2, 200000)
	go TransferWithSyncCondition(&user2, &user1, 100000)
	go TransferWithSyncCondition(&user2, &user1, 100000)
	group.Wait()
	fmt.Println("User: ", user1.Name, "Balance: ", user1.Balance)
	fmt.Println("User: ", user2.Name, "Balance: ", user2.Balance)

}
