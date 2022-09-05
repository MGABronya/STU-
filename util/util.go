package util

import (
	"STU/common"
	"STU/vo"
	"container/heap"
	"fmt"
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
	db := common.GetDB()
	// 记录的总条数
	db.Model(vo.Point{}).Count(&PointTotal)

	var edges []vo.Edge
	db.Find(&edges)

	db.Model(vo.Edge{}).Count(&EdgeTotal)

	edge = make([]node, EdgeTotal*2, EdgeTotal*2)
	head = make([]int, PointTotal, PointTotal)
	d = make([]float64, PointTotal, PointTotal)
	dpPoint = make([]int, PointTotal, PointTotal)
	dpEdge = make([]int, PointTotal, PointTotal)
	resPoint = make([][][]int, PointTotal, PointTotal)
	resEdge = make([][][]int, PointTotal, PointTotal)
	resLength = make([][]float64, PointTotal, PointTotal)

	for i := 0; i < PointTotal; i++ {
		resPoint[i] = make([][]int, PointTotal, PointTotal)
		resEdge[i] = make([][]int, PointTotal, PointTotal)
		resLength[i] = make([]float64, PointTotal, PointTotal)
	}

	cnt = 0

	for i := range head {
		head[i] = -1
	}

	for i := range edges {
		add_edge(edges[i].PointA, edges[i].PointB, int(edges[i].ID), edges[i].Length)
		add_edge(edges[i].PointB, edges[i].PointA, int(edges[i].ID), edges[i].Length)
	}
	for i := 0; i < PointTotal; i++ {
		for j := 0; j < PointTotal; j++ {
			resLength[i][j] = dijkstra(i + 1, j + 1)
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
	heap.Push(que, p{start, 0})
	for que.Len() != 0 {
		pp := heap.Pop(que).(p)
		fmt.Print(pp)
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

func DFS(ResPoint *[]int, ResEdge *[]int, path *vo.Path) float64 {
	*ResPoint = make([]int, len(path.End), len(path.End))
	(*ResPoint)[0] = path.Start
	PossiblePaths := permute(path.End)
	var res float64 = 2020611073
	for i := range PossiblePaths {
		tmp := resLength[path.Start][PossiblePaths[i][0]]
		for j := 1; j < len(PossiblePaths[i]); j++ {
			tmp += resLength[PossiblePaths[i][j-1]][PossiblePaths[i][j]]
		}
		if tmp < res {
			res = tmp
			for j := 0; j < len(PossiblePaths[i]); j++ {
				(*ResPoint)[j+1] = PossiblePaths[i][j]
			}
			(*ResEdge) = make([]int, 0)
			for _, v := range resEdge[path.Start][PossiblePaths[i][0]] {
				(*ResEdge) = append(*ResEdge, v)
			}
			for j := 1; j < len(PossiblePaths[i]); j++ {
				for _, v := range resEdge[PossiblePaths[i][j-1]][PossiblePaths[i][j]] {
					(*ResEdge) = append(*ResEdge, v)
				}
			}
		}
	}
	return res
}
