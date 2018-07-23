package eth

// Constants to match up protocol versions and messages
const (
	apa01 = 01
)

// ApaProtocolName is the official short name of the protocol used during capability negotiation.
var ApaProtocolName = "apa"

// ApaProtocolVersions are the upported versions of the ethapa protocol (first is primary).
var ApaProtocolVersions = []uint{apa01}

// ApaProtocolLengths are the number of implemented message corresponding to different protocol versions.
var ApaProtocolLengths = []uint64{2}

const ApaProtocolMaxMsgSize = 10 * 1024 * 1024 // Maximum cap on the size of a apaprotocol message

// apa protocol message codes
const (
	// Protocol messages belonging to apa01
	SendTestMsg = 0x11
	RecvTestMsg = 0x12
)

const (
	ErrSendTest = iota
	ErrRecvTest
)

// XXX change once legacy code is out
var apaErrorToString = map[int]string{
	ErrSendTest: "Test error in send",
	ErrRecvTest: "Test error in receive",
}
