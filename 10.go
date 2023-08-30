package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	thermoMapa := map[int][]float32{}
	for _, v := range arr {
		//Делим на 10 чтобы получить целочисленное (-2, 1, 2 и т.д)
		//После получения целочислонного умножаем обратно на 10 и получаем группы(-20, 10, 20 и т.д)
		//Записываем под ключом полученной группы
		key := int(math.Trunc(float64(v))/10) * 10
		thermoMapa[key] = append(thermoMapa[key], v)
	}
	fmt.Println(thermoMapa)
}
