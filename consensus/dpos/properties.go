package dpos

type globalStaticProperty struct {
	Account
	NumberOfWitnesses  int
	WitnessList        []string
	PendingWitnessList []string
}
