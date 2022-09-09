package util

import (
	"STU/common"
	"STU/vo"
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

var EdgeTotal int
var PointTotal int

func add_edge(a int, b int, id int, c float64) {
	edge[cnt] = node{b, head[a], id, c}
	head[a] = cnt
	cnt++
}

func InitPath() {
	db := common.GetDB()
	// 记录的总条数
	db.Model(vo.Point{}).Count(&PointTotal)

	var edges []vo.Edge
	db.Find(&edges)

	db.Model(vo.Edge{}).Count(&EdgeTotal)

	edge = make([]node, EdgeTotal*2, EdgeTotal*2)
	head = make([]int, PointTotal, PointTotal)
	d = make([]float64, PointTotal, PointTotal)

	cnt = 0

	for i := range head {
		head[i] = -1
	}

	for i := range edges {
		add_edge(edges[i].PointA-1, edges[i].PointB-1, int(edges[i].ID), edges[i].Length)
		add_edge(edges[i].PointB-1, edges[i].PointA-1, int(edges[i].ID), edges[i].Length)
	}
}

func dijkstra(start int, end int) (float64, *[]int, *[]int) {
	que := &P{}
	heap.Init(que)
	dpPoint := make([]int, PointTotal, PointTotal)
	dpEdge := make([]int, PointTotal, PointTotal)
	for i := 0; i < PointTotal; i++ {
		d[i] = 2020611073
		dpPoint[i] = -1
		dpEdge[i] = -1
	}
	d[start-1] = 0
	heap.Push(que, p{start - 1, 0})
	for que.Len() != 0 {
		pp := heap.Pop(que).(p)
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
	var Points, Edges []int
	for i, j := dpPoint[end-1], dpEdge[end-1]; i != -1; i, j = dpPoint[i], dpEdge[i] {
		Points = append(Points, i+1)
		Edges = append(Edges, j)
	}
	for i, j := 0, len(Points)-1; i < j; i, j = i+1, j-1 {
		Points[i], Points[j] = Points[j], Points[i]
	}
	for i, j := 0, len(Edges)-1; i < j; i, j = i+1, j-1 {
		Edges[i], Edges[j] = Edges[j], Edges[i]
	}
	return d[end-1], &Points, &Edges
}

func permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, n := range nums {
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}
	dfs([]int{})
	return res
}

func DFS(path *vo.Path) (float64, *[]int, *[]int) {
	if len(path.Location) == 0 {
		res, ResPoint, ResEdge := dijkstra(path.Start, path.End)
		*ResPoint = append(*ResPoint, path.End)
		return res, ResPoint, ResEdge
	}
	res, ResPoint, ResEdge := 2020611073.0, &[]int{}, &[]int{}
	PossiblePaths := permute(path.Location)
	for i := range PossiblePaths {
		Tmp, TmpPoint, TmpEdge := dijkstra(path.Start, PossiblePaths[i][0])
		for j := 1; j < len(PossiblePaths[i]); j++ {
			tmp, tmpPoint, tmpEdge := dijkstra(PossiblePaths[i][j-1], PossiblePaths[i][j])
			Tmp += tmp
			*TmpPoint = append(*TmpPoint, (*tmpPoint)[:]...)
			*TmpEdge = append(*TmpEdge, (*tmpEdge)[:]...)
		}
		tmp, tmpPoint, tmpEdge := dijkstra(PossiblePaths[i][len(PossiblePaths[i])-1], path.End)
		Tmp += tmp
		*TmpPoint = append(*TmpPoint, (*tmpPoint)[:]...)
		*TmpEdge = append(*TmpEdge, (*tmpEdge)[:]...)
		if Tmp < res {
			res, ResPoint, ResEdge = Tmp, TmpPoint, TmpEdge
		}
	}
	*ResPoint = append(*ResPoint, path.End)
	return res, ResPoint, ResEdge
}
