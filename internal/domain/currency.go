package domain

type Currency string

const (
	USD Currency = "USD"
	VND Currency = "VND"
)

func (c Currency) IsValid() bool {
	switch c {
	case USD, VND:
		return true
	}
	return false
}
