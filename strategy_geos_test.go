package geos

import (
	"reflect"
	"testing"
)

func TestGEOSAlgorithm_Centroid(t *testing.T) {
	const multipoint = `MULTIPOINT ( -1 0, -1 2, -1 3, -1 4, -1 7, 0 1, 0 3, 1 1, 2 0, 6 0, 7 8, 9 8, 10 6 )`
	geometry, _ := UnmarshalString(multipoint)
	const pointresult = `POINT(2.3076923076923075 3.3076923076923075)`

	type args struct {
		g Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.

		{name: "point", args: args{g: geometry}, want: pointresult, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.Centroid(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Centroid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			s := MarshalString(got)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("Centroid() got = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_IsSimple(t *testing.T) {
	const polygon = `POLYGON((1 2, 3 4, 5 6, 1 2))`
	const linestring = `LINESTRING(1 1,2 2,2 3.5,1 3,1 2,2 1)`
	poly, _ := UnmarshalString(polygon)
	line, _ := UnmarshalString(linestring)

	type args struct {
		g Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "polygon", args: args{g: poly}, want: true, wantErr: false},
		{name: "line", args: args{g: line}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.IsSimple(tt.args.g)
			t.Log(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsSimple() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsSimple() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_Area(t *testing.T) {
	const wkt = `POLYGON((-1 -1, 1 -1, 1 1, -1 1, -1 -1))`
	geometry, _ := UnmarshalString(wkt)
	type args struct {
		g Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "area", args: args{g: geometry}, want: 4.0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.Area(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Area() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_Boundary(t *testing.T) {
	const sourceLine = `LINESTRING(1 1,0 0, -1 1)`
	const expectLine = `MULTIPOINT(1 1,-1 1)`
	sLine, _ := UnmarshalString(sourceLine)
	eLine, _ := UnmarshalString(expectLine)

	const sourcePolygon = `POLYGON((1 1,0 0, -1 1, 1 1))`
	const expectPolygon = `LINESTRING(1 1,0 0,-1 1,1 1)`
	sPolygon, _ := UnmarshalString(sourcePolygon)
	ePolygon, _ := UnmarshalString(expectPolygon)

	const multiPolygon = `POLYGON (( 10 130, 50 190, 110 190, 140 150, 150 80, 100 10, 20 40, 10 130 ),
	( 70 40, 100 50, 120 80, 80 110, 50 90, 70 40 ))`
	const expectMultiPolygon = `MULTILINESTRING((10 130,50 190,110 190,140 150,150 80,100 10,20 40,10 130),
	(70 40,100 50,120 80,80 110,50 90,70 40))`

	smultiPolygon, _ := UnmarshalString(sourceLine)
	emultiPolygon, _ := UnmarshalString(expectLine)

	type args struct {
		g Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    Geometry
		wantErr bool
	}{
		{name: "line", args: args{g: sLine}, want: eLine, wantErr: false},
		{name: "polygon", args: args{g: sPolygon}, want: ePolygon, wantErr: false},
		{name: "multiPolygon", args: args{g: smultiPolygon}, want: emultiPolygon, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.Boundary(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Boundary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(MarshalString(got))
			t.Log(MarshalString(tt.want))
			if !reflect.DeepEqual(MarshalString(got), MarshalString(tt.want)) {
				t.Errorf("Boundary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_Length(t *testing.T) {
	const line = `LINESTRING(0 0, 1 1)`
	geometry, _ := UnmarshalString(line)
	type args struct {
		g Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "lengh", args: args{g: geometry}, want: 1.4142135623730951, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.Length(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Length() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Length() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_HausdorffDistance(t *testing.T) {
	const g1 = `LINESTRING (0 0, 2 0)`
	const g2 = `MULTIPOINT (0 1, 1 0, 2 1)`
	geom1, _ := UnmarshalString(g1)
	geom2, _ := UnmarshalString(g2)

	type args struct {
		g1 Geometry
		g2 Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "HausdorffDistance", args: args{
			g1: geom1,
			g2: geom2,
		}, want: 1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.HausdorffDistance(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("HausdorffDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HausdorffDistance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_IsEmpty(t *testing.T) {
	const wkt = `POLYGON((1 2, 3 4, 5 6, 1 2))`
	geometry, _ := UnmarshalString(wkt)
	type args struct {
		g Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "empty", args: args{g: geometry}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.IsEmpty(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsEmpty() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_Crosses(t *testing.T) {
	const g1 = `LINESTRING(0 0, 10 10)`
	const g2 = `LINESTRING(10 0, 0 10)`

	geom1, _ := UnmarshalString(g1)
	geom2, _ := UnmarshalString(g2)
	type args struct {
		g1 Geometry
		g2 Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "crosses", args: args{
			g1: geom1,
			g2: geom2,
		}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.Crosses(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Crosses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Crosses() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_Within(t *testing.T) {
	const polygon = `POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))`
	const point1 = `POINT(3 3)`
	const point2 = `POINT(-1 35)`

	p1, _ := UnmarshalString(point1)
	p2, _ := UnmarshalString(point2)
	poly, _ := UnmarshalString(polygon)

	type args struct {
		g1 Geometry
		g2 Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "in", args: args{
			g1: p1,
			g2: poly,
		}, want: true, wantErr: false},
		{name: "notin", args: args{
			g1: p2,
			g2: poly,
		}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.Within(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Within() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Within() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_Contains(t *testing.T) {
	const polygon = `POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))`
	const point1 = `POINT(3 3)`
	const point2 = `POINT(-1 35)`

	p1, _ := UnmarshalString(point1)
	p2, _ := UnmarshalString(point2)
	poly, _ := UnmarshalString(polygon)
	type args struct {
		g1 Geometry
		g2 Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "contain", args: args{
			g1: poly,
			g2: p1,
		}, want: true, wantErr: false},
		{name: "notcontain", args: args{
			g1: poly,
			g2: p2,
		}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.Contains(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Contains() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_UniquePoints(t *testing.T) {
	const polygon = `POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))`
	const multipoint = `MULTIPOINT((0 0),(6 0),(6 6),(0 6))`

	poly, _ := UnmarshalString(polygon)

	type args struct {
		g Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "uniquepoints", args: args{g: poly}, want: multipoint, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.UniquePoints(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("UniquePoints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			res := MarshalString(got)
			t.Log(res)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("UniquePoints() got = %v, want %v", res, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_SharedPaths(t *testing.T) {
	const mullinestring = `MULTILINESTRING((26 125,26 200,126 200,126 125,26 125),
	   (51 150,101 150,76 175,51 150))`
	const linestring = `LINESTRING(151 100,126 156.25,126 125,90 161, 76 175)`
	const res = `GEOMETRYCOLLECTION (MULTILINESTRING ((126.0000000000000000 156.2500000000000000, 126.0000000000000000 125.0000000000000000), (101.0000000000000000 150.0000000000000000, 90.0000000000000000 161.0000000000000000), (90.0000000000000000 161.0000000000000000, 76.0000000000000000 175.0000000000000000)), MULTILINESTRING EMPTY)`

	mulline, _ := UnmarshalString(mullinestring)
	line, _ := UnmarshalString(linestring)

	type args struct {
		g1 Geometry
		g2 Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "sharepath", args: args{
			g1: line,
			g2: mulline,
		}, want: res, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.SharedPaths(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("SharedPaths() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SharedPaths() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_Snap(t *testing.T) {
	const input = `POINT(0.05 0.05)`
	const refernce = `POINT(0 0)`
	const expect = `POINT(0 0)`

	inputGeom, _ := UnmarshalString(input)
	referenceGeom, _ := UnmarshalString(refernce)

	type args struct {
		input     Geometry
		reference Geometry
		tolerance float64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "snap", args: args{
			input:     inputGeom,
			reference: referenceGeom,
			tolerance: 0.1,
		}, want: expect, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.Snap(tt.args.input, tt.args.reference, tt.args.tolerance)

			s := MarshalString(got)
			t.Log(s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Snap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("Snap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_Buffer(t *testing.T) {
	geometry, _ := UnmarshalString("POINT(100 90)")
	expectGeometry, _ := UnmarshalString("POLYGON((150 90,146.193976625564 70.8658283817455,135.355339059327 54.6446609406727,119.134171618255 43.8060233744357,100 40,80.8658283817456 43.8060233744356,64.6446609406727 54.6446609406725,53.8060233744357 70.8658283817454,50 89.9999999999998,53.8060233744356 109.134171618254,64.6446609406725 125.355339059327,80.8658283817453 136.193976625564,99.9999999999998 140,119.134171618254 136.193976625564,135.355339059327 125.355339059328,146.193976625564 109.134171618255,150 90))")
	type args struct {
		g        Geometry
		width    float64
		quadsegs int32
	}
	tests := []struct {
		name string
		G    GEOSAlgorithm
		args args
		want Geometry
	}{
		{name: "buffer", args: args{g: geometry, width: 50, quadsegs: 4}, want: expectGeometry},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			gotGeometry := G.Buffer(tt.args.g, tt.args.width, tt.args.quadsegs)
			isEqual, _ := G.EqualsExact(gotGeometry, tt.want, 0.000001)
			if !isEqual {
				t.Errorf("GEOSAlgorithm.Buffer() = %v, want %v", MarshalString(gotGeometry), MarshalString(tt.want))
			}
		})
	}
}

func TestGEOSAlgorithm_EqualsExact(t *testing.T) {
	geometry1, _ := UnmarshalString("POINT(116.309878625564 40.0427783817455)")
	geometry2, _ := UnmarshalString("POINT(116.309878725564 40.0427783827455)")
	geometry3, _ := UnmarshalString("POINT(116.309877625564 40.0427783827455)")
	type args struct {
		g1        Geometry
		g2        Geometry
		tolerance float64
	}
	tests := []struct {
		name    string
		G       GEOSAlgorithm
		args    args
		want    bool
		wantErr bool
	}{
		{name: "equals exact", args: args{g1: geometry1, g2: geometry2, tolerance: 0.000001}, want: true, wantErr: false},
		{name: "not equals exact", args: args{g1: geometry1, g2: geometry3, tolerance: 0.000001}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.EqualsExact(tt.args.g1, tt.args.g2, tt.args.tolerance)
			if (err != nil) != tt.wantErr {
				t.Errorf("GEOSAlgorithm.EqualsExact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GEOSAlgorithm.EqualsExact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGEOSAlgorithm_NGeometry(t *testing.T) {
	multiPoint, _ := UnmarshalString(`MULTIPOINT ( -1 0, -1 2, -1 3, -1 4, -1 7, 0 1, 0 3, 1 1, 2 0, 6 0, 7 8, 9 8, 10 6 )`)
	multiLineString, _ := UnmarshalString(`MULTILINESTRING((10 130,50 190,110 190,140 150,150 80,100 10,20 40,10 130),
	(70 40,100 50,120 80,80 110,50 90,70 40))`)
	multiPolygon, _ := UnmarshalString(`MULTIPOLYGON (((40 40, 20 45, 45 30, 40 40)),
	((20 35, 10 30, 10 10, 30 5, 45 20, 20 35)),
	((30 20, 20 15, 20 25, 30 20)))`)
	type args struct {
		g Geometry
	}
	tests := []struct {
		name    string
		G       GEOSAlgorithm
		args    args
		want    int
		wantErr bool
	}{
		{name: "ngeometry multiPoint", args: args{g: multiPoint}, want: 13, wantErr: false},
		{name: "ngeometry multiLineString", args: args{g: multiLineString}, want: 2, wantErr: false},
		{name: "ngeometry multiPolygon", args: args{g: multiPolygon}, want: 3, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := GEOSAlgorithm{}
			got, err := G.NGeometry(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("GEOSAlgorithm.NGeometry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GEOSAlgorithm.NGeometry() = %v, want %v", got, tt.want)
			}
		})
	}
}
