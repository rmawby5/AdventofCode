f = open('Day17\InputTest.txt', 'r')
Data = f.readlines()
f.close 

Map = []
mapDict = {}

for i in Data:
    Map.append(list(i).pop)

for i in range(0, len(Map)):
    print(Map[i])
    for i in range(0, len(Map[i])):
        dict