// Account API
package account

import "sync"

type Account struct {
	balance int64
	isOpen  bool
	sync.Mutex
}

// Take an opening deposit and return an Account.
func Open(d int64) *Account {
	if d < 0 {
		return nil
	}
	return &Account{
		balance: d,
		isOpen:  true,
		// zero value of Mutex is an unlocked Mutex.
		// sync.Mutex{},
	}
}

// Account.Close() closes an account. Set isOpen: false, balance: 0, and return payout. Fail if the account is already closed.
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.isOpen {
		p := a.balance
		a.balance = 0
		a.isOpen = false
		return p, true
	}
	return 0, false
}

// Account.Balance() returns the account balance.
func (a *Account) Balance() (b int64, ok bool) {
	// Fails go test -race without this. But does that matter with mere reads?
	a.Lock()
	defer a.Unlock()
	if a.isOpen {
		return a.balance, true
	}
	return 0, false
}

// Account.Deposit() deposits and withdraws amounts. Withdrawals must fail if the balance would become negative.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	switch {
	case a.balance+amount < 0:
		return a.balance, false
	case a.isOpen:
		a.balance += amount
		return a.balance, true
	default:
		return 0, false
	}
}
