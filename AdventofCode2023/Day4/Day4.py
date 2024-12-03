f = open('Day4\Input.txt', 'r')
Data = f.readlines()
f.close

cardCount = len(Data)
cardsInstances = [1]*cardCount
P1 = 0
P2 = 0

def Part1(arr):
    exponent = 0
    Total = 0
    winning = arr[0].strip().replace('  ', ' ').split(' ')
    ElfNum = arr[1].strip().replace('  ', ' ').split(' ')
    for i in winning:
        if ElfNum.count(i) != 0:
            exponent += 1
    if exponent != 0:
        Total = 2 ** (exponent-1)
    return Total, exponent

def Part2(matchCount, cardNum):
    count = cardsInstances[cardNum]
    for i in range(cardNum + 1, cardNum + 1 + matchCount):
        cardsInstances[i] = cardsInstances[i] + (1*count)
    return None



for i in range(0, cardCount):
    game = Data[i].strip().split(':')[1].split('|')
    score = Part1(game)
    P1 += score[0]
    Part2(score[1], i)

for i in cardsInstances:
    P2 += i

print('Part 1 Total : ', P1 )
print('Part 1 Total : ', P2 )



