package dpos

import "errors"

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

func (prop *globalStaticProperty) addPendingWitness(id string) error {
	for _, wit := range prop.WitnessList {
		if id == wit {
			return errors.New("already in witness list")
		}
	}
	for _, wit := range prop.PendingWitnessList {
		if id == wit {
			return errors.New("already in pending list")
		}
	}
	prop.PendingWitnessList = append(prop.PendingWitnessList, id)
	return nil
}

func (prop *globalStaticProperty) deletePendingWitness(id string) error {
	i := 0
	for _, wit := range prop.PendingWitnessList {
		if id == wit {
			newList := append(prop.PendingWitnessList[:i], prop.PendingWitnessList[i+1:]...)
			prop.PendingWitnessList = newList
			return nil
		}
		i++
	}
	return errors.New("witness not in pending list")
}


