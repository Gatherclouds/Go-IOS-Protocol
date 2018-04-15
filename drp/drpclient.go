package drp

type Client interface {
	RequestRan(seed []byte) []byte
	AddServer(server Server)
	OnReceivingProof()
	OnReceivingSignOff()
}
