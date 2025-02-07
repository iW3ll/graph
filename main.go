package main

import (
    "fmt"
)

// Graph structure
type Graph struct {
    vertices []*Vertex
}

// Vertex structure
type Vertex struct {
    key      int
    adjacent []*Vertex
}

// Add a vertex to the graph
func (g *Graph) AddVertex(k int) {
    if contains(g.vertices, k) {
        err := fmt.Errorf("Vertice %v não adicionado pois ja existe a chave", k)
        fmt.Println(err.Error())
    } else {
        g.vertices = append(g.vertices, &Vertex{key: k})
    }
}

// Add an edge between two vertices
func (g *Graph) AddEdge(from, to int) {
    fromVertex := g.getVertex(from)
    toVertex := g.getVertex(to)
    if fromVertex == nil || toVertex == nil {
        err := fmt.Errorf("Aresta invalida (%v-->%v)", from, to)
        fmt.Println(err.Error())
    } else if contains(fromVertex.adjacent, to) {
        err := fmt.Errorf("Aresta já existe (%v-->%v)", from, to)
        fmt.Println(err.Error())
    } else {
        fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
    }
}

// Get a vertex by its key
func (g *Graph) getVertex(k int) *Vertex {
    for i, v := range g.vertices {
        if v.key == k {
            return g.vertices[i]
        }
    }
    return nil
}

// Check if a vertex with the given key exists in the slice
func contains(s []*Vertex, k int) bool {
    for _, v := range s {
        if k == v.key {
            return true
        }
    }
    return false
}

// Print the graph
func (g *Graph) Print() {
    for _, v := range g.vertices {
        fmt.Printf("\nVertice %v : ", v.key)
        for _, adj := range v.adjacent {
            fmt.Printf(" %v", adj.key)
        }
    }
    fmt.Println()
}

// Calculate in degree and out degree for each vertex
func (g *Graph) calculateDegrees() (map[int]int, map[int]int) {
    inDegree := make(map[int]int)
    outDegree := make(map[int]int)

    // Initialize degrees for all vertices
    for _, vertex := range g.vertices {
        inDegree[vertex.key] = 0
        outDegree[vertex.key] = 0
    }

    // Calculate in degree and out-degree
    for _, vertex := range g.vertices {
        for _, adj := range vertex.adjacent {
            outDegree[vertex.key]++
            inDegree[adj.key]++
        }
    }

    return inDegree, outDegree
}

// Check if the graph is connected
func (g *Graph) isConnected() bool {
    if len(g.vertices) == 0 {
        return true // An empty graph is trivially connected
    }

    visited := make(map[int]bool)

    // Find the first non-isolated vertex
    var start *Vertex
    for _, vertex := range g.vertices {
        if len(vertex.adjacent) > 0 {
            start = vertex
            break
        }
    }

    if start == nil {
        return true // All vertices are isolated
    }

    // Perform DFS from the start vertex
    var dfs func(*Vertex)
    dfs = func(v *Vertex) {
        visited[v.key] = true
        for _, adj := range v.adjacent {
            if !visited[adj.key] {
                dfs(adj)
            }
        }
    }

    dfs(start)

    // Check if all non-isolated vertices were visited
    for _, vertex := range g.vertices {
        if len(vertex.adjacent) > 0 && !visited[vertex.key] {
            return false
        }
    }

    return true
}

// Determine if the graph is eulerian
func (g *Graph) isEulerian() string {
    if !g.isConnected() {
        return "Não é Eureliano: Grafo não está conectado"
    }

    inDegree, outDegree := g.calculateDegrees()
    oddCount := 0

    for _, vertex := range g.vertices {
        totalDegree := inDegree[vertex.key] + outDegree[vertex.key]
        if totalDegree%2 != 0 {
            oddCount++
        }
    }
    //All vertices have the same degree
    if oddCount == 0 {
        return "É Eureliano"   
    } else if oddCount == 2 {
       return "Semi-Eureliano: Exatamente dois vertices tem graus impares"
     //Two or more than two vertices have odd degree
    } else {
        return "Não é Eureliano"
    }
}

func main() {
    test := &Graph{}
    for i := 0; i < 4; i++ {
        test.AddVertex(i)
    }
    test.AddEdge(0, 1)
    test.AddEdge(1, 2)
    test.AddEdge(2, 3)
    test.AddEdge(3, 0)
    
    test.Print()

    // Check if the graph is eulerian
    result := test.isEulerian()
    fmt.Println(result)
}