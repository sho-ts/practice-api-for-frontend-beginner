package object

import (
	"errors"
	"strconv"
)

type Limit struct {
	Value int
}

func NewLimit(q string) (Limit, error) {
	if q == "" {
		return Limit{Value: 15}, nil
	}

	r, err := strconv.Atoi(q)

	if r > 50 {
		return Limit{}, errors.New("limit must be less than 50")
	}

	return Limit{Value: r}, err
}
