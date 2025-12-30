package domain

import (
	"fmt"
	"simplebank/internal/common"
	pb "simplebank/internal/domain/events"
	es "simplebank/pkg/hephaistos/core"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
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
	balance       int64
	currency      Currency
	accountStatus AccountStatus
}

func (a *Account) ID() uuid.UUID {
	return a.id
}

func (a *Account) Status() string {
	return string(a.accountStatus)
}

func NewAccount(balance int64, currencyStr string) (*Account, error) {
	if balance < 0 {
		return nil, fmt.Errorf("initial balance cannot be negative")
	}

	account := &Account{
		Aggregate: &es.Aggregate{},
	}

	currency := Currency(currencyStr)
	if !currency.IsValid() {
		return nil, fmt.Errorf("invalid currency")
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	event := &pb.AccountCreated{
		AccountId: id.String(),
		Balance:   balance,
		Currency:  string(currency),
		Status:    string(AccountStatusActive),
	}

	if err := account.Commit(event); err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) Deposit(amount int64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	event := &pb.MoneyDeposited{
		Amount:    amount,
		AccountId: a.id.String(),
	}

	err := a.Commit(event)
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) Withdraw(amount int64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}

	if a.balance < amount {
		return fmt.Errorf("insufficient balance")
	}

	event := &pb.MoneyWithdrawn{
		Amount:    amount,
		AccountId: a.id.String(),
	}

	err := a.Commit(event)
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) Commit(event proto.Message) error {
	a.Apply(event)

	err := a.TrackChange(a.id.String(), event)
	if err != nil {
		return err
	}

	return nil
}

func (a *Account) Apply(event proto.Message) {
	switch e := event.(type) {
	case *pb.AccountCreated:
		a.id = uuid.MustParse(e.AccountId)
		a.balance = e.Balance
		a.currency = Currency(e.Currency)
		a.accountStatus = AccountStatus(e.Status)
	case *pb.MoneyDeposited:
		a.balance += e.Amount
	case *pb.MoneyWithdrawn:
		a.balance -= e.Amount
	default:
		fmt.Printf("Don't know what event %s!\n", common.GetEventName(event))
	}
}
