package caixeiro_viajante

import (
	"fmt"
	"math"
)

func main() {
	distanceMatrix := [][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	}
	n := len(distanceMatrix)
	minRoute, minCost := travellingSalesmanBruteForce(distanceMatrix, n)
	fmt.Println("Menor custo:", minCost)
	fmt.Println("Rota:", minRoute)
}

func travellingSalesmanBruteForce(distanceMatrix [][]int, n int) ([]int, int) {
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i
	}

	minCost := math.MaxInt32
	var minRoute []int

	permute(perm, func(perm []int) {
		cost := 0
		for i := 0; i < n-1; i++ {
			cost += distanceMatrix[perm[i]][perm[i+1]]
			fmt.Printf("Visitando cidade %d, custo atual: %d\n", perm[i+1], cost)
		}
		cost += distanceMatrix[perm[n-1]][perm[0]] // Return to starting city
		fmt.Printf("Voltando para a cidade %d, custo total: %d\n", perm[0], cost)

		if cost < minCost {
			minCost = cost
			minRoute = make([]int, n)
			copy(minRoute, perm)
		}
	})

	return minRoute, minCost
}

func permute(a []int, f func([]int)) {
	permuteHelper(a, f, 0)
}

func permuteHelper(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	permuteHelper(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permuteHelper(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
