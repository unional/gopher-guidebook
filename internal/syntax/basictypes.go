package syntax

var Bool bool = true
var Str string = "abc"

var U8 uint8 = 16*16 - 1
var U16 uint16 = 256*256 - 1
var U32 uint32 = 65536*65536 - 1
var U64 uint64 = 1 << 63

// `int` == `int32` | `int64` depending on the system
var Int int = 1<<63 - 1

// `uint` == `uint32` | `uint64` depending on the system
var Uint uint = 1 << 63

// `uintptr` is 32bits or 64bits depending on the system
// large enough to host any pointer
var Uptr uintptr = 0xFFFFFFFFFFFFFFFF

var F32 float32 = 3.4e38
var F64 float64 = 1.79e308

var C64 complex64 = 3.4e38 + 3.4e38i
var C128 complex128 = 1.79e308 + 1.79e308i
