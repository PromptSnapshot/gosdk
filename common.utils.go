package utils

import (
	"time"

	"github.com/google/uuid"
)

func BoolAdr(b bool) *bool {
	return &b
}

func UUIDAdr(in uuid.UUID) *uuid.UUID {
	return &in
}

func GetCurrentTimePtr() *time.Time {
	tmp := time.Now()
	return &tmp
}

func GetCurrentYear2Digit() int {
	return (time.Now().Year() + 543) % 100
}
