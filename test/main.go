package util

import (
	"container/heap"
)

type p struct {
	point int
	len   float64
}
type P []p

func (t *P) Len() int {
	return len(*t) //
}

func (t *P) Less(i, j int) bool {
	return (*t)[i].len < (*t)[j].len
}

func (t *P) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *P) Push(x interface{}) {
	*t = append(*t, x.(p))
}

func (t *P) Pop() interface{} {
	n := len(*t)
	x := (*t)[n-1]
	*t = (*t)[:n-1]
	return x
}

type node struct {
	to, next, id int
	cost         float64
}

var edge []node

var cnt int
var head []int  //由点进入，内存了该点的所有边
var d []float64 //记录距离
var dpPoint []int
var dpEdge []int

var EdgeTotal int
var PointTotal int

var resPoint [][][]int
var resEdge [][][]int
var resLength [][]float64

func add_edge(a int, b int, id int, c float64) {
	edge[cnt] = node{b, head[a], id, c}
	head[a] = cnt
	cnt++
}

func InitPath() {

	edge = make([]node, EdgeTotal*2, EdgeTotal*2)
	head = make([]int, PointTotal, PointTotal)
	d = make([]float64, PointTotal, PointTotal)
	dpPoint = make([]int, PointTotal, PointTotal)
	dpEdge = make([]int, PointTotal, PointTotal)
	resPoint = make([][][]int, PointTotal, PointTotal)
	resEdge = make([][][]int, PointTotal, PointTotal)
	resLength = make([][]float64, PointTotal, PointTotal)

	cnt = 0

	for i := range head {
		head[i] = -1
	}
	add_edge(2, 3, int(edges[i].ID), edges[i].Length)
	add_edge(edges[i].PointB, edges[i].PointA, int(edges[i].ID), edges[i].Length)

	for i := 0; i < PointTotal; i++ {
		for j := 0; j < PointTotal; j++ {
			resLength[i][j] = dijkstra(i, j)
		}
	}
}

func dijkstra(start int, end int) float64 {
	que := &P{}
	heap.Init(que)
	for i := 0; i < PointTotal; i++ {
		d[i] = 2020611073
		dpPoint[i] = -1
		dpEdge[i] = -1
	}
	d[start] = 0
	for que.Len() != 0 {
		var pp p
		pp = heap.Pop(que).(p)
		if d[pp.point] < pp.len {
			continue
		}
		for i := head[pp.point]; i != -1; i = edge[i].next { //遍历这个点的所有边
			if edge[i].cost+d[pp.point] < d[edge[i].to] { //对出现更优距离进行处理
				d[edge[i].to] = edge[i].cost + d[pp.point]
				dpPoint[edge[i].to] = pp.point
				dpEdge[edge[i].to] = edge[i].id
				heap.Push(que, p{edge[i].to, d[edge[i].to]})
			}
		}
	}

	for i := end; i != -1; i = dpPoint[i] {
		resPoint[end][start] = append(resPoint[end][start], i)
	}

	for i := dpEdge[end]; i != -1; i = dpEdge[i] {
		resEdge[end][start] = append(resEdge[end][start], i)
	}
	return d[end]
}
