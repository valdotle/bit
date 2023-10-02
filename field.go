// package bit provides an integer alias to provide methods for bitfields of type integer
package bit

import "github.com/valdotle/bit/field"

// integer alias used to represent bitfields within integer limitations
type Field int

// check if the bitfield includes the provided flag
func (f Field) Has(flag int) bool {
	return field.Has(int(f), flag)
}

// add the provided flag to the bitfield (set bit to 1 / true)
func (f Field) Add(flag int) Field {
	return Field(field.Add(int(f), flag))
}

// flip the flag's corresponding bit in the bitfield
func (f Field) Flip(flag int) Field {
	return Field(field.Flip(int(f), flag))
}

// remove the provided flag from the bitfield (set bit to 0 / false)
func (f Field) Remove(flag int) Field {
	return Field(field.Remove(int(f), flag))
}
