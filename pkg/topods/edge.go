package topods

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/occ_types.h>
import "C"

type TopoDSEdge C.TopoDSEdge

type Edge struct {
	Edge C.TopoDSEdge
}

func NewEdgeFromRef(ref TopoDSEdge) Edge {
	return Edge{C.TopoDSEdge(ref)}
}
