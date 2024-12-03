import itertools, copy
from functools import cache

f = open('Day 12\Input.txt', 'r')
Data = f.readlines()
f.close

permcount = 0

@cache
def checkCompliance(patten, gol):
    pattern = list(patten)
    goal = gol.split(',')
    blockcount = []
    curBlock = 0
    try:
        ind = pattern.index('?')
    except:
        ind = len(pattern) - 1
    #print(pattern, ' : ', goal)
    for i in range(0, ind + 1):
        if pattern[i] == '#':
            curBlock += 1
            #print('+1')
            if i == ind:                         #reached end of pattern, add current block
                blockcount.append(curBlock)
                curBlock = 0
        if (pattern[i] == '.' or pattern[i] == '?') and i!= 0:
            if pattern[i-1] == '#':
                blockcount.append(curBlock)
            curBlock = 0
    #print(blockcount)
    if len(blockcount) != 0:
        for i in range(0, len(blockcount)-1):
            if blockcount[i] != int(goal[i]):
                #print('Failed check')
                return False             
        if blockcount[-1] > int(goal[len(blockcount) - 1]):
           # print('Failed check')
            return False             
    return True


def CalcCombinations(pattern, goal, goaltotal):
    global permcount
    ptn = copy.deepcopy(pattern)                                             
    if ptn.count('?') == 0:                             #No '?' in current pattern, check if combination is valid and add to total perm count
        if checkCompliance(ptn, goal) == True:
            permcount += 1
        return None    
    ind = ptn.index('?')                         #get position of first instance of '?'
    if ptn.count('#') < goaltotal:                 
        #ptn[ind] = '#'                          #try with '#' in place of ?'  
        ptnNew = ptn[:ind] + '#' + ptn[ind + 1:]
        if checkCompliance(ptnNew, goal) == True:              #if combination so far is valid after change, continue with next step.
            CalcCombinations(ptnNew, goal, goaltotal)
    if ptn.count('#') + ptn.count('?') > goaltotal:
        #ptn[ind] = '.'
        ptnNew = ptn[:ind] + '.' + ptn[ind + 1:]
        if checkCompliance(ptnNew, goal) == True:              #if combination so far is valid after change, continue with next step.
            CalcCombinations(ptnNew, goal, goaltotal)
    return None


def springTotal(goal):
    tot = 0
    for i in goal:
        if i != ',':
            tot += int(i)
    return tot

"""
pat = '???.###'
gol = '1,1,3'

springs = springTotal(gol)
CalcCombinations(pat, gol, springs)

print(permcount)

"""

count = 1
for i in Data:
    p = i.strip().split(' ')
    pat = p[0]
    num = p[1]

    #extra processing for part 2
    pat += '?'
    pat = (pat*5)[:-1]
    num += ','
    num = (num*5)[:-1]
    springs = springTotal(num)
    CalcCombinations(pat, num, springs)
    print('Done :', count)
    count += 1


print('Part Total : ', permcount)
print('Done')
