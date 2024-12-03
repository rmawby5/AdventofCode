import time

startTime = time.time()

f = open('Day1\Input.txt', 'r')
Data = f.readlines()
f.close

ls1 = []
ls2 = []

for i in Data:
    i = i.split()
    #print(i)
    ls1.append(int(i[0]))
    ls2.append(int(i[1]))

ls1.sort()
ls2.sort()

#Part1
TotalDistance = 0
for i in list(range(0, len(ls1))):
    TotalDistance = TotalDistance + abs(ls1[i] - ls2[i])

#Part2
SimilarityScore = 0
for i in ls1:
    s2Occ = ls2.count(i)
    SimilarityScore = SimilarityScore + (s2Occ * i)


elapsed= time.time() - startTime
print("Part 1: ", TotalDistance)
print("Part 2: ", SimilarityScore)



print("Elapsed 2", elapsed)