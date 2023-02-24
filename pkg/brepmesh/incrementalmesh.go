package brepmesh

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepmesh.h>
import "C"
import "github.com/marcuswu/gooccwrapper/pkg/topods"

type IncrementalMesh struct {
	IncrementalMesh C.BRepMeshIncrementalMesh
}

func NewIncrementalMesh(
	shape topods.Shape,
	linearDefl float64,
	isRelative bool,
	angularDefl float64,
	isParallel bool,
) IncrementalMesh {
	return IncrementalMesh{C.BRepMeshIncrementalMesh_Init(
		C.TopoDSShape(shape.Shape),
		C.double(linearDefl),
		C.bool(isRelative),
		C.double(angularDefl),
		C.bool(isParallel),
	)}
}

func (im IncrementalMesh) Free() {
	C.BRepMeshIncrementalMesh_Free(im.IncrementalMesh)
}
