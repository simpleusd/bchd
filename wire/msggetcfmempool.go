// Copyright (c) 2018 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"io"
)

// MsgGetCFMempool is a request for a filter of the remote peer's mempool.
type MsgGetCFMempool struct{}

// BchDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgGetCFMempool) BchDecode(r io.Reader, pver uint32, _ MessageEncoding) error {
	return nil
}

// BchEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgGetCFMempool) BchEncode(w io.Writer, pver uint32, _ MessageEncoding) error {
	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgGetCFMempool) Command() string {
	return CmdGetCFMempool
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgGetCFMempool) MaxPayloadLength(pver uint32) uint32 {
	return 0
}

// NewMsgGetCFMempool returns a new bitcoin getcfmempool message that conforms
// to the Message interface using the passed parameters and defaults for the
// remaining fields.
func NewMsgGetCFMempool() *MsgGetCFMempool {
	return &MsgGetCFMempool{}
}
