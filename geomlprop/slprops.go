package geomlprop

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/geomlprop_slprops.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/geom"
	"github.com/marcuswu/gooccwrapper/gp"
)

type SLProps struct {
	Props C.GeomLPropSLProps
}

func NewSLProps(surface geom.Surface, umin float64, vmin float64, n float64, res float64) SLProps {
	return SLProps{C.GeomLPropSLProps_Init(C.gcSurface(surface.Surface), C.double(umin), C.double(vmin), C.double(n), C.double(res))}
}

func (p SLProps) Normal() gp.Dir {
	return gp.NewDirFromRef(gp.GPDir(C.GeomLPropsSLProps_Normal(p.Props)))
}

func (p SLProps) Free() {
	C.GeomLPropSLProps_Free(p.Props)
}
