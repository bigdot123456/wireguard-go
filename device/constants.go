/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2019 WireGuard LLC. All Rights Reserved.
 */

package device

import (
	"time"
)

/* Specification constants */

const (
	RekeyAfterMessages      = (1 << 64) - (1 << 16) - 1
	RejectAfterMessages     = (1 << 64) - (1 << 4) - 1
	RekeyAfterTime          = time.Second * 120
	RekeyAttemptTime        = time.Second * 90
	RekeyTimeout            = time.Second * 5
	MaxTimerHandshakes      = 90 / 5 /* RekeyAttemptTime / RekeyTimeout */
	RekeyTimeoutJitterMaxMs = 334
	RejectAfterTime         = time.Second * 180
	KeepaliveTimeout        = time.Second * 10
	CookieRefreshTime       = time.Second * 120
	HandshakeInitationRate  = time.Second / 50
	PaddingMultiple         = 16
)

const (
	MinMessageSize = MessageKeepaliveSize                  // minimum size of transport message (keepalive)
	MaxMessageSize = MaxSegmentSize                        // maximum size of transport message
	MaxContentSize = MaxSegmentSize - MessageTransportSize // maximum size of transport message content
)

/* Implementation constants */

const (
	UnderLoadQueueSize = QueueHandshakeSize / 8
	UnderLoadAfterTime = time.Second // how long does the device remain under load after detected
	MaxPeers           = 1 << 16     // maximum number of configured peers
)
