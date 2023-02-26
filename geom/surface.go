package geom

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/occ_types.h>
// #include <occwrapper/brep_tool.h>
import "C"

type GeomSurface C.gcSurface

type Surface struct {
	Surface C.gcSurface
}

func NewSurfaceFromRef(ref GeomSurface) Surface {
	return Surface{C.gcSurface(ref)}
}

func (s Surface) Free() {
	C.GCSurface_Free(s.Surface)
}
