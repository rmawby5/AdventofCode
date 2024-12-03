f = open('Day 9\Input.txt', 'r')
Data = f.readlines()
f.close 

def zeroElements(inputList):
    if len(inputList) == inputList.count(0):
        Total = 0
    else:
        Total = 1
    return Total

def difference(TopSeq):
    LowSeq = []
    lastElmt = 0
    for i in list(range(0, len(TopSeq)-1)):
        LowSeq.append(TopSeq[i+1] - TopSeq[i])
    if zeroElements(LowSeq) != 0:
        lastElmt = difference(LowSeq)
        lastElmt = lastElmt + TopSeq[len(TopSeq) - 1]
    elif zeroElements(LowSeq) == 0:
        lastElmt = TopSeq[len(TopSeq) - 1]
    return lastElmt

def differenceReverse(TopSeq):
    LowSeq = []
    firstElmt = 0
    for i in list(range(0, len(TopSeq)-1)):
        LowSeq.append(TopSeq[i+1] - TopSeq[i])
    if zeroElements(LowSeq) != 0:
        firstElmt = differenceReverse(LowSeq)
        firstElmt = TopSeq[0] - firstElmt
    elif zeroElements(LowSeq) == 0:
        firstElmt = TopSeq[0]
    return firstElmt
    
SequenceTotalP1 = 0
SequenceTotalP2 = 0
for i in Data:
    Seq = []
    i = i.split()
    for j in i:
        Seq.append(int(j))
    SequenceTotalP1 = SequenceTotalP1 + difference(Seq)
    SequenceTotalP2 = SequenceTotalP2 + differenceReverse(Seq)

print (SequenceTotalP1)
print(SequenceTotalP2)