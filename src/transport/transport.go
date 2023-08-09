// (C) 2016-2023 Ant Group Co.,Ltd.
// SPDX-License-Identifier: Apache-2.0

package transport

import (
	"bkr-go/transport/http"
	"bkr-go/transport/info"
	"bkr-go/transport/message"

	"go.uber.org/zap"
)

type Transport interface {
	Broadcast(msg *message.ConsMessage)
}

// InitTransport executes transport layer initiliazation, which returns transport, a channel
// for received ConsMessage, a channel for received requests, and a channel for reply
func InitTransport(lg *zap.Logger, id info.IDType, port int, peers []string) (Transport,
	chan *message.ConsMessage, chan []byte, chan []byte) {
	return http.InitTransport(lg, id, port, peers)
}
