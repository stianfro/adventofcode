package main

func main() {
	println(race(49979494, 263153213781851))
}

func race(time, distance int) int {
	var win int
	for i := 0; i < time; i++ {
		remainingTime := time - i
		distanceTraveled := remainingTime * i
		if distanceTraveled > distance {
			win++
		}
	}
	return win
}
