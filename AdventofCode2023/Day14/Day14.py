import numpy
import copy

f = open('Day14\InputTest.txt', 'r')
Data = f.readlines()
f.close

def northShift(grid):
    gridn = numpy.transpose(grid)
    lastInd = 0
    for i in gridn:
        if i[0] == '#' or i[0] == 'O':
            lastInd = 1
        for j in range(1, len(i)):
            if (i[j-1] == '#' or i[j-1] == 'O') and i[j] == '.':
                lastInd = j
            if i[j] == 'O':
                i[j] = '.'
                i[lastInd] = 'O'
                lastInd += 1
            if i[j] == '#':
                lastInd = j + 1
        lastInd = 0
    newgrid = numpy.transpose(gridn)
    return newgrid

def eastShift(grid):
    gride = []
    for i in grid:
        gride.append(copy.deepcopy(i.tolist()))
    lastInd = 0
    for i in gride:
        i.reverse()
        if i[0] == '#' or i[0] == 'O':
            lastInd = 1
        for j in range(1, len(i)):
            if (i[j-1] == '#' or i[j-1] == 'O') and i[j] == '.':
                lastInd = j
            if i[j] == 'O':
                i[j] = '.'
                i[lastInd] = 'O'
                lastInd += 1
            if i[j] == '#':
                lastInd = j + 1
        i.reverse()
        lastInd = 0
    return gride   

def westShift(grid):
    lastInd = 0
    gridw = copy.deepcopy(grid)
    for i in gridw:
        if i[0] == '#' or i[0] == 'O':
            lastInd = 1
        for j in range(1, len(i)):
            if (i[j-1] == '#' or i[j-1] == 'O') and i[j] == '.':
                lastInd = j
            if i[j] == 'O':
                i[j] = '.'
                i[lastInd] = 'O'
                lastInd += 1
            if i[j] == '#':
                lastInd = j + 1
        lastInd = 0
    return gridw


def southShift(grids):
    grids = numpy.transpose(numpy.flip(grids))
    lastInd = 0
    for i in grids:
        if i[0] == '#' or i[0] == 'O':
            lastInd = 1
        for j in range(1, len(i)):
            if (i[j-1] == '#' or i[j-1] == 'O') and i[j] == '.':
                lastInd = j
            if i[j] == 'O':
                i[j] = '.'
                i[lastInd] = 'O'
                lastInd += 1
            if i[j] == '#':
                lastInd = j + 1
        lastInd = 0
    newgrid = numpy.transpose(grids)
    newgrid = numpy.flip(newgrid)
    return newgrid

def cycle(grid):
    gridn = northShift(grid)
    gridw = westShift(gridn)
    grids = southShift(gridw)
    gride = eastShift(grids)
    return gride

def loadCount(grid):
    loadTotal = 0
    for row in range(0, len(grid)):
        loadDistance = len(grid) - row
        for i in grid[row]:
            if i == 'O':
                loadTotal += loadDistance
    return loadTotal

def part1(grid):
    grid = northShift(grid)
    total = loadCount(grid)
    return total

def part2(grid):
    c = 0
    results = []
    while c < 1000:
        #print('new cycle')
        grid = cycle(grid)
        c += 1
        j = loadCount(grid)
        results.append(j)
        #print(c, ' : ', j)

    return results

platform = []

for i in Data:
    platform.append(list(i.strip()))

print('Part 1 Total: ', part1(platform))
print('Part 2 Total: ', part2(platform))







