package domain

import (
	"fmt"
	es "simplebank/internal/hephaistos/event_sourcing"

	"github.com/google/uuid"
)

type Account struct {
	*es.Aggregate

	id uuid.UUID

	// Money represents a monetary amount in the smallest currency unit (e.g., cents, satoshi).
	// Int64 is used to store data efficiently and avoid floating-point precision issues.
	//
	// Example:
	//   - USD: 1050 represents $10.50
	//   - VND: 10000 represents 10,000 VND
	balance int64
}

func (a *Account) Deposit(money int64) {

	event := MoneyDeposited{
		AccountId: a.id,
		Amount:    money,
		Balance:   a.balance,
	}

	a.TrackChanges(event)
	a.Apply(event)
}

func (a *Account) WithDraw(money int64) error {
	if a.balance < money {
		return fmt.Errorf("balance less than amount")
	}

	event := MoneyWithdrawn{
		AccountId: a.id,
		Amount:    money,
		Balance:   a.balance,
	}

	a.TrackChanges(event)
	a.Apply(event)

	return nil
}

func (a *Account) Apply(event es.Event) {
	switch e := event.(type) {
	case MoneyDeposited:
		a.balance += e.Amount
	case MoneyWithdrawn:
		a.balance -= e.Amount
	default:
		fmt.Printf("Don't know what event %s!\n", event.TypeName())
	}
}

type MoneyDeposited struct {
	AccountId uuid.UUID
	Amount    int64
	Balance   int64
}

func (e MoneyDeposited) TypeName() string {
	return "MoneyDeposited"
}

type MoneyWithdrawn struct {
	AccountId uuid.UUID
	Amount    int64
	Balance   int64
}

func (m MoneyWithdrawn) TypeName() string {
	return "MoneyWithdrawn"
}
