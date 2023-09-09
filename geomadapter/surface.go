package geomadapter

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/geom_adapter.h>
import "C"
import "github.com/marcuswu/gooccwrapper/geom"

type Surface struct {
	Surface C.GeomAdapterSurface
}

func NewSurface(surface geom.Surface) Surface {
	return Surface{C.GeomAdapterSurface_Init(C.gcSurface(surface.Surface))}
}

func (s Surface) Free() {
	C.GeomAdapterSurface_Free(s.Surface)
}

func (s Surface) IsConical() bool {
	return bool(C.GeomAdapterSurface_IsConical(s.Surface))
}

func (s Surface) IsCylindrical() bool {
	return bool(C.GeomAdapterSurface_IsCylindrical(s.Surface))
}

func (s Surface) IsPlanar() bool {
	return bool(C.GeomAdapterSurface_IsPlanar(s.Surface))
}
