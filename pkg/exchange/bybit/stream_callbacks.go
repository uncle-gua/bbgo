// Code generated by "callbackgen -type Stream"; DO NOT EDIT.

package bybit

import (
	"github.com/c9s/bbgo/pkg/exchange/bybit/bybitapi"
)

func (s *Stream) OnBookEvent(cb func(e BookEvent)) {
	s.bookEventCallbacks = append(s.bookEventCallbacks, cb)
}

func (s *Stream) EmitBookEvent(e BookEvent) {
	for _, cb := range s.bookEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnWalletEvent(cb func(e []bybitapi.WalletBalances)) {
	s.walletEventCallbacks = append(s.walletEventCallbacks, cb)
}

func (s *Stream) EmitWalletEvent(e []bybitapi.WalletBalances) {
	for _, cb := range s.walletEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnKLineEvent(cb func(e KLineEvent)) {
	s.kLineEventCallbacks = append(s.kLineEventCallbacks, cb)
}

func (s *Stream) EmitKLineEvent(e KLineEvent) {
	for _, cb := range s.kLineEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnOrderEvent(cb func(e []OrderEvent)) {
	s.orderEventCallbacks = append(s.orderEventCallbacks, cb)
}

func (s *Stream) EmitOrderEvent(e []OrderEvent) {
	for _, cb := range s.orderEventCallbacks {
		cb(e)
	}
}

func (s *Stream) OnTradeEvent(cb func(e []TradeEvent)) {
	s.tradeEventCallbacks = append(s.tradeEventCallbacks, cb)
}

func (s *Stream) EmitTradeEvent(e []TradeEvent) {
	for _, cb := range s.tradeEventCallbacks {
		cb(e)
	}
}
