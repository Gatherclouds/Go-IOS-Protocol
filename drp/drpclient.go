package drp

type UserClient interface {
	RequestRandom(seed []byte) []byte
	AddServer(server Server)
	OnReceivingProof()
	OnReceivingSignOff()
}
