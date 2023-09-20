// Package kdtree Contains classes which implement a k-D tree index over 2-D point data.
package kdtree

import (
	"fmt"
	"os"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
)

func TestKdTree_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"IsEmpty", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := indexTree
			if got := k.IsEmpty(); got != tt.want {
				t.Errorf("KdTree.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKdTree_InsertNoData(t *testing.T) {
	type args struct {
		p matrix.Matrix
	}
	tests := []struct {
		name string
		args args
	}{
		{"inset no data", args{matrix.Matrix{100, 100}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := indexTree
			if got := k.InsertNoData(tt.args.p); got == nil {
				t.Errorf("KdTree.InsertNoData() = %v, want %v", got, "not nil")
			}
		})
	}
}

func TestKdTree_Insert(t *testing.T) {
	type args struct {
		p    matrix.Matrix
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"insert ", args{matrix.Matrix{101, 101}, envelope.Matrix(matrix.Matrix{101, 101})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := indexTree
			if got := k.InsertMatrix(tt.args.p, tt.args.data); got == nil {
				t.Errorf("KdTree.Insert() = %v, want %v", got, "not nil")
			}
		})
	}
}

func TestKdTree_QueryVisitor(t *testing.T) {

	type args struct {
		queryEnv *envelope.Envelope
	}
	tests := []struct {
		name string
		args args
	}{
		{"inset ", args{envelope.Matrix(matrix.Matrix{3, 1})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := indexTree
			visitor := &BestMatchVisitor{Matrix: matrix.Matrix{3, 2}, tolerance: 3}
			if err := k.QueryVisitor(tt.args.queryEnv, visitor); err != nil {
				t.Errorf("KdTree.QueryVisitor() = %v, want %v", visitor, "not nil")
			}
		})
	}
}

func TestKdTree_Query(t *testing.T) {
	type args struct {
		queryPt matrix.Matrix
	}
	tests := []struct {
		name string
		args args
	}{
		{"query ", args{matrix.Matrix{2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := indexTree
			if got := k.QueryMatrix(tt.args.queryPt); got == nil {
				t.Errorf("KdTree.Query() = %v, want %v", got, "not nil")
			}
		})
	}
}

func TestKdTree_Depth(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Depth", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := buildTree()
			if got := k.Depth(); got != tt.want {
				t.Errorf("KdTree.Depth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKdTree_Size(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Size", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := buildTree()
			if got := k.Size(); got != tt.want {
				t.Errorf("KdTree.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

var indexTree *KdTree

func TestMain(m *testing.M) {

	fmt.Println("test start")
	buildTree()
	code := m.Run()
	fmt.Println("test end")
	os.Exit(code)
}

func buildTree() *KdTree {
	indexTree = &KdTree{}
	var ms matrix.Collection = matrix.Collection{
		matrix.Matrix{1, 1},
		matrix.Matrix{1, 1},
		matrix.Matrix{2, 1},
		matrix.Matrix{2, 2},
		matrix.Matrix{3, 1},
		matrix.Matrix{3, 2},
	}
	for i := 0; i < len(ms); i++ {
		indexTree.InsertMatrix(ms[i].(matrix.Matrix), nil)
	}
	return indexTree
}
