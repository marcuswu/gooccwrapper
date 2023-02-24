package topods

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/topods_compound.h>
import "C"

type Compound struct {
	Compound C.TopoDSCompound
}

func NewCompound() Compound {
	return Compound{C.TopoDSCompound_Init()}
}

func (c Compound) Free() {
	C.TopoDSCompound_Free(c.Compound)
}
