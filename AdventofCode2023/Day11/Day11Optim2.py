f = open('InputTest.txt', 'r')
Data = f.readlines()
f.close

expFact = 1
disSum = 0
universe = []
emptyRows = []
emptyColums = []
galaxyCoords = []

#create universe map and count empty rows
rowCo = 0
for i in Data:
    row = list(i)
    if row[len(row) - 1] == '\n':      
        row.pop(len(row) - 1)
    universe.append(row)
    if row.count('#') == 0:
        emptyRows.append(rowCo)
    rowCo += 1
#count empty columns
i = 0
for l in list(range(0, len(universe[0]))):
    galaxyCount = 0
    for j in list(range(0, len(universe))):
        if universe[j][i] == '#':
            galaxyCount += 1
    if galaxyCount == 0: #add new column
        emptyColums.append(i)
    i += 1
for i in list(range(0, len(universe))):
    for j in list(range(0, len(universe[i]))):
        if universe[i][j] == '#':
            galaxyCoords.append([j, i])
#count empty rows/columns before galaxy
for i in galaxyCoords:
    emtR = 0
    emtC = 0
    for j in emptyColums:
        if j < i[0]:
            emtC += 1
    for k in emptyRows:
        if k < i[1]:
            emtR += 1
    i[0] = i[0] + (emtC * expFact)
    i[1] = i[1] + (emtR * expFact)
#calculate distances
for i in list(range(len(galaxyCoords) - 1, -1, -1)):
    for j in list(range(i - 1, -1, -1)):
        disSum += (abs(galaxyCoords[i][0] - galaxyCoords[j][0]) + abs(galaxyCoords[i][1] - galaxyCoords[j][1]))
    galaxyCoords.pop(i)
print(disSum)