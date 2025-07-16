package topods

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/occ_types.h>
// #include <occwrapper/topexp_explorer.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/gp"
)

type TopoDSEdge C.TopoDSEdge

type Edge struct {
	Edge C.TopoDSEdge
}

func NewEdgeFromRef(ref TopoDSEdge) Edge {
	return Edge{C.TopoDSEdge(ref)}
}

func (e Edge) FirstVertex() Vertex {
	return NewVertexFromRef(TopoDSVertex(C.TopExp_FirstVertex(e.Edge)))
}

func (e Edge) LastVertex() Vertex {
	return NewVertexFromRef(TopoDSVertex(C.TopExp_LastVertex(e.Edge)))
}

func (e Edge) Midpoint() gp.Pnt {
	pnt1 := e.FirstVertex().Pnt()
	pnt2 := e.LastVertex().Pnt()
	return gp.NewMidPnt(pnt1, pnt2)
}
