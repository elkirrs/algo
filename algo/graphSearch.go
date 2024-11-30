package algo

import "fmt"

func BreadthSearch() {
	var graph = map[string][]string{
		"a": {"b", "c"},
		"b": {"f"},
		"c": {"d", "e"},
		"d": {"f"},
		"e": {"f"},
		"f": {"g"},
	}
	breadthSearchRun(graph, "a", "f")
}

func breadthSearchRun(graph map[string][]string, start, end string) bool {
	queue := []string{start}
	visited := map[string]bool{}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == end {
			fmt.Printf("SearchWidth: start: %s end: %s result: true\n", start, end)
			return true
		}

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	fmt.Printf("SearchWidth: start: %s end: %s result: false\n", start, end)
	return false
}

func DijkstraSearch() {
	var graph = map[string]map[string]int{
		"a": {"b": 2, "c": 1},
		"b": {"f": 7},
		"c": {"d": 5, "e": 2},
		"d": {"f": 2},
		"e": {"f": 1},
		"f": {"g": 1},
		"g": {},
	}
	dijkstraSearchRun(graph, "a", "g")
}

func dijkstraSearchRun(graph map[string]map[string]int, start, end string) map[string]int {
	distances := make(map[string]int)
	previous := make(map[string]string)
	visited := make(map[string]bool)

	for node := range graph {
		if node == start {
			distances[node] = 0
		} else {
			distances[node] = int(1e9)
		}
	}

	for {
		node := findNodeLowestDistance(distances, visited)
		if node == "" || node == end {
			break
		}

		distance := distances[node]
		neighbors := graph[node]
		for neighbor, weight := range neighbors {
			if visited[neighbor] {
				continue
			}
			newDistance := distance + weight
			if newDistance < distances[neighbor] {
				distances[neighbor] = newDistance
				previous[neighbor] = node
			}
		}
		visited[node] = true
	}

	fmt.Println("Distances:", distances)
	fmt.Println("Path:", reconstructPath(previous, start, end))

	return distances
}

func findNodeLowestDistance(distances map[string]int, visited map[string]bool) string {
	lowestDistance := int(1e9)
	lowestNode := ""

	for node, distance := range distances {
		if distance < lowestDistance && !visited[node] {
			lowestDistance = distance
			lowestNode = node
		}
	}

	return lowestNode
}

func reconstructPath(previous map[string]string, start, end string) []string {
	var path []string
	for node := end; node != ""; node = previous[node] {
		path = append([]string{node}, path...)
		if node == start {
			break
		}
	}
	return path
}
