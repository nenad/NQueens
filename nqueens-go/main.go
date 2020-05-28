package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Configuration struct {
	Population     int
	MaxGenerations int
	MutationChance float64
	BoardSize      int
}

func loadConfig() Configuration {
	var (
		population     = flag.Int("population", 5000, "Sets the population")
		mutationChance = flag.Float64("mutation", 0.02, "Sets the mutation chance")
		boardSize      = flag.Int("boardSize", 8, "Sets the NxN board size with N queens")
		maxGenerations = flag.Int("maxGenerations", 5000, "Sets the maximum number of generations")
	)

	flag.Parse()

	return Configuration{
		Population:     *population,
		MaxGenerations: *maxGenerations,
		MutationChance: *mutationChance,
		BoardSize:      *boardSize,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := run(loadConfig()); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func run(cfg Configuration) error {
	population := generatePopulation(cfg.Population, cfg.BoardSize)
	calculateFitness(population)

	maxDNA := DNA{}

	for gen := 0; gen <= cfg.MaxGenerations; gen++ {
		// Compare fitness
		totalFitness := float64(0)
		n := len(population)

		for i := 0; i < n; i++ {
			if population[i].Fitness == 1 {
				fmt.Printf("Winner (gen %d): %#v\n", gen, population[i])
				return nil
			}
			if population[i].Fitness > maxDNA.Fitness {
				maxDNA = population[i]
			}
			totalFitness += population[i].Fitness
		}

		newGen := make([]DNA, cfg.Population)
		for i := 0; i < cfg.Population; i++ {
			r1 := rand.Float64() * totalFitness
			r2 := rand.Float64() * totalFitness

			var p1, p2 *DNA

			gotParents := 0
			// best parent selection
			for j := 0; j < n; j++ {
				r1 -= population[j].Fitness
				if r1 <= 0 && p1 == nil {
					p1 = &population[j]
					gotParents++
				}

				r2 -= population[j].Fitness
				if r2 <= 0  && p2 == nil{
					p2 = &population[j]
					gotParents++
				}

				if gotParents == 2 {
					break
				}
			}

			child := p1.Crossover(*p2)
			child.Mutate(cfg.MutationChance)
			child.CalculateFitness()
			newGen[i] = child
		}

		population = newGen

		fmt.Printf("Gen: %d, Max fitness now: %.2f\r", gen, maxDNA.Fitness)
	}

	fmt.Printf("Best candidate: %#v\n", maxDNA)

	return nil
}
