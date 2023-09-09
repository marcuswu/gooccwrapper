package brep

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brep_builder.h>
import "C"
import "github.com/marcuswu/gooccwrapper/topods"

type Builder struct {
	builder C.BRepBuilder
}

func NewBuilder() Builder {
	return Builder{C.BRepBuilder_Init()}
}

func (b Builder) MakeCompound(res topods.Compound) {
	C.BRepBuilder_MakeCompound(b.builder, C.TopoDSCompound(res.Compound))
}

func (b Builder) Add(res topods.Shape, shape topods.Shape) {
	C.BRepBuilder_Add(b.builder, C.TopoDSShape(res.Shape), C.TopoDSShape(shape.Shape))
}

func (b Builder) Free() {
	C.BRepBuilder_Free(b.builder)
}
