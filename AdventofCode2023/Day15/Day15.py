f = open('Day15\Input.txt', 'r')
Data = f.readlines()
f.close

totalP1 = 0
totalP2 = 0

def charValue(cha, curValue):
     newValue = (curValue + cha) * 17
     newValue = newValue % 256
     return newValue
    
def getBox(stg):
    #used for p2:
    elmt = list(stg)
    currentValue = 0
    for i in elmt:
        if i == '=' or i == '-':
            break
        currentValue = charValue(ord(i), currentValue)
    return currentValue

def getValue(stg):
    #used for P1
    elmt = list(stg)
    currentValue = 0
    for i in elmt:
        currentValue = charValue(ord(i), currentValue)
    return currentValue


for i in Data:
    code = i.strip().split(',')

#Part 1
for i in code:
    totalP1 += getValue(i)

#Part 2
boxls = []
for i in range(0, 256):
    boxls.append([])

for i in code:
    box = getBox(i)
    if i[-2] == '=':
        lable = i[:-2]
        if boxls[box].count(lable) != 0:     #Lense of same label is already in box
           oldInd = boxls[box].index(lable)  #get posistion of exsiting lense
           boxls[box].pop(oldInd)            #remove lable
           boxls[box].pop(oldInd)            #remove lense
           boxls[box].insert(oldInd, i[-1])  #add lense then lable to box at same position
           boxls[box].insert(oldInd, lable)           
        else:                               #lens is not in box
            boxls[box].append(lable)        #add lable to back
            boxls[box].append(i[-1])        #add lense to back
    elif i[-1] == '-':                      # removing lense with no replacement
        lable = i[:-1]
        try:
            oldlens = boxls[box].index(lable)   #get position of old lense if present
            boxls[box].pop(oldlens)             #remove old lable, then remove old lense
            boxls[box].pop(oldlens) 
        except:
            None

for i in range(0, len(boxls)):
    lensPos = 1
    if len(boxls[i]) != 0:
        for j in range(1, len(boxls[i]), 2):
            totalP2 += (int(boxls[i][j]) * lensPos * (i + 1))
            lensPos += 1

print('Part 1 Total : ', totalP1)
print('Part 2 Total : ', totalP2)
        