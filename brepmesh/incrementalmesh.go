package brepmesh

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepmesh.h>
import "C"
import "github.com/marcuswu/gooccwrapper/topods"

type IncrementalMesh struct {
	IncrementalMesh C.BRepMeshIncrementalMesh
}

func NewIncrementalMesh(
	shape topods.Compound,
	linearDefl float64,
	isRelative bool,
	angularDefl float64,
	isParallel bool,
) IncrementalMesh {
	return IncrementalMesh{C.BRepMeshIncrementalMesh_Init(
		C.TopoDSShape(shape.Compound),
		C.double(linearDefl),
		C.bool(isRelative),
		C.double(angularDefl),
		C.bool(isParallel),
	)}
}

func (im IncrementalMesh) Free() {
	C.BRepMeshIncrementalMesh_Free(im.IncrementalMesh)
}
