#stupid, sorta
def allzero(list):
    if list == []:
        return False
    for i in list:
        if i != 0:
            return False
    return True

f = open("Day 9\Input.txt", 'r')
file = f.read().strip().split('\n')
p1sum = 0
p2sum = 0
for i in file:
    nums = [int(x) for x in i.split()]
    difflist = [nums]
    diffs = []
    newdiffs = []
    for j in range(1, len(nums)):
        newdiffs.append(nums[j] - nums[j-1])
    diffs = newdiffs
    difflist.append(diffs)
    while not allzero(newdiffs):
        newdiffs = []
        for j in range(1, len(diffs)):
            newdiffs.append(diffs[j] - diffs[j-1])
        diffs = newdiffs
        difflist.append(diffs)
    nextnum = 0
    p2nextnum = 0
    for j in enumerate(difflist):
        if j[0] == len(difflist):
            break
        nextnum += j[1][-1]
    for j in enumerate(difflist[::-1]):
        p2nextnum = j[1][0] - p2nextnum
    #print(nextnum)
    p1sum += nextnum
    p2sum += p2nextnum
print(p1sum)
print(p2sum)