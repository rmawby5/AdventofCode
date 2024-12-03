import more_itertools

f = open('Day 12\InputTest.txt', 'r')
Data = f.readlines()
f.close

def getPerms(charlist):
    perms = more_itertools.distinct_combinations(charlist, len(charlist))
    return perms

totalcombinations = 0

for o in Data:
    p = o.split(' ')
    j = list(p[0])
    matchCount = 0
    num = list(p[1])
    if num[-1] == '\n':
        num.pop(-1)   
    numStr = ''.join(num)
    num = numStr.split(',')
    total = 0
    for i in list(range(0, len(num), 1)):
        total += int(num[i])
    k = j.count('?')
    h = j.count('#')
    dot = k - (total - h)
    l = []
    for i in list(range(0, dot)):
        l.append('.')
    for i in list(range(0, (total - h))):
        l.append('#')
    #combinations = getPerms(l)
    combinations = more_itertools.distinct_permutations(l, len(l))
    count = 0
    for y in combinations:
        count += 1
    print(count)
    print('stop')
    for item in combinations:
        ind = 0
        seq = []
        for n in j:
            seq.append(n)        
        for m in list(range(0, len(seq))):
            if seq[m] == '?':
               seq[m] = item[ind]
               ind += 1      
        blockCount = 0
        stringMatch = []
        for i in seq:
            if i == '#':
                blockCount += 1
            elif i != '#' and blockCount != 0:
                stringMatch.append(str(blockCount))
                blockCount = 0
        if blockCount != 0:        
            stringMatch.append(str(blockCount))
        if stringMatch == num:
            matchCount += 1       
    totalcombinations += matchCount
print(totalcombinations)
