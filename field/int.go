package field

func Has(field, flag int) bool {
	return field&flag == flag
}

func Add(field, flag int) int {
	return field | flag
}

func Flip(field, flag int) int {
	return field ^ flag
}

func Remove(field, flag int) int {
	return field &^ flag
}
