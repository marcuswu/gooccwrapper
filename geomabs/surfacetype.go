package geomabs

type SurfaceType int

// This list echos opencascade/GeomAbs_SurfaceType.hxx and must be updated if opencascade is updated
const (
	Plane SurfaceType = iota
	Cylinder
	Cone
	Sphere
	Torus
	BezierSurface
	BSplineSurface
	SurfaceOfRevolution
	SurfaceOfExtrusion
	OffsetSurface
	OtherSurface
)
