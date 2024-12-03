import math
f = open('Input.txt', 'r')
Data = f.readlines()
f.close

path = open('Path.txt')
Seq = path.readline()
path.close
SeqL = list(Seq)

LeftDict = {}
RightDict = {}
StartingNodes = []

for i in Data:
    i = i.replace(' ', '')
    if i[2] == 'A':
        StartingNodes.append(i[:3])
    
    LeftDict[i[:3]] = i[5:8]
    RightDict[i[:3]] = i[9:12]


TotalSteps = []
for nextStr in StartingNodes:
    Steps = 0
    while nextStr[2] != 'Z':
        for i in SeqL:

            if i == 'L':
                nextStr = LeftDict[nextStr]
            elif i == 'R':
                nextStr = RightDict[nextStr]
            Steps = Steps + 1
        
            if nextStr[2] == 'Z':
                
                break


    TotalSteps.append(Steps)


#Total = math.lcm(TotalSteps[0], TotalSteps[1], TotalSteps[2], TotalSteps[3], TotalSteps[4], TotalSteps[5])
#print(Total)