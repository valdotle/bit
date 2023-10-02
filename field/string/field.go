// package string (bit/field/field) exposes functions to use with bitfields of type string
//
// Intended for use with bitfields that exceed integer limitations. Errors resulting from parsing strings to bigints are silently discarded and an empty string ("") is returned.
package string

import "math/big"

func Has(field string, flag int) bool {
	fieldInt, success := (&big.Int{}).SetString(field, 0)
	if !success {
		return false
	}

	flagBigInt := big.NewInt(int64(flag))

	return fieldInt.And(fieldInt, flagBigInt).Cmp(flagBigInt) == 0
}

func Add(field string, flag int) string {
	fieldInt, success := (&big.Int{}).SetString(field, 0)
	if !success {
		return ""
	}

	return fieldInt.Or(fieldInt, big.NewInt(int64(flag))).String()
}

func Flip(field string, flag int) string {
	fieldInt, success := (&big.Int{}).SetString(field, 0)
	if !success {
		return ""
	}

	return fieldInt.Xor(fieldInt, big.NewInt(int64(flag))).String()
}

func Remove(field string, flag int) string {
	fieldInt, success := (&big.Int{}).SetString(field, 0)
	if !success {
		return ""
	}

	return fieldInt.AndNot(fieldInt, big.NewInt(int64(flag))).String()
}

func HasString(field, flag string) bool {
	fieldInt, success := (&big.Int{}).SetString(field, 0)
	if !success {
		return false
	}

	flagInt, success := (&big.Int{}).SetString(flag, 0)
	if !success {
		return false
	}

	return fieldInt.And(fieldInt, flagInt).Cmp(flagInt) == 0
}

func AddString(field, flag string) string {
	fieldInt, success := (&big.Int{}).SetString(field, 0)
	if !success {
		return ""
	}

	flagInt, success := (&big.Int{}).SetString(flag, 0)
	if !success {
		return ""
	}

	return fieldInt.Or(fieldInt, flagInt).String()
}

func FlipString(field, flag string) string {
	fieldInt, success := (&big.Int{}).SetString(field, 0)
	if !success {
		return ""
	}

	flagInt, success := (&big.Int{}).SetString(flag, 0)
	if !success {
		return ""
	}

	return fieldInt.Xor(fieldInt, flagInt).String()
}

func RemoveString(field, flag string) string {
	fieldInt, success := (&big.Int{}).SetString(field, 0)
	if !success {
		return ""
	}

	flagInt, success := (&big.Int{}).SetString(flag, 0)
	if !success {
		return ""
	}

	return fieldInt.AndNot(fieldInt, flagInt).String()
}

func HasBig(field string, flag *big.Int) bool {
	fieldInt, success := (&big.Int{}).SetString(field, 0)
	if !success {
		return false
	}

	return fieldInt.And(fieldInt, flag).Cmp(flag) == 0
}

func AddBig(field string, flag *big.Int) string {
	fieldInt, success := (&big.Int{}).SetString(field, 0)
	if !success {
		return ""
	}

	return fieldInt.Or(fieldInt, flag).String()
}

func FlipBig(field string, flag *big.Int) string {
	fieldInt, success := (&big.Int{}).SetString(field, 0)
	if !success {
		return ""
	}

	return fieldInt.Xor(fieldInt, flag).String()
}

func RemoveBig(field string, flag *big.Int) string {
	fieldInt, success := (&big.Int{}).SetString(field, 0)
	if !success {
		return ""
	}

	return fieldInt.AndNot(fieldInt, flag).String()
}
