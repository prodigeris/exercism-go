package account

import "math"

type Account struct {
	active  bool
	balance int64
}

// Define the Account type here.

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, active: true}
}

func (a *Account) Balance() (int64, bool) {
	if !a.active {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if !a.active {
		return 0, false
	}
	if amount >= 0 {
		a.balance += amount
		return a.balance, true
	}

	d := int64(math.Abs(float64(amount)))

	if a.balance < d {
		return 0, false
	}

	a.balance -= d
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	if !a.active {
		return 0, false
	}
	var b int64
	a.active = false
	b, a.balance = a.balance, 0

	return b, true
}
