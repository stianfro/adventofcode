package main

func main() {
	// first := race(7, 9)
	// second := race(15, 40)
	// third := race(30, 200)
	first := race(49, 263)
	second := race(97, 1532)
	third := race(94, 1378)
	fourth := race(94, 1851)

	answer := first * second * third * fourth
	println(answer)
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
