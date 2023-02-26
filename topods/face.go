package topods

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/occ_types.h>
import "C"

type TopoDSFace C.TopoDSFace

type Face struct {
	Face C.TopoDSFace
}

func NewFaceFromRef(ref TopoDSFace) Face {
	return Face{C.TopoDSFace(ref)}
}
