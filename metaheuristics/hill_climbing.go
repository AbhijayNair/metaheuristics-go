package metaheuristics

var upper_bound float64
var lower_bound float64

func getNeighbors(curr [2]float64, step_size float64) [][2]float64 {
	var neighbors [][2]float64

	for i := 0; i < 2; i++ {
		var temp [2]float64 = curr
		temp[i] += step_size
		temp[i] = max(min(temp[i], upper_bound), lower_bound)
		neighbors = append(neighbors, temp)
		temp = curr
		temp[i] -= step_size
		temp[i] = max(min(temp[i], upper_bound), lower_bound)
		neighbors = append(neighbors, temp)
	}
	var temp [2]float64 = curr
	temp[0] += step_size
	temp[0] = max(min(temp[0], upper_bound), lower_bound)
	temp[1] += step_size
	temp[1] = max(min(temp[1], upper_bound), lower_bound)
	neighbors = append(neighbors, temp)
	temp = curr
	temp[0] -= step_size
	temp[0] = max(min(temp[0], upper_bound), lower_bound)
	temp[1] -= step_size
	temp[1] = max(min(temp[1], upper_bound), lower_bound)
	neighbors = append(neighbors, temp)
	return neighbors

}

func solve() {}
