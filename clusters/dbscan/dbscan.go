// Package dbscan clusters incoming points into clusters with params (eps, minPoints).
package dbscan

import (
	"github.com/spatial-go/geoos/clusters"
	"github.com/spatial-go/geoos/space"
)

// DBSCAN in pseudocode (from http://en.wikipedia.org/wiki/DBSCAN):

// EpsFunction is a function that returns eps based on point pt
type EpsFunction func(pt space.Point) float64

// DBScan clusters incoming points into clusters with params (eps, minPoints)
//
// eps is clustering radius in km
// minPoints in minimum number of points in eps-neighborhood (density)
func DBScan(points clusters.PointList, eps float64, minPoints int) (clusterArray []clusters.Cluster, noise []int) {
	visited := make([]bool, len(points))
	members := make([]bool, len(points))
	clusterArray = []clusters.Cluster{}
	noise = []int{}
	C := 0
	kdTree := NewKDTree(points)

	// Our SphericalDistanceFast returns distance which is not multiplied
	// by EarthR * DegreeRad, adjust eps accordingly
	eps = eps / EarthR / DegreeRad

	for i := 0; i < len(points); i++ {
		if visited[i] {
			continue
		}
		visited[i] = true

		neighborPts := kdTree.InRange(points[i], eps, nil)
		if len(neighborPts) < minPoints {
			noise = append(noise, i)
		} else {
			// init cluster with center point
			cluster := clusters.Cluster{C: C, Points: []int{i}, PointList: []space.Point{points[i]}}
			members[i] = true
			C++
			// expandCluster goes here inline
			neighborUnique := make(map[int]int)
			for j := 0; j < len(neighborPts); j++ {
				neighborUnique[neighborPts[j]] = neighborPts[j]
			}

			for j := 0; j < len(neighborPts); j++ {
				k := neighborPts[j]
				if !visited[k] {
					visited[k] = true
					moreNeighbors := kdTree.InRange(points[k], eps, nil)
					if len(moreNeighbors) >= minPoints {
						for _, p := range moreNeighbors {
							if _, ok := neighborUnique[p]; !ok {
								neighborPts = append(neighborPts, p)
								neighborUnique[p] = p
							}
						}
					}
				}

				if !members[k] {
					cluster.Points = append(cluster.Points, k)
					cluster.PointList = append(cluster.PointList, points[k])
					members[k] = true
				}
			}
			cluster.Recenter()
			clusterArray = append(clusterArray, cluster)
		}
	}
	return
}

// RegionQuery is simple way O(N) to find points in neighborhood
//
// It is roughly equivalent to kdTree.InRange(points[i], eps, nil)
func RegionQuery(points clusters.PointList, P space.Point, eps float64) []int {
	result := []int{}

	for i := 0; i < len(points); i++ {
		// if points[i].sqDist(P) < eps*eps {
		dis := DistanceSphericalFast(points[i], P)
		if dis < eps*eps {
			result = append(result, i)
		}
	}
	return result
}
