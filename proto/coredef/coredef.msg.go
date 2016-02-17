// This code is generated by github.com/davyxu/cellnet/protoc-gen-msg, DO NOT EDIT
// Generated from: coredef.proto
package coredef

import (
	"github.com/davyxu/cellnet"
)

func init() {
	cellnet.RegisterMessage("coredef.SessionAccepted", (*SessionAccepted)(nil))
	cellnet.RegisterMessage("coredef.SessionConnected", (*SessionConnected)(nil))
	cellnet.RegisterMessage("coredef.SessionClosed", (*SessionClosed)(nil))
	cellnet.RegisterMessage("coredef.PeerInit", (*PeerInit)(nil))
	cellnet.RegisterMessage("coredef.PeerStart", (*PeerStart)(nil))
	cellnet.RegisterMessage("coredef.PeerStop", (*PeerStop)(nil))
	cellnet.RegisterMessage("coredef.UpstreamACK", (*UpstreamACK)(nil))
	cellnet.RegisterMessage("coredef.CloseClientACK", (*CloseClientACK)(nil))
	cellnet.RegisterMessage("coredef.DownstreamACK", (*DownstreamACK)(nil))
	cellnet.RegisterMessage("coredef.RemoteCallREQ", (*RemoteCallREQ)(nil))
	cellnet.RegisterMessage("coredef.RemoteCallACK", (*RemoteCallACK)(nil))
	cellnet.RegisterMessage("coredef.TestEchoACK", (*TestEchoACK)(nil))
}
