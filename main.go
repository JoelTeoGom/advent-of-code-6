package main

import (
	"fmt"
	"time"
)

type Data struct {
	TotalMul int
}

func main() {

	// timeInput := []int{53, 91, 67, 68}
	// distance := []int{250, 1330, 1081, 1025}

	data := Data{TotalMul: 1}

	// //main loop
	// for i := 0; i < 4; i++ {
	// 	resolve(&data, timeInput[i], distance[i])
	// }

	// //part 2
	// data.TotalMul = 1

	start := time.Now()
	resolve(&data, 53916768, 250133010811025)
	duration := time.Since(start)

	fmt.Println("Exe time:", duration)
	fmt.Println(data.TotalMul)

}
func resolve(data *Data, time int, distance int) {
	sum := 0
	max := 0
	//hold button
	for timeHeld := 1; timeHeld < time; timeHeld++ {
		remainTime := time - timeHeld
		newDistance := timeHeld * remainTime
		if newDistance >= distance {

			if newDistance > max {
				max = newDistance
				sum++
			} else {
				data.TotalMul = (sum * 2) - 1
				return
			}

		}
	}

}
