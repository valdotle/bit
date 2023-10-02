// package field (bit/field) exposes functions to use with bitfields of type integer as well as string and big.Int aliases to offer methods on bitfields of these types
package field

import (
	"math/big"

	stringfield "github.com/valdotle/bit/field/string"
)

type String string

func (f String) Has(flag int) bool {
	return stringfield.Has(string(f), flag)
}

func (f String) Add(flag int) String {
	return String(stringfield.Add(string(f), flag))
}

func (f String) Flip(flag int) String {
	return String(stringfield.Flip(string(f), flag))
}

func (f String) Remove(flag int) String {
	return String(stringfield.Remove(string(f), flag))
}

func (f String) HasString(flag string) bool {
	return stringfield.HasString(string(f), flag)
}

func (f String) AddString(flag string) String {
	return String(stringfield.AddString(string(f), flag))
}

func (f String) FlipString(flag string) String {
	return String(stringfield.FlipString(string(f), flag))
}

func (f String) RemoveString(flag string) String {
	return String(stringfield.RemoveString(string(f), flag))
}

func (f String) HasBig(flag *big.Int) bool {
	return stringfield.HasBig(string(f), flag)
}

func (f String) AddBig(flag *big.Int) String {
	return String(stringfield.AddBig(string(f), flag))
}

func (f String) FlipBig(flag *big.Int) String {
	return String(stringfield.FlipBig(string(f), flag))
}

func (f String) RemoveBig(flag *big.Int) String {
	return String(stringfield.RemoveBig(string(f), flag))
}
