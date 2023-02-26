package brepadapter

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brep_adapter.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/pkg/geomabs"
	"github.com/marcuswu/gooccwrapper/pkg/gp"
	"github.com/marcuswu/gooccwrapper/pkg/topods"
)

type Surface struct {
	surf C.BRepAdapterSurface
}

func NewSurface(face topods.Face) Surface {
	return Surface{C.BRepAdapterSurface_Init(C.TopoDSFace(face.Face))}
}

func NewSurfaceRestriction(face topods.Face, restriction bool) Surface {
	return Surface{C.BRepAdapterSurface_InitRestriction(C.TopoDSFace(face.Face), C.bool(restriction))}
}

func (s Surface) Plane() gp.Pln {
	return gp.NewPlnFromRef(gp.GPPln(C.BRepAdapterSurface_Plane(s.surf)))
}

func (s Surface) Type() geomabs.SurfaceType {
	return geomabs.SurfaceType(C.BRepAdapterSurface_Type(s.surf))
}

func (s Surface) Free() {
	C.BRepAdapterSurface_Free(s.surf)
}
