package trucktour

func truckTour(petrolpumps [][]int32) int {
	var fuel, ind int
	for i := 0; i < len(petrolpumps); i++ {
		fuel += int(petrolpumps[i][0] - petrolpumps[i][1])
		if fuel < 0 {
			ind = i + 1
			fuel = 0
		}
	}
	return ind

}
