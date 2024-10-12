package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := getData()
	n := float64(len(data))
	var x, y, xy, x2, y2 float64
	for i := 0; i < len(data); i++ {
		fi := float64(i)
		x += fi
		y += data[i]
		xy += fi * data[i]
		x2 += fi * fi
		y2 += data[i] * data[i]
	}
	m := (n*xy - x*y) / (n*x2 - x*x)
	b := (y - m*x) / n
	numerator := (n*xy - x*y)
	denominator := math.Sqrt((n*x2 - x*x) * (n*y2 - y*y))
	r := numerator / denominator
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}

func getData() []float64 {
	arg := os.Args[1:]
	if len(arg) != 1 {
		log.Fatal("wrong arguments")
	}
	filePath := arg[0]
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var intSlice []float64
	number := strings.Split(string(data), "\n")
	for i := 0; i < len(number); i++ {
		if number[i] != "" {
			nb, err := strconv.Atoi(number[i])
			if err != nil {
				fmt.Println(number[i])
				log.Fatal(err)
			}
			intSlice = append(intSlice, float64(nb))
		}
	}

	return intSlice
}
