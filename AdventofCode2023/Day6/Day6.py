from math import ceil
f = open('Day6\Input.txt', 'r')
Data = f.readlines()
f.close

P1 = 1
P2 = 0

Times = Data[0].strip().split()
Times.pop(0)
Distance = Data[1].strip().split()
Distance.pop(0)


def minMaxTime(TotalTime, Distance):
    max = (TotalTime + (((TotalTime ** 2) - (4 * Distance)) ** 0.5)) / 2
    min = (TotalTime - (((TotalTime ** 2) - (4 * Distance)) ** 0.5)) / 2
    if max == int(max):
        max -= 1
    else:
        max = int(max)
    if min == int(min):
        min += 1
    else:
        min = ceil(min)
    return max, min


def Part1(time, dis):
    max, min = minMaxTime(time, dis)
    Options = max - min + 1
    #print(min, max, Options)
    return Options



for t, d in zip(Times, Distance):
   P1 = P1 * Part1(int(t), int(d))




print('Part 1 Total : ', P1 )
print('Part 2 Total : ', Part1(60947882, 475213810151650) )
