package object

import (
	"strconv"
)

type Offset struct {
	Value int
}

func NewOffset(q string, limit Limit) (Offset, error) {
	if q == "" {
		return Offset{Value: 0}, nil
	}

	r, err := strconv.Atoi(q)

	// SQLのOffsetは0からなので1を引いておく
	return Offset{Value: (r - 1) * limit.Value}, err
}
