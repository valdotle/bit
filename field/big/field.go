// package big (bit/field/big) exposes functions to use with bitfields of type big.Int
//
// Intended for use with bitfields that exceed integer limitations.
package big

import "math/big"

func Has(field *big.Int, flag int) bool {
	flagInt := big.NewInt(int64(flag))

	return field.And(field, flagInt).Cmp(flagInt) == 0
}

func Add(field *big.Int, flag int) *big.Int {
	return field.Or(field, big.NewInt(int64(flag)))
}

func Flip(field *big.Int, flag int) *big.Int {
	return field.Xor(field, big.NewInt(int64(flag)))
}

func Remove(field *big.Int, flag int) *big.Int {
	return field.AndNot(field, big.NewInt(int64(flag)))
}

func HasString(field *big.Int, flag string) bool {
	flagInt, success := (&big.Int{}).SetString(flag, 0)
	if !success {
		return false
	}

	return field.And(field, flagInt).Cmp(flagInt) == 0
}

func AddString(field *big.Int, flag string) *big.Int {
	flagInt, success := (&big.Int{}).SetString(flag, 0)
	if !success {
		return nil
	}

	return field.Or(field, flagInt)
}

func FlipString(field *big.Int, flag string) *big.Int {
	flagInt, success := (&big.Int{}).SetString(flag, 0)
	if !success {
		return nil
	}

	return field.Xor(field, flagInt)
}

func RemoveString(field *big.Int, flag string) *big.Int {
	flagInt, success := (&big.Int{}).SetString(flag, 0)
	if !success {
		return nil
	}

	return field.AndNot(field, flagInt)
}

func HasBig(field, flag *big.Int) bool {
	return field.And(field, flag).Cmp(flag) == 0
}

func AddBig(field, flag *big.Int) *big.Int {
	return field.Or(field, flag)
}

func FlipBig(field, flag *big.Int) *big.Int {
	return field.Xor(field, flag)
}

func RemoveBig(field, flag *big.Int) *big.Int {
	return field.AndNot(field, flag)
}
