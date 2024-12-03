from shapely.geometry import Point, Polygon
import matplotlib.path as mlptpath


f = open('Day10\Input.txt', 'r')
Data = f.readlines()
f.close 

startX = 0
startY = 0

nextX = 0
nextY = 0

StepCounter = 0

PipeMatrice = []
rowCount = 0
for i in Data:
    PipeMatrice.append(list(i))
    if i.count('s') == 1:
        startY = rowCount
        startX = i.index('s')
    elif i.count('S') == 1:
        startY = rowCount
        startX = i.index('S')
    rowCount += 1




"""
print('x: ')
print(startX)

print('y: ')
print(startY)
"""

for i in PipeMatrice:
   i.pop(len(i) - 1)


#for i in list(range(0, len(PipeMatrice) - 1)):
 #   print (PipeMatrice[i])


BlankMatrice = [[0]*(len(PipeMatrice[0])) for i in range(len(PipeMatrice))]


if PipeMatrice[startY][startX] != 's' and PipeMatrice[startY][startX] != 'S':
    print ('incorrect start point')
else:
    print ('start located')




#get initial direction

if PipeMatrice[startY + 1][startX] == '|' or PipeMatrice[startY + 1][startX] == 'J' or PipeMatrice[startY + 1][startX] == 'L':
    nextX = startX
    nextY = startY + 1
    #print('successful start')
    StepCounter += 1

elif PipeMatrice[startY - 1][startX] == '|' or PipeMatrice[startY + 1][startX] == 'F' or PipeMatrice[startY + 1][startX] == '7':
    nextX = startX
    nextY = startY - 1
    #print('successful start')
    StepCounter += 1
    

elif PipeMatrice[startY][startX + 1] == '-' or PipeMatrice[startY + 1][startX] == 'J' or PipeMatrice[startY + 1][startX] == 'J':
    nextX = startX + 1
    nextY = startY
    #print('successful start')
    StepCounter += 1

elif PipeMatrice[startY][startX - 1] == '-' or PipeMatrice[startY + 1][startX] == 'F' or PipeMatrice[startY + 1][startX] == 'L':
    nextX = startX - 1
    nextY = startY
    #print('successful start')
    StepCounter += 1
else:
    print('no path found!')

currentX = startX
currentY = startY


fullCord = []


while PipeMatrice[nextY][nextX] != 's' and PipeMatrice[nextY][nextX] != 'S':
    #print(PipeMatrice[nextY][nextX], nextY, nextX, currentY, currentX)
    currentCord = [currentX, currentY]

    fullCord.append(currentCord)

    BlankMatrice[nextY][nextX] = StepCounter
    if PipeMatrice[nextY][nextX] == '|':
        if nextY < currentY:
            currentY = nextY
            nextY = nextY - 1
        else:
            currentY = nextY
            nextY = nextY + 1
        StepCounter += 1


    elif PipeMatrice[nextY][nextX] == '-':
        if nextX < currentX:
            currentX = nextX
            nextX = nextX - 1
        else:
            currentX = nextX
            nextX = nextX + 1
        StepCounter += 1


    elif PipeMatrice[nextY][nextX] == 'L':
        if nextY > currentY:
            currentX = nextX
            currentY = nextY
            nextX = nextX + 1
        else:
            currentY = nextY
            currentX = nextX
            nextY = nextY - 1
        StepCounter += 1


    elif PipeMatrice[nextY][nextX] == 'J':
        if nextY > currentY:
            currentX = nextX
            currentY = nextY
            nextX = nextX - 1
        else:
            currentY = nextY
            currentX = nextX
            nextY = nextY - 1
        StepCounter += 1


    elif PipeMatrice[nextY][nextX] == '7':
        if nextY < currentY:
            currentX = nextX
            currentY = nextY
            nextX = nextX - 1
        else:
            currentY = nextY
            currentX = nextX
            nextY = nextY + 1
        StepCounter += 1



    elif PipeMatrice[nextY][nextX] == 'F':
        if nextY < currentY:
            currentX = nextX
            currentY = nextY
            nextX = nextX + 1
        else:
            currentY = nextY
            currentX = nextX
            nextY = nextY + 1
        StepCounter += 1



#if PipeMatrice[nextY][nextX] == PipeMatrice[startY][startX]:
    #print('Path found!')

def isEven(num):
    if (num % 2) == 0:
        rtn = True
    else:
        rtn = False
    return rtn
    


if (StepCounter % 2) == 0:
    path = StepCounter // 2
else:
    path = (StepCounter // 2) + 1

#print('path length: ')
print('Steps: ', path)





#replace S character
if PipeMatrice[startY][startX] == 'S':
    if BlankMatrice[startY+1][startX] == 1 or BlankMatrice[startY-1][startX] == 1:
        
        if BlankMatrice[startY+1][startX] + BlankMatrice[startY-1][startX] == StepCounter:
            PipeMatrice[startY][startX] = '|'

    if BlankMatrice[startY+1][startX] == 1 or BlankMatrice[startY][startX-1] == 1:
        
        if BlankMatrice[startY+1][startX] + BlankMatrice[startY][startX-1] == StepCounter:
            PipeMatrice[startY][startX] = '7'

    if BlankMatrice[startY+1][startX] == 1 or BlankMatrice[startY][startX+1] == 1:
        
        if BlankMatrice[startY+1][startX] + BlankMatrice[startY][startX+1] == StepCounter:
            PipeMatrice[startY][startX] = 'F'

    if BlankMatrice[startY][startX-1] == 1 or BlankMatrice[startY][startX+1] == 1:
        
        if BlankMatrice[startY][startX-1] + BlankMatrice[startY][startX+1] == StepCounter:
            PipeMatrice[startY][startX] = '-'

    if BlankMatrice[startY][startX-1] == 1 or BlankMatrice[startY-1][startX] == 1:
        
        if BlankMatrice[startY][startX-1] + BlankMatrice[startY-1][startX] == StepCounter:
            PipeMatrice[startY][startX] = 'J'

    if BlankMatrice[startY-1][startX] == 1 or BlankMatrice[startY][startX-1] == 1:
        
        if BlankMatrice[startY-1][startX] + BlankMatrice[startY][startX-1] == StepCounter:
            PipeMatrice[startY][startX] = 'L'




tileTotal = 0




for i in list(range(1, len(BlankMatrice) - 1)):
    lineTotal = 0
    verLine = 0
    k = BlankMatrice[i]
    lastpipe = 's'
    for j in (list(range(0, len(k)))):
        if BlankMatrice[i][j] - BlankMatrice[i-1][j] == -1 or BlankMatrice[i + 1][j] - BlankMatrice[i][j] == -1:
            #print('u')
            #line is going up
            #check if line is continuous with previous row
        
            if PipeMatrice[i][j] == 'J' and lastpipe != 'F':
                #if lineTotal != 1:
                lineTotal += 1
                verLine += 1
            elif PipeMatrice[i][j] == '7' and lastpipe != 'L':
                lineTotal += 1
                verLine += 1    
            elif PipeMatrice[i][j] == '|'  or PipeMatrice[i][j] == 'L' or PipeMatrice[i][j] == 'F':
                lineTotal += 1
                verLine += 1   
            lastpipe = PipeMatrice[i][j]


        elif BlankMatrice[i][j] - BlankMatrice[i-1][j] == 1 or BlankMatrice[i + 1][j] - BlankMatrice[i][j] == 1:
            #line is going down
            #check if line is continuous with previous row
            #print('d')
            if PipeMatrice[i][j] == 'J' and lastpipe != 'F':
                
                #if lineTotal != 1:
                lineTotal -= 1
                verLine += 1
            elif PipeMatrice[i][j] == '7' and lastpipe != 'L':
                lineTotal -= 1
                verLine += 1
            
            elif PipeMatrice[i][j] == '|' or PipeMatrice[i][j] == 'L' or PipeMatrice[i][j] == 'F':
                lineTotal -= 1
                verLine += 1  

            lastpipe = PipeMatrice[i][j]
        elif BlankMatrice[i][j] == 0 and lineTotal != 0:
            tileTotal += 1
        #print (lineTotal)
    #print(k, lineTotal, tileTotal)
    

#for i in list(range(0, len(BlankMatrice))):
#  print (BlankMatrice[i])


#for i in fullCord:
#    print (i)





print (len(fullCord))
tileTotal2 = 0

print('Tiles : ', tileTotal)


loop = Polygon(fullCord)
for i in list(range(0, len(BlankMatrice))):
    for j in list(range(0, len(BlankMatrice[i]))):
        tile = Point(j,i)
        
        if tile.within(loop) == True:
            tileTotal2 += 1
            print(tile, 'true')
"""
            

loop = mlptpath.Path(fullCord)
for i in list(range(0, len(BlankMatrice))):
    for j in list(range(0, len(BlankMatrice[i]))):

        
        if loop.contains_points([(j, i)]) == True:
            tileTotal2 += 1
            #print(tile, 'true')
"""

print('Tiles from poly: ', tileTotal2)



