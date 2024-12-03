import numpy
f = open('Day7\Input.txt', 'r')
Data = f.readlines()
f.close

P1 = 0
P2 = 0

def getType(hand):
    c1count = hand.count(hand[0])
    c2count = hand.count(hand[1])
    c3count = hand.count(hand[2])
    c4count = hand.count(hand[3])
    c5count = hand.count(hand[4])
    countls = [c1count, c2count, c3count, c4count, c5count]
    if c1count == 5:    #five of a kind
        HandType = 6
    elif c1count == 4 or c2count == 4:    #four of a kind, 1st or 2nd must match
        HandType = 5
    elif (c1count == 3 or c2count == 3 or c3count == 3) and (c1count == 2 or c2count == 2 or c3count == 2 or c4count == 2):    # full house
        HandType = 4
    elif countls.count(3) == 3:    #three of a kind, 1st or 2nd or 3rd must match
        HandType = 3
    elif countls.count(2) == 4:   #two pair
        HandType = 2
    elif countls.count(2) == 2:     #one pair
        HandType = 1
    elif countls.count(1) == 5:
        HandType = 0   
    return HandType

def getTypeP2(hand):
    c1count = hand.count(hand[0])
    c2count = hand.count(hand[1])
    c3count = hand.count(hand[2])
    c4count = hand.count(hand[3])
    c5count = hand.count(hand[4])
    jcount = hand.count('J')
    countls = [c1count, c2count, c3count, c4count, c5count]
    if countls.count(5 - jcount) == (5 - jcount):    #five of a kind
        HandType = 6
    elif (countls.count(3) == (3) and jcount == 1) or (countls.count(4 - jcount) == 4 and jcount == 2) or (countls.count(1) == 2 and jcount == 3) or countls.count(4) == 4:    #four of a kind, 1st or 2nd must match
        HandType = 5
    elif ((c1count == 3 or c2count == 3 or c3count == 3) and (c1count == 2 or c2count == 2 or c3count == 2 or c4count == 2)) or (countls.count(2) == 4 and jcount == 1):    # full house
        HandType = 4
    elif countls.count(3 - jcount) == (3 - jcount) or countls.count(3 - jcount) == 3:    #three of a kind, 1st or 2nd or 3rd must match
        HandType = 3
    elif countls.count(2) == 4 or (countls.count(2) == 2 and jcount == 1):   #two pair
        HandType = 2
    elif countls.count(2) == 2 or jcount == 1:     #one pair
        HandType = 1
    elif countls.count(1) == 5:
        HandType = 0     
    return HandType

#Part 1 resources
cards = [[],[],[],[],[],[],[]]
cardValues = [[999999],[999999],[999999],[999999],[999999],[999999],[999999]]
curRank = 1
hextrans = str.maketrans('AKQJT', 'EDCBA')

#Part 2 resources
cardsP2 = [[],[],[],[],[],[],[]]
cardValuesP2 = [[999999],[999999],[999999],[999999],[999999],[999999],[999999]]
curRankP2 = 1
hextransP2 = str.maketrans('AKQTJ', 'DCBA1')

for i in Data:
    hand = i.strip().split()
    HandType = getType(hand[0]) 
    cardValues[HandType].append(int(hand[0].translate(hextrans), 16))
    cards[HandType].append(hand)
    #Part 2
    HandType = getTypeP2(hand[0])
    cardValuesP2[HandType].append(int(hand[0].translate(hextransP2), 16))
    cardsP2[HandType].append(hand)


#Part 1
for i in range(0, len(cardValues)):
    while len(cardValues[i]) > 1:
        k = numpy.argmin(cardValues[i])
        P1 += int(cards[i][k-1][1]) * curRank
        curRank += 1
        cardValues[i].pop(k)
        cards[i].pop(k-1)

#Part2
for i in range(0, len(cardValuesP2)):
    while len(cardValuesP2[i]) > 1:
        k = numpy.argmin(cardValuesP2[i])
        P2 += int(cardsP2[i][k-1][1]) * curRankP2
        curRankP2 += 1
        cardValuesP2[i].pop(k)
        cardsP2[i].pop(k-1)

print('Part 1 Total : ', P1 )
print('Part 2 Total : ', P2 )