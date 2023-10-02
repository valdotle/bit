[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/valdotle/bit.svg)](https://github.com/valdotle/bit)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/valdotle/bit)
[![GoReportCard example](https://goreportcard.com/badge/github.com/valdotle/bit)](https://goreportcard.com/report/github.com/valdotle/bit)
[![GitHub license](https://badgen.net/github/license/valdotle/bit)](https://github.com/valdotle/bit/blob/main/LICENSE)
[![Github tag](https://badgen.net/github/tag/valdotle/bit)](https://github.com/valdotle/bit/tags/)

Module **bit** comes with a variety of utility functions for use with **bitfields**.

The module offers functionality for use with go built-in types (`string`, `int` and `big.Int`) as well as aliases for these types to use methods instead.

The module exposes variations of the following functions:

```go
    Has() bool // whether a bitfield includes the flag specified; always only returns a boolean
    Add()      // add a flag to a bitfield (set corresponding bit to 1 / true)
    Flip()     // flip a flag's state in a bitfield (flip corresponding bit)
    Remove()   // remove a flag from a bitfield (set corresponding bit to 0 / false)
```

Note that all functions' (except for `Has()`) return value matches the type of the bitfield. This means, that methods defined on an alias type bitfield return a bitfield of that alias type as well.

The combination of **import path** and function / method **suffix** determines the data type of the bitfield and flag. This is explained in more detail below. **`int` is assumed as the default / most common bitfield type and thus doesn't need to be explicitly specified.** Bitfields of type `string` and `big.Int` are intended to be used where integer limitations are exceeded and **only then**. Since no checks are performed whether a string or bigint passed actually has to be handled as such, all string and bigint related functions **always** use `math/big`'s `big.Int` and thus have a significantly slower performance. 

The **import path** determines the type of the bitfield:

```go
    import (
        "github.com/valdotle/bit/field"                    // bitfield of type int - Note how it's not "bit/field/int"
        stringfield "github.com/valdotle/bit/field/string" // bitfield of type string
        "github.com/valdotle/bit/field/big"                // bitfield of type big.Int
    )

    flag := 2
    field.Has(9, flag)           // false

    stringfield.Has("9", flag)   // false

    big.Has(big.NewInt(9), flag) // false
```

**Methods** are defined at the parent folder of the corresponding function and are implemented as aliases:

```go
    import (
        "github.com/valdotle/bit"       // methods for bitfield of type alias int
        "github.com/valdotle/bit/field" // functions for bitfield of type int
    )

    flag := 2
    field.Has(9, flag)     // false

    bit.Field(9).Has(flag) // false
```

Additionally, function and method name **suffixes** determine the flag's type:

```go
    import stringfield "github.com/valdotle/bit/field/string"

    field := "9"

    stringfield.Has(field, 2)                // flag of type int

    stringfield.HasString(field, "2")        // flag of type string

    stringfield.HasBig(field, big.NewInt(2)) // flag of type big.Int
```

Note, that **`int` bitfields don't allow checking against flags of type `string` or `big.Int`** (meaning they don't expose `HasString()` nor `HasBig()`) as those are intended to be used where integer limitations are exceeded and as such can't be parsed as integers anymore.
