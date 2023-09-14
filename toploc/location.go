package toploc

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/toploc_location.h>
import "C"
import "github.com/marcuswu/gooccwrapper/gp"

type TopLocLocation C.TopLocLocation

type Location struct {
	loc C.TopLocLocation
}

func NewLocationFromRef(ref TopLocLocation) Location {
	return Location{C.TopLocLocation(ref)}
}

func (l Location) Transformation() gp.Trsf {
	return gp.NewTrsfFromRef(gp.GPTrsf(C.TopLocLocation_Transformation(l.loc)))
}

func (l Location) Free() {
	C.TopLocLocation_Free(l.loc)
}
