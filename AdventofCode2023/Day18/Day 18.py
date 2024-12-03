from shapely.geometry import Polygon

f = open('Day18\Input.txt', 'r')
Data = f.readlines()
f.close 

FullCoordslist = [[0, 0]]
FullCoordslistP2 = [[0, 0]]

directionMod = {  #L,R diretion : [x mod, y mod]
    'U' : [0, -1],
    'L' : [-1, 0],
    'D' : [0, 1],
    'R' : [1, 0]
}

directionModP2 = {  #L,R diretion : [x mod, y mod]
    '3' : [0, -1],
    '2' : [-1, 0],
    '1' : [0, 1],
    '0' : [1, 0]
}

def digP1(operation,):
    global FullCoordslist
    coords = [(FullCoordslist[-1][0] + ((int(operation[1])) * directionMod[operation[0]][0])), (FullCoordslist[-1][1] + ((int(operation[1])) * directionMod[operation[0]][1]))]       
    return coords  


def digP2(operation,):
    global FullCoordslistP2
    dist = int(operation[2][2:-2], 16)
    dir = operation[2][-2]
    coords = [(FullCoordslist[-1][0] + (dist * directionModP2[dir][0])), (FullCoordslist[-1][1] + (dist * directionModP2[dir][1]))]       
    return coords 

for i in Data:
    j = i.split()
    lakecord = digP1(j)
    FullCoordslist.append(lakecord)
    Seacord = digP2(j)
    FullCoordslistP2.append(Seacord)

lake = Polygon(FullCoordslist)
P1 = lake.area + (lake.length / 2) + 1
print('Part 1 : ', P1)

Sea = Polygon(FullCoordslistP2)
P2 = Sea.area + (Sea.length / 2) + 1
print('Part 2 : ', P2)
