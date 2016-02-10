package base3

import (
	"errors"
	"fmt"
)

type Base3Uint struct {
	data   []uint
	length uint
}

func (b *Base3Uint) String() string {
	retStr := ""
	for _, digit := range b.data {
		retStr += fmt.Sprintf("%d", digit)
	}
	return retStr
}

func (b *Base3Uint) addDigit(d uint) {
	b.data = append(b.data, d)
	b.length += 1
}

func (b *Base3Uint) Encode(i uint) {
	b.length = 0
	b.data = make([]uint, 0)
	for i > 0 {
		b.addDigit(i % 3)
		i /= 3
	}
}

func (b *Base3Uint) Decode() uint {
	var decoded uint = 0
	var places uint = 1
	for _, digit := range b.data {
		decoded += digit * places
		places *= 3
	}
	return decoded
}

func ParseString(s string) (*Base3Uint, error) {
	var b Base3Uint
	for _, char := range s {
		if !('0' < char && char < '3') {
			return nil, errors.New(fmt.Sprintf("Invalid character at %c", char))
		}
		b.addDigit(uint(char - '0'))
	}
	return &b, nil
}

func New(i uint) *Base3Uint {
	var b Base3Uint
	b.Encode(i)
	return &b
}
