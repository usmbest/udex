package types

import (
	"fmt"
)

// SubscriptionEvent is an enum signifies whether the incoming message is of type Subscribe or unsubscribe
type SubscriptionEvent string

// Enum members for SubscriptionEvent
const (
	SUBSCRIBE   SubscriptionEvent = "SUBSCRIBE"
	UNSUBSCRIBE SubscriptionEvent = "UNSUBSCRIBE"
	Fetch       SubscriptionEvent = "fetch"
)

const TradeChannel = "trades"
const OrderbookChannel = "order_book"
const OrderChannel = "orders"
const OHLCVChannel = "ohlcv"

type WebsocketMessage struct {
	Channel string         `json:"channel"`
	Event   WebsocketEvent `json:"event"`
}

func (ev *WebsocketMessage) String() string {
	return fmt.Sprintf("%v/%v", ev.Channel, ev.Event.String())
}

type WebsocketEvent struct {
	Type    string      `json:"type"`
	Hash    string      `json:"hash,omitempty"`
	Payload interface{} `json:"payload"`
}

func (ev *WebsocketEvent) String() string {
	return fmt.Sprintf("%v", ev.Type)
}

// Params is a sub document used to pass parameters in Subscription messages
type Params struct {
	From     int64  `json:"from"`
	To       int64  `json:"to"`
	Duration int64  `json:"duration"`
	Units    string `json:"units"`
	PairID   string `json:"pair"`
}

type OrderPendingPayload struct {
	Matches *Matches `json:"matches"`
}

type OrderSuccessPayload struct {
	Matches *Matches `json:"matches"`
}

type OrderMatchedPayload struct {
	Matches *Matches `json:"matches"`
}

type SubscriptionPayload struct {
	PairName   string `json:"pairName,omitempty"`
	QuoteToken string `json:"quoteToken,omitempty"`
	BaseToken  string `json:"baseToken,omitempty"`
	From       int64  `json:"from"`
	To         int64  `json:"to"`
	Duration   int64  `json:"duration"`
	Units      string `json:"units"`
}

func NewOrderWebsocketMessage(o *Order) *WebsocketMessage {
	return &WebsocketMessage{
		Channel: "orders",
		Event: WebsocketEvent{
			Type:    "NEW_ORDER",
			Hash:    o.Hash,
			Payload: o,
		},
	}
}

func NewOrderAddedWebsocketMessage(o *Order, p *Pair, filled int64) *WebsocketMessage {
	o.Process(p)
	o.FilledAmount = filled
	o.Status = "OPEN"
	return &WebsocketMessage{
		Channel: "orders",
		Event: WebsocketEvent{
			Type:    "ORDER_ADDED",
			Hash:    o.Hash,
			Payload: o,
		},
	}
}

func NewOrderCancelWebsocketMessage(oc *OrderCancel) *WebsocketMessage {
	return &WebsocketMessage{
		Channel: "orders",
		Event: WebsocketEvent{
			Type:    "CANCEL_ORDER",
			Hash:    oc.OrderHash,
			Payload: oc,
		},
	}
}
