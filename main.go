package main

import "fmt"

type Data struct {
	TotalMul int
}

func main() {

	time := []int{53, 91, 67, 68}
	distance := []int{250, 1330, 1081, 1025}

	data := Data{TotalMul: 1}

	//main loop
	for i := 0; i < 4; i++ {
		resolve(&data, time[i], distance[i])
	}
	fmt.Println(data.TotalMul)
}

func resolve(data *Data, time int, distance int) {

	sum := 0
	//peak := false

	//hold button
	for timeHeld := 1; timeHeld < time; timeHeld++ {

		remainTime := time - timeHeld
		newDistance := timeHeld * remainTime
		if newDistance >= distance {
			sum++
		}

	}

	data.TotalMul = data.TotalMul * sum
}
