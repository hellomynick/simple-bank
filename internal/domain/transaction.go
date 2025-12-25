package domain

import "time"

type Transaction struct {
	Money    float32
	CreateAt time.Time
}
