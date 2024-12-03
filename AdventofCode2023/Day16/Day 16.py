import sys, threading
from functools import cache

sys.setrecursionlimit(10000)


f = open('Day16\Input.txt', 'r')
Data = f.readlines()
f.close 

startX = 4
startY = 0

PipeMatrice = []

tileTotal = 0
StepCounter = 0


for i in Data:
    PipeMatrice.append(list(i))

for i in PipeMatrice:                       #remove newline char, add wall condition
   i.pop(len(i) - 1)
   i.insert(0, '*')
   i.append('*')

PipeMatrice.insert(0, ['*']*(len(PipeMatrice[0])))
PipeMatrice.append(['*']*(len(PipeMatrice[0])))


def rayTrace(curX, curY, prevX, prevY):

    while PipeMatrice[curY][curX] != '*':
        global StepCounter
        StepCounter += 1
        if PipeMatrice[curY][curX] == '|':
            if prevX != curX:     #hits mirror
                rayTrace(curX, curY - 1, curX, curY)

                prevX = curX
                prevY = curY
                curY = curY + 1          
            else:                           #inline with mirror
                dif = curY - prevY
                prevY = curY
                prevX = curX
                curY = curY + dif
        elif PipeMatrice[curY][curX] == '-':
            if prevY != curY:     #hits mirror
                rayTrace(curX - 1, curY, curX, curY)
                prevX = curX
                prevY = curY
                curX = curX + 1          
            else:                       #inline with mirror
                dif = curX - prevX
                prevX = curX
                prevY = curY
                curX = curX + dif
        elif PipeMatrice[curY][curX] == '\\':
            if prevY != curY:     #hits mirror
                dif = curY - prevY
                prevX = curX
                prevY = curY              
                curX = curX + dif          
            else:                       #inline with mirror
                dif = curX - prevX
                prevX = curX
                prevY = curY
                curY = curY + dif
        elif PipeMatrice[curY][curX] == '/':
            if prevY != curY:     #hits mirror
                dif = curY - prevY
                prevX = curX
                prevY = curY              
                curX = curX - dif          
            else:                       #inline with mirror
                dif = curX - prevX
                prevX = curX
                prevY = curY
                curY = curY - dif
        elif PipeMatrice[curY][curX] == '.':
            BlankMatrice[curY][curX] = StepCounter
            BlankMatrice2[curY][curX] = '#'
            StepCounter += 1
            difX = curX - prevX
            difY = curY - prevY
            prevY = curY
            prevX = curX
            curX = curX + difX
            curY = curY + difY       
            while PipeMatrice[curY][curX] == '.':
                BlankMatrice[curY][curX] = StepCounter
                BlankMatrice2[curY][curX] = '#'
                prevX = curX
                prevY = curY
                curX = curX + difX
                curY = curY + difY

                StepCounter += 1

        if BlankMatrice[curY][curX] != 0:
            if BlankMatrice[curY][curX] - BlankMatrice[prevY][prevX] == 1:         #path has already been passed, exit loop
                return None                                        
            else:
                BlankMatrice[prevY][prevX] = StepCounter
                BlankMatrice2[prevY][prevX] = '#'
        else:
            BlankMatrice[prevY][prevX] = StepCounter
            BlankMatrice2[prevY][prevX] = '#'
    return None

BlankMatrice = []
BlankMatrice2 = []

#sum non-zero elements in results

def Part1(starX, starY):
    tileTotal = 0
    if starY == 0:
        currY = 1
        currX = starX
    elif starX == 0:
        currY = starY
        currX = 1
    elif starY == 11:
        currY = 10
        currX = starX
    elif starX == 11:
        currY = starY
        currX = 10
    global BlankMatrice
    BlankMatrice = [[0]*(len(PipeMatrice[0])) for i in range(len(PipeMatrice))]
    global BlankMatrice2
    BlankMatrice2 = [['.']*(len(PipeMatrice[0])) for i in range(len(PipeMatrice))]
    global StepCounter
    StepCounter = 0
    rayTrace(currX, currY, starX, starY)
    for i in range(1, len(BlankMatrice2) - 1):
        l = len(BlankMatrice2[i]) - BlankMatrice2[i].count('.')
        if BlankMatrice2[i][0] != '.':
            l -= 1
        if BlankMatrice2[i][-1] != '.':
            l -= 1
        tileTotal += l
    return tileTotal
    
print('Part 1 Total : ', Part1(0, 1))

currentMax = 0


def Toprow():
    global BlankMatrice
    global BlankMatrice2
    global currentMax
    for i in range(1, len(PipeMatrice) - 1):
        BlankMatrice = []
        BlankMatrice2 = []
        tot = Part1(i, 0)
        if tot > currentMax:
            currentMax = tot

def bottomrow():
    global BlankMatrice
    global BlankMatrice2
    global currentMax
    for i in range(1, len(PipeMatrice) - 1):
        BlankMatrice = []
        BlankMatrice2 = []
        tot = Part1(0, i)
        if tot > currentMax:
            currentMax = tot


def leftrow():
    global BlankMatrice
    global BlankMatrice2
    global currentMax
    for i in range(1, len(PipeMatrice) - 1):
        BlankMatrice = []
        BlankMatrice2 = []
        tot = Part1(i, 11)
        if tot > currentMax:
            currentMax = tot



def rightrow():
    global BlankMatrice
    global BlankMatrice2
    global currentMax
    for i in range(1, len(PipeMatrice) - 1):
        BlankMatrice = []
        BlankMatrice2 = []
        tot = Part1(11, i)

        if tot > currentMax:
            currentMax = tot


print('Part 2 Total : ' )


t1 = threading.Thread(target=Toprow)
t2 = threading.Thread(target=bottomrow)


print(t1.is_alive())

t1.start()
t2.start()


print(t1.is_alive())

t1.join()
t2.join()


print('done')
print (currentMax)

