// @Title  util
// @Description  使用迪杰斯特拉算法以及全排列算法，用于计算最短路径
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:47
package util

import (
	"STU/common"
	"STU/vo"
	"container/heap"
)

// p			定义了用于迪杰斯特拉算法的结构体的基本信息
type p struct {
	point int     // 当前到达的点
	len   float64 // 当前到达该点所走过的路程
}
type P []p //P 用于最小堆

// @title    Len
// @description   返回最小堆中的元素个数
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     void void		入参为空
// @return    int      返回参数为最小堆中的元素个数
func (t *P) Len() int {
	return len(*t)
}

// @title    Less
// @description   定义结构体的对比方式
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     i, j int		入参为最小堆数组的两个下标
// @return    bool          返回对比后是否为true
func (t *P) Less(i, j int) bool {
	return (*t)[i].len < (*t)[j].len
}

// @title    Swap
// @description   定义结构体的交换规则
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     i, j int		入参为最小堆数组的两个下标
// @return    void          没有返回值
func (t *P) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

// @title    Push
// @description   定义如何向最小堆中添加元素
// @auth      MGAronya（张健）             2022-9-16 10:33
// @param     x interface{}		入参应当要可以转化为p
// @return    void          没有返回值
func (t *P) Push(x interface{}) {
	*t = append(*t, x.(p))
}

// @title    Pop
// @description   定义如何从最小堆中弹出元素
// @auth      MGAronya（张健）             2022-9-16 10:35
// @param     void			没有入参
// @return    interface{}   返回接口可以转化为p
func (t *P) Pop() interface{} {
	n := len(*t)
	x := (*t)[n-1]
	*t = (*t)[:n-1]
	return x
}

// node				定义了链式前向星节点的基本信息
type node struct {
	to, next, id int     // to为该边所能到达的点，next为链式前向星结构中的下一节点，id为当前边的id
	cost         float64 // 该节点所代表边的长度
}

var edge []node
var pointTran []int
var pointMap map[int]int

var cnt int
var head []int  //由点进入，内存了该点的所有边
var d []float64 //记录距离

var EdgeTotal int
var PointTotal int

// @title    add_edge
// @description   定义如何向链式前向星中添加边
// @auth      MGAronya（张健）             2022-9-16 10:41
// @param     a int, b int, id int, c float64			边的基本信息
// @return    void			没有返回值
func add_edge(a int, b int, id int, c float64) {
	// TODO 在链式前向星的cnt处记录该边
	edge[cnt] = node{b, head[a], id, c}
	// TODO 记录点a新的入口
	head[a] = cnt
	// TODO 更新cnt
	cnt++
}

// @title    InitPath
// @description   定义如何初始化链式前向星
// @auth      MGAronya（张健）             2022-9-16 10:41
// @param     void			没有入参
// @return    void			没有返回值
func InitPath() {
	db := common.GetDB()
	// TODO 记录的总条数
	db.Model(vo.Point{}).Count(&PointTotal)

	var edges []vo.Edge
	db.Find(&edges)

	db.Model(vo.Edge{}).Count(&EdgeTotal)

	edge = make([]node, EdgeTotal*2, EdgeTotal*2)
	head = make([]int, PointTotal, PointTotal)
	d = make([]float64, PointTotal, PointTotal)

	pointMap = make(map[int]int)
	pointTran = make([]int, PointTotal, PointTotal)

	cnt = 0
	c := 0

	// TODO 初始化所有点的入口
	for i := range head {
		head[i] = -1
	}

	// TODO 使用pointMap进行点的离散化处理，pointTran进行点的翻译
	for i := range edges {
		if _, ok := pointMap[edges[i].PointA]; !ok {
			pointMap[edges[i].PointA] = c
			pointTran[c] = edges[i].PointA
			c++
		}
		if _, ok := pointMap[edges[i].PointB]; !ok {
			pointMap[edges[i].PointB] = c
			pointTran[c] = edges[i].PointB
			c++
		}
	}

	// TODO 链式前向星建边
	for i := range edges {
		add_edge(pointMap[edges[i].PointA], pointMap[edges[i].PointB], int(edges[i].ID), edges[i].Length)
		add_edge(pointMap[edges[i].PointB], pointMap[edges[i].PointA], int(edges[i].ID), edges[i].Length)
	}
}

// @title    dijkstra
// @description   迪杰斯特拉算法，计算两点之间的最短路径
// @auth      MGAronya（张健）             2022-9-16 10:41
// @param     start int, end int		起点与终点
// @return    float64, *[]int, *[]int	路径的总长度，路径上的点，路径上的边
func dijkstra(start int, end int) (float64, *[]int, *[]int) {
	que := &P{}
	heap.Init(que)
	dpPoint := make([]int, PointTotal, PointTotal)
	dpEdge := make([]int, PointTotal, PointTotal)

	// TODO 初始化路径数组
	for i := 0; i < PointTotal; i++ {
		d[i] = 2020611073
		dpPoint[i] = -1
		dpEdge[i] = -1
	}

	d[start] = 0
	heap.Push(que, p{start, 0})

	// TODO 循环直到没有更优路径
	for que.Len() != 0 {
		pp := heap.Pop(que).(p)

		// TODO 当前路径长度高于记录值，则跳过此次循环
		if d[pp.point] < pp.len {
			continue
		}

		// TODO 遍历该点的所有边
		for i := head[pp.point]; i != -1; i = edge[i].next {
			// TODO 对更优距离进行处理
			if edge[i].cost+d[pp.point] < d[edge[i].to] {
				d[edge[i].to] = edge[i].cost + d[pp.point]
				dpPoint[edge[i].to] = pp.point
				dpEdge[edge[i].to] = edge[i].id
				heap.Push(que, p{edge[i].to, d[edge[i].to]})
			}
		}
	}

	var Points, Edges []int

	// TODO 取出记录的路径
	for i, j := dpPoint[end], dpEdge[end]; i != -1; i, j = dpPoint[i], dpEdge[i] {
		Points = append(Points, pointTran[i])
		Edges = append(Edges, j)
	}

	// TODO 因为是从终点回溯，此处需要将路径颠倒
	for i, j := 0, len(Points)-1; i < j; i, j = i+1, j-1 {
		Points[i], Points[j] = Points[j], Points[i]
	}

	// TODO 此处需要将路径颠倒
	for i, j := 0, len(Edges)-1; i < j; i, j = i+1, j-1 {
		Edges[i], Edges[j] = Edges[j], Edges[i]
	}
	return d[end], &Points, &Edges
}

// @title    permute
// @description   计算全排列
// @auth      MGAronya（张健）             2022-9-16 10:41
// @param     nums []int		原数组
// @return    [][]int			数组全排列后的所有结果
func permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)

	// TODO 通过递归计算全排列
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

// @title    permute
// @description   计算从起点，经过多个必经点，最后到达终点的最短路径
// @auth      MGAronya（张健）             2022-9-16 10:41
// @param     path *vo.Path		path结构体，包含起点，终点，必经点
// @return    float64, *[]int, *[]int		返回最短路径长度，最短路径经过的点，最短路径经过的边
func DFS(path *vo.Path) (float64, *[]int, *[]int) {
	path.Start = pointMap[path.Start]

	// TODO 先将传入的数据进行先前记录的离散化处理
	for _, i := range path.Location {
		path.Location[i] = pointMap[path.Location[i]]
	}
	path.End = pointMap[path.End]

	// TODO 如果只是简单的起点终点查询
	if len(path.Location) == 0 {
		res, ResPoint, ResEdge := dijkstra(path.Start, path.End)
		*ResPoint = append(*ResPoint, path.End)
		return res, ResPoint, ResEdge
	}
	res, ResPoint, ResEdge := 2020611073.0, &[]int{}, &[]int{}
	PossiblePaths := permute(path.Location)

	// TODO 遍历必经点的全排序，找出最佳次序
	for i := range PossiblePaths {
		Tmp, TmpPoint, TmpEdge := dijkstra(path.Start, PossiblePaths[i][0])
		// TODO 遍历当前必经点
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
		// TODO 如果找出了更佳的路径，进行路径的记录
		if Tmp < res {
			res, ResPoint, ResEdge = Tmp, TmpPoint, TmpEdge
		}
	}
	*ResPoint = append(*ResPoint, path.End)
	return res, ResPoint, ResEdge
}
