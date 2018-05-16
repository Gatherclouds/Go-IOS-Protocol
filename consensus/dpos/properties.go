package dpos

type globalStaticProperty struct {
	Account
	NumberOfWitnesses  int
	WitnessList        []string
	PendingWitnessList []string
}

func newGlobalStaticProperty(acc Account, witnessList []string) globalStaticProperty {
	prop := globalStaticProperty{
		Account:            acc,
		NumberOfWitnesses:  len(witnessList),
		WitnessList:        witnessList,
		PendingWitnessList: []string{},
	}
	return prop
}

