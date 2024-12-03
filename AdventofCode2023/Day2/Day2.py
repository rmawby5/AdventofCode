f = open('Day2\Input.txt', 'r')
Data = f.readlines()
f.close

gameCount = len(Data)

P1 = 0
P2 = 0

cubeMax = {
    'blue': 14,
    'red': 12,
    'green': 13
}

def Part1(game):
    for i in game:
        pull = i.split(',')
        #print(pull)
        for j in pull:
            cubes = j.split(' ')
            #print(cubes)
            if cubeMax[cubes[2]] < int(cubes[1]):
                return False

    return True

def Part2(game):
    r_min = 0
    b_min = 0
    g_min = 0
    for i in game:
        pull = i.split(',')
        #print(pull)
        for j in pull:
            cubes = j.split(' ')
            #print(cubes)
            if cubes[2] == 'blue' and b_min < int(cubes[1]):
                b_min = int(cubes[1])
            if cubes[2] == 'red' and r_min < int(cubes[1]):
                r_min = int(cubes[1])
            if cubes[2] == 'green' and g_min < int(cubes[1]):
                g_min = int(cubes[1])
    return (r_min*b_min*g_min)

for i in range(0, gameCount):
    game = Data[i].strip().split(':')[1].split(';')
    if Part1(game) == True:
        P1 += (i + 1) 
    P2 += Part2(game)
    #print(cardsInstances)

print('Part 1 Total : ', P1 )
print('Part 1 Total : ', P2 )
