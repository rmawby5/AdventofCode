f = open('Day 11\Input.txt', 'r')
Data = f.readlines()
f.close

expFact = 1

universe = []

emptyRows = []
emptyColums = []

def pointDistance(pointA, pointb):
    x = abs(pointA[0] - pointb[0])
    y = abs(pointA[1] - pointb[1])
    dist = x + y
    return dist


def blankColumnCount(pointA, pointb):
    numRow = 0
    if pointA[0] > pointb[0]:
        for i in emptyColums:
            if i > pointb[0] and i < pointA[0]:
                numRow += 1
        #print ('abig')
    elif pointA[0] < pointb[0]:
        for i in emptyColums:
            if i < pointb[0] and i > pointA[0]:
                numRow += 1
        #print('bbig')
    #print(numRow)
    return numRow



def blankrowCount(pointA, pointb):
    numRow = 0
    if pointA[1] > pointb[1]:
        for i in emptyRows:
            if i > pointb[1] and i < pointA[1]:
                numRow += 1
    elif pointA[1] < pointb[1]:
        for i in emptyRows:
            if i < pointb[1] and i > pointA[1]:
                numRow += 1
    return numRow

#create universe and expand height

rowCo = 0
for i in Data:
    row = list(i)
    if row[len(row) - 1] == '\n':      
        row.pop(len(row) - 1)
    universe.append(row)
    if row.count('#') == 0:
        emptyRows.append(rowCo)
    rowCo += 1

        



i = 0

#expand universe width
for l in list(range(0, len(universe[0]))):
    galaxyCount = 0
    for j in list(range(0, len(universe))):
        #print (universe[j][i])
        if universe[j][i] == '#':
            galaxyCount += 1
    #print('Adding new row')
    if galaxyCount == 0: #add new column
            #print(k, i)
                emptyColums.append(i)
            
            
            #print(universe[k])
        
    #print('row added')
    i += 1


     
#show universe
#for i in list(range(0, len(universe))):
#  print (universe[i])

#get coords of each galaxy

galaxyCoords = []

for i in list(range(0, len(universe))):
    for j in list(range(0, len(universe[i]))):
        if universe[i][j] == '#':
            galaxyCoords.append([j, i])

#print (galaxyCoords)


disSum = 0
#calculate distances
for i in list(range(len(galaxyCoords) - 1, -1, -1)):
    for j in list(range(i - 1, -1, -1)):
        pointdis = pointDistance(galaxyCoords[i], galaxyCoords[j])
        expandRow = blankrowCount(galaxyCoords[i], galaxyCoords[j]) * (expFact)
        expandCol = blankColumnCount(galaxyCoords[i], galaxyCoords[j]) * (expFact)
        totaldis = pointdis + expandCol + expandRow
        disSum = disSum + totaldis
        #print(totaldis, expandRow, expandCol)
    galaxyCoords.pop(i)



print(disSum)