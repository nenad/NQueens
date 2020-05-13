__author__ = 'Nenad'

from numpy.random import randint, random_sample
import numpy
from operator import add
from functools import reduce

def clamp(minimum, x, maximum):
    return max(minimum, min(x, maximum))

def individual(max_n):
    return randint(0, max_n, max_n)

def population(count, n):
    return [individual(n) for _ in range(0, count)]

def is_valid(q1, q2, debug=False):
    err = 0
    if q1['row'] == q2['row']:
        err += 1
        if debug:
            print("Rows -> ({0}, {1}) ({2}, {3})".format(q1['row'], q1['column'], q2['row'], q2['column']))
    if q1['column'] == q2['column']:
        err += 1
        if debug:
            print("Columns -> ({0}, {1}) ({2}, {3})".format(q1['row'], q1['column'], q2['row'], q2['column']))
    if abs(q2['column'] - q1['column']) == abs(q2['row'] - q1['row']):
        err += 1
        if debug:
            print("Diagonal -> ({0}, {1}) ({2}, {3})".format(q1['row'], q1['column'], q2['row'], q2['column']))
    return err

def draw_table(queen_array):
    n = queen_array.size

    for i in range(0, n):
        line = ""
        for j in range(0, n):
            q_pos = queen_array[i]
            if j == q_pos:
                line += 'X  '
            else:
                line += '_  '
        print(line)



def fitness(queen_array, debug=False):
    if debug:
        print(queen_array)
    fit = 0
    n = queen_array.size
    for i in range(0, n):
        q1 = {'row': i, 'column': queen_array[i]}
        q_err = 0
        for j in range(i+1, n):
            # print "{0}x{1} = ({2}, {3})".format(i,j,queen_array[i],queen_array[j])
            q2 = {'row': j, 'column': queen_array[j]}
            q_err += is_valid(q1, q2, debug)
        fit += q_err
    return fit

def crossover(array_1, array_2):

    n = len(array_1)
    male = array_1[:]
    female = array_2[:]

    child = male
    start = randint(0, n)
    stop = randint(0, n)
    if start > stop:
        start, stop = stop, start

    for i in range(start, stop + 1):
        child = numpy.delete(child, i)
        child = numpy.insert(child, i, female[i])

    #print "Crossover from {0} and {1} to {2} with start, stop ({3}, {4})".format(array_1, array_2, child, start, stop)
    return child

def mutate(queen_array, strength, count=1):
    n = queen_array.size
    #print "Before: {0}".format(queen_array)
    for i in range(0, count):
        selected_index = randint(0, n)
        selected_int = queen_array[selected_index]
        while True:
            new_int = randint(-strength + selected_int, strength+1 + selected_int)
            if new_int != selected_int:
                break

        queen_array[selected_int] = clamp(0, new_int, n - 1)
    #print "After:  {0}".format(queen_array)
    return queen_array

def average_fitness(pop):
    fitnesses = [fitness(ind) for ind in pop]
    return min(fitnesses)
    #summation = reduce(add, fitnesses, 0)
    #return summation * 1.0 / len(pop)

def sort_arrays(pop):
    fit_population = [(fitness(p), p) for p in pop]
    fit_population = sorted(fit_population, key=lambda fit_p: fit_p[0])
    return [p[1] for p in fit_population]

def evolve(pop, random_select=0.05, elite_percent=0.2, mutation=0.03):
    count = len(pop)
    pop = sort_arrays(pop)
    elite = int(count * elite_percent)
    # keep the top 20% of all population
    parents = pop[:elite]
    #print "Top parents: "
    #for parent in parents:
    #    print "{0} ---> {1}".format(parent, fitness(parent))

    # loop through all the rest, pick 5% random
    for ind in pop[elite:]:
        if random_select > random_sample():
            parents.append(ind)

    for ind in parents:
        if mutation > random_sample():
            ind = mutate(ind, len(pop[0]), 2)
            parents.append(ind)

    parents_len = len(parents)
    desired_length = count - parents_len
    children = []
    while len(children) < desired_length:
        parent_1 = randint(0, parents_len)
        parent_2 = randint(0, parents_len)
        if parent_1 != parent_2:
            child = crossover(parents[parent_1], parents[parent_2])
            children.append(child)
            #print "Crossover child: {0} from {1} and {2} ({3} and {4})".format(child, parent_1, parent_2,
                                                                            #   parents[parent_1], parents[parent_2])

    parents.extend(children)
    #print "{0} new children. Total parents: {1}".format(len(children), len(parents))

    #for p in parents:
    #    print "{0} ---> {1}".format(p, fitness(p))

    return parents

pops = population(100, 8)

for x in range(5000):
    pops = evolve(pops, random_select=0.02, elite_percent=0.2, mutation=0.05)
    avg = average_fitness(pops)
    print("Gen: {0}, Fitness: {1}\r".format(x, avg), end="")
    if avg == 0:
        draw_table(pops[0])
        break
