f = open('Day 11\Input.txt', 'r')
Data = f.readlines()
f.close

universe = []

expFact = 2000

def pointDistance(pointA, pointb):
    x = abs(pointA[0] - pointb[0])
    y = abs(pointA[1] - pointb[1])
    dist = x + y
    return dist



#create universe and expand height
rowCount = 0
for i in Data:
    row = list(i)
    if row[len(row) - 1] == '\n':      
        row.pop(len(row) - 1)
    universe.append(row)
    if row.count('#') == 0:
        for r in list(range(0, expFact)):
            universe.append(['.']*len(row))

   
        



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
        for k in list(range(0, len(universe))):
            #print(k, i)
            for r in list(range(0, expFact)):
                universe[k].insert(i, '.')
            
            
            #print(universe[k])
        i += expFact
    #print('row added')
    i += 1

#expand universe height
"""
row = 0   
for i in list(range(0, len(universe))):
    if universe[row].count('#') == 0:
        universe.insert(row, universe[row])
        print('new row')
        row += 1
    row += 1
 """           
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
        dis = pointDistance(galaxyCoords[i], galaxyCoords[j])
        disSum += dis
        #print (dis)
    galaxyCoords.pop(i)

print(disSum)

