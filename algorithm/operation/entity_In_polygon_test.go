package operation

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestInPolygon(t *testing.T) {
	type args struct {
		point matrix.Matrix
		poly  matrix.LineMatrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{matrix.Matrix{5, 5}, matrix.LineMatrix{{1, 1}, {2, 1}, {2, 6}, {7, 6}, {7, 1}, {10, 1}, {10, 10}, {1, 10}, {1, 1}}}, false},
		{"test2", args{matrix.Matrix{5, 7}, matrix.LineMatrix{{1, 1}, {2, 1}, {2, 6}, {7, 6}, {7, 1}, {10, 1}, {10, 10}, {1, 10}, {1, 1}}}, true},
		{"test2", args{matrix.Matrix{5, 6}, matrix.LineMatrix{{1, 1}, {2, 1}, {2, 6}, {7, 6}, {7, 1}, {10, 1}, {10, 10}, {1, 10}, {1, 1}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPnPolygon(tt.args.point, tt.args.poly); got != tt.want {
				t.Errorf("InPolygon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pointInRing(t *testing.T) {
	type args struct {
		p matrix.Matrix
		r matrix.LineMatrix
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{matrix.Matrix{5, 5}, matrix.LineMatrix{{1, 1}, {2, 1}, {2, 6}, {7, 6}, {7, 1}, {10, 1}, {10, 10}, {1, 10}, {1, 1}}}, OnlyOutPolygon},
		{"test2", args{matrix.Matrix{5, 7}, matrix.LineMatrix{{1, 1}, {2, 1}, {2, 6}, {7, 6}, {7, 1}, {10, 1}, {10, 10}, {1, 10}, {1, 1}}}, OnlyInPolygon},
		{"test3", args{matrix.Matrix{5, 6}, matrix.LineMatrix{{1, 1}, {2, 1}, {2, 6}, {7, 6}, {7, 1}, {10, 1}, {10, 10}, {1, 10}, {1, 1}}}, OnlyInLine},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pointInRing(tt.args.p, tt.args.r)
			if got != tt.want {
				t.Errorf("pointInRing() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rayIntersectsSegment(t *testing.T) {
	type args struct {
		p matrix.Matrix
		a matrix.Matrix
		b matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"rayIntersectsSegment", args{matrix.Matrix{6.0, 7.0}, matrix.Matrix{6.0, 5.0}, matrix.Matrix{8.0, 5.0}}, false},
		{"rayIntersectsSegment", args{matrix.Matrix{6, 4}, matrix.Matrix{6, 5}, matrix.Matrix{8, 5}}, false},
		{"rayIntersectsSegment", args{matrix.Matrix{6, 7}, matrix.Matrix{6, 5}, matrix.Matrix{7, 8}}, true},
		{"rayIntersectsSegment", args{matrix.Matrix{7, 7}, matrix.Matrix{6, 5}, matrix.Matrix{7, 8}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rayIntersectsSegment(tt.args.p, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("rayIntersectsSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInPolygon(t *testing.T) {
	type args struct {
		arg  matrix.Steric
		poly matrix.PolygonMatrix
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"test line poly0", args{matrix.LineMatrix{{0.5, 1.5}, {2.5, 1.5}}, matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}}}, OnlyOutPolygon, BothPolygon},
		{"test line poly1", args{matrix.LineMatrix{{0.5, 1.5}, {1, 1.5}}, matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}}}, PartOutPolygon, OnlyOutPolygon},
		{"test line poly2", args{matrix.LineMatrix{{1, 1.5}, {2, 1.5}}, matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}}}, OnlyInLine, OnlyInPolygon},
		{"test line poly3", args{matrix.LineMatrix{{2, 1.5}, {2.5, 1.5}}, matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}}}, PartOutPolygon, OnlyOutPolygon},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := InPolygon(tt.args.arg, tt.args.poly)
			if got != tt.want {
				t.Errorf("IsInPolygon() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsInPolygon() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
