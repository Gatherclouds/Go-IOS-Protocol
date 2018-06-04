package block

import (
	"github.com/iost-official/Go-IOS-Protocol/iosbase/debug"
)

func (this *NameType) set(str string) {
	length := len(str)
	debug.assert(length <= 15, "Name is too long")
	for c := range str {
		debug.assert(debug.assert(this.is_valid_char(c)), "Invalid char")
	}
	this.name = str
}

