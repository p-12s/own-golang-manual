package onederx

import "github.com/shopspring/decimal"

// GetWsL2SubscribeRequest - формирование JSON-стурктуры для подписки на канал (https://docs.onederx.com/#l2-channel)
/*
user -> exchange
{
  "type": "subscribe",
  "payload": {
    "subscriptions": [
      {
        "params": {
          "symbol": "BTCUSD_P"
        },
        "channel": "ticker"
      }
    ]
  }
}
*/
func GetWsL2SubscribeRequest() interface{} {
	subscribeReq := struct {
		Type    string
		Payload struct {
			Subscriptions []struct {
				Channel string
				Params  struct {
					Symbol string
				}
			}
		}
	}{}
	subscribeReq.Type = "subscribe"
	subscribeReq.Payload.Subscriptions = make([]struct {
		Channel string
		Params  struct {
			Symbol string
		}
	}, 1)
	subscribeReq.Payload.Subscriptions[0].Channel = "l2"
	subscribeReq.Payload.Subscriptions[0].Params.Symbol = "BTCUSD_P" // нам нужен тольео этот инструмент, поэтому он захардкожен

	// возвращаем анонимную структуру, глобальной не делаем - она больше нигде не нужна
	return subscribeReq
}

type WsL2Item struct {
	Price     decimal.Decimal `json:"price"`
	Volume    uint64          `json:"volume,string"`
	Side      string          `json:"side"`
	Timestamp int64           `json:"timestamp"`
}

type WsL2Update struct {
	Params struct {
		Symbol string `json:"symbol"`
	} `json:"params"`
	Payload WsL2Item `json:"payload"`
}

type WsL2Snapshot struct {
	Params struct {
		Symbol string `json:"symbol"`
	} `json:"params"`
	Payload struct {
		Snapshot []*WsL2Item `json:"snapshot"`
		Updates  []*WsL2Item `json:"updates"`
	} `json:"payload"`
}
