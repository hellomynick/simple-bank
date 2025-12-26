package domain

type AccountStatus string

const (
	AccountStatusActive AccountStatus = "Active"
	AccountStatusLocked AccountStatus = "Locked"
	AccountStatusClosed AccountStatus = "Closed"
)

func (s AccountStatus) IsValid() bool {
	switch s {
	case AccountStatusActive, AccountStatusLocked, AccountStatusClosed:
		return true
	}
	return false
}
