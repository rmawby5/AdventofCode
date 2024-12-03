f = open('Day 11\Input.txt', 'r')
Data = f.readlines()
f.close

cdef int expFact
expFact = 1

universe = []

emptyRows = []
emptyColums = []

def pointDistance(pointA, pointb):
    x = abs(pointA[0] - pointb[0])
    y = abs(pointA[1] - pointb[1])
    dist = x + y
    return dist



#create universe and count empty rows

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





disSum = 0
#calculate distances
for i in list(range(len(galaxyCoords) - 1, -1, -1)):
    for j in list(range(i - 1, -1, -1)):
        disSum += pointDistance(galaxyCoords[i], galaxyCoords[j])

        #print(totaldis, expandRow, expandCol)
    galaxyCoords.pop(i)



print(disSum)