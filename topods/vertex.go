package topods

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/occ_types.h>
// #include <occwrapper/brep_tool.h>
import "C"
import "github.com/marcuswu/gooccwrapper/gp"

type TopoDSVertex C.TopoDSVertex

type Vertex struct {
	Vertex C.TopoDSVertex
}

func NewVertexFromRef(ref TopoDSVertex) Vertex {
	return Vertex{C.TopoDSVertex(ref)}
}

func (v Vertex) Pnt() gp.Pnt {
	return gp.NewPntFromRef(gp.GPPnt(C.BRepTool_Pnt(C.TopoDSVertex(v.Vertex))))
}

// TODO: need to be able to get X, Y, Z values

// I didn't define this in the C API -- seems like an oversight
// func (v Vertex) Free() {
// 	C.TopoDSVertex_Free(v.Vertex)
// }
