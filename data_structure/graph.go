package data_structure

import "fmt"

// Graph 代表有向图的数据结构
type Graph struct {
	vertices map[string][]string
}

// NewGraph 创建一个新的图
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string][]string),
	}
}

// AddEdge 添加有向边
func (g *Graph) AddEdge(from, to string) {
	g.vertices[from] = append(g.vertices[from], to)
}

// PrintGraph 打印图的邻接列表
func (g *Graph) PrintGraph() {
	for vertex, edges := range g.vertices {
		fmt.Printf("%s -> %v\n", vertex, edges)
	}
}
