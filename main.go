package main

import (
	account "exercism/bank-account"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	a := account.Open(0)
	const amt = 10
	const c = 1000

	var negBal int32
	var start, g sync.WaitGroup

	start.Add(1)
	g.Add(3 * c)

	for i := 0; i < c; i++ {
		go func() { // deposit
			start.Wait()
			a.Deposit(amt) // ignore return values
			g.Done()
		}()
		go func() { // withdraw
			start.Wait()
			for {
				if _, ok := a.Deposit(-amt); ok {
					break
				}
				time.Sleep(time.Microsecond) // retry
			}
			g.Done()
		}()
		go func() { // watch that balance stays >= 0
			start.Wait()
			if p, _ := a.Balance(); p < 0 {
				atomic.StoreInt32(&negBal, 1)
			}
			g.Done()
		}()
	}
	start.Done()
	g.Wait()

	if negBal == 1 {
		fmt.Println("Balance went negative with concurrent deposits and " +
			"withdrawals.  Want balance always >= 0.")
	}
	if p, ok := a.Balance(); !ok || p != 0 {
		fmt.Printf("After equal concurrent deposits and withdrawals, "+
			"a.Balance = %d, %t.  Want 0, true", p, ok)
	}

	fmt.Println(a.Balance())
}
