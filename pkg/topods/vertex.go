package topods

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/occ_types.h>
import "C"

type TopoDSVertex C.TopoDSVertex

type Vertex struct {
	Vertex C.TopoDSVertex
}

func NewVertexFromRef(ref TopoDSVertex) Vertex {
	return Vertex{C.TopoDSVertex(ref)}
}

// I didn't define this in the C API -- seems like an oversight
// func (v Vertex) Free() {
// 	C.TopoDSVertex_Free(v.Vertex)
// }
