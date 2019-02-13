package orderbook

import "errors"

var (
	ErrInvalidQuantity = errors.New("orderbook: invalid order quantity")
	ErrInvalidPrice    = errors.New("orderbook: invalid order price")
	ErrOrderExists     = errors.New("orderbook: order already exists")
	ErrOrderNotExists  = errors.New("orderbook: order does not exist")
)
