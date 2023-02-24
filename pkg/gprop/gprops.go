package gprop

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/gprop.h>
import "C"
import "github.com/marcuswu/gooccwrapper/pkg/gp"

type GProps struct {
	Props C.GPropGProps
}

func NewGProps() GProps {
	return GProps{C.GPropGProps_Init()}
}

func (g GProps) Free() {
	C.GPropGProps_Free(g.Props)
}

func (g GProps) Mass() float64 {
	return float64(C.GPropGProps_Mass(g.Props))
}

func (g GProps) CenterOfMass() gp.Pnt {
	return gp.NewPntFromRef(gp.GPPnt(C.GPropGProps_CentreOfMass(g.Props)))
}
