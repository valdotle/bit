package field

import (
	"math/big"

	bigfield "github.com/valdotle/bit/field/big"
)

type Big big.Int

func (f *Big) Has(flag int) bool {
	return bigfield.Has((*big.Int)(f), flag)
}

func (f *Big) Add(flag int) *Big {
	return (*Big)(bigfield.Add((*big.Int)(f), flag))
}

func (f *Big) Flip(flag int) *Big {
	return (*Big)(bigfield.Flip((*big.Int)(f), flag))
}

func (f *Big) Remove(flag int) *Big {
	return (*Big)(bigfield.Remove((*big.Int)(f), flag))
}

func (f *Big) HasString(flag string) bool {
	return bigfield.HasString((*big.Int)(f), flag)
}

func (f *Big) AddString(flag string) *Big {
	return (*Big)(bigfield.AddString((*big.Int)(f), flag))
}

func (f *Big) FlipString(flag string) *Big {
	return (*Big)(bigfield.FlipString((*big.Int)(f), flag))
}

func (f *Big) RemoveString(flag string) *Big {
	return (*Big)(bigfield.RemoveString((*big.Int)(f), flag))
}

func (f *Big) HasBig(flag *big.Int) bool {
	return bigfield.HasBig((*big.Int)(f), flag)
}

func (f *Big) AddBig(flag *big.Int) *Big {
	return (*Big)(bigfield.AddBig((*big.Int)(f), flag))
}

func (f *Big) FlipBig(flag *big.Int) *Big {
	return (*Big)(bigfield.FlipBig((*big.Int)(f), flag))
}

func (f *Big) RemoveBig(flag *big.Int) *Big {
	return (*Big)(bigfield.RemoveBig((*big.Int)(f), flag))
}
