package main

import (
	"math/rand"
)

type DNA struct {
	Positions []int
	Fitness   float64
}

func (dna *DNA) CalculateFitness() {
	n := len(dna.Positions)
	collisions := float64(0)

	for i, v := range dna.Positions {

		q := i + v
		// Right side
		for fw := i + 1; fw < n; fw++ {
			vv := dna.Positions[fw]
			fwq := fw + vv
			if q == fwq {
				collisions++
				continue
			}

			if i-fw == v-vv {
				collisions++
				continue
			}

			if v == vv {
				collisions++
				continue
			}
		}

		// Left side
		for bw := i - 1; bw >= 0; bw-- {
			vv := dna.Positions[bw]
			fwq := bw + dna.Positions[bw]
			if q == fwq {
				collisions++
				continue
			}

			if i-bw == v-vv {
				collisions++
				continue
			}

			if v == vv {
				collisions++
				continue
			}
		}
	}

	max := float64(n * (n - 1))

	dna.Fitness = float64(1) - collisions/max
}

func (dna *DNA) Crossover(dna2 DNA) DNA {
	child := DNA{
		Positions: make([]int, len(dna.Positions)),
	}

	for i := 0; i < len(dna.Positions); i++ {
		if rand.Float64() < 0.5 {
			child.Positions[i] = dna.Positions[i]
		} else {
			child.Positions[i] = dna2.Positions[i]
		}
	}

	return child
}

func (dna *DNA) Mutate(chance float64) {
	for i := 0; i < len(dna.Positions); i++ {
		if rand.Float64() < chance {
			dna.Positions[i] = rand.Intn(len(dna.Positions))
		}
	}
}

func randomDNA(size int) DNA {
	dna := DNA{
		Positions: make([]int, size),
	}

	for i := 0; i < size; i++ {
		dna.Positions[i] = rand.Intn(size)
	}
	return dna
}

func generatePopulation(n int, size int) []DNA {
	population := make([]DNA, n)
	for i := 0; i < n; i++ {
		population[i] = randomDNA(size)
	}

	return population
}

func calculateFitness(population []DNA) {
	for i, _ := range population {
		population[i].CalculateFitness()
	}
}
