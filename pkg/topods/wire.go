package topods

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/occ_types.h>
import "C"

type TopoDSWire C.TopoDSWire

type Wire struct {
	Wire C.TopoDSWire
}

func NewWireFromRef(ref TopoDSWire) Wire {
	return Wire{C.TopoDSWire(ref)}
}
