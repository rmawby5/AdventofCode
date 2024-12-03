import copy

f = open('Day19\Input.txt', 'r')
Data = f.readlines()
f.close

workflows = {}
parts = []
acceptedParts = []

partIndexValues = {
    'x':0,
    'm':1,
    'a':2,
    's':3
}


#Part 1
def compareCondition(condition, part):
    condition = condition.split(':')
    operator = condition[0][1]
    prop = condition[0][0]
    if operator == '>':
        if int(part[partIndexValues[prop]][2:]) > int(condition[0][2:]):
            return condition[1]
    elif operator == '<':
        if int(part[partIndexValues[prop]][2:]) < int(condition[0][2:]):
            return condition[1]
    return False

def comparePart(part, curworkflow):
    flowConditions= workflows[curworkflow]
    currentCond = 0
    result = False
    while result == False:
        if currentCond != len(flowConditions) - 1:
            i = flowConditions[currentCond]
            result = compareCondition(i, part)
        else:
            result = flowConditions[currentCond]
        currentCond += 1
    if result == 'A':                           #Add part to accepted list, then exit stack
        acceptedParts.append(part)
        return None
    elif result == 'R' or False:
        return None                             # part is rejected, exit stack
    else:                                    #result is for new workflow, add new level to stack for new workflow
        comparePart(part, result)
    return None

def sumTotal(part):
    partTotal = 0
    for cond in part:
        partTotal += int(cond[2:])
    return partTotal


for i in Data:
    if i != '\n':
        if i[0] != '{':
            flow = i.strip().split('{')
            flowName = flow[0]
            flowconditions = flow[1][:-1].split(',')
            workflows[flowName] = flowconditions
        else:
            parts.append(i[1:-2].split(','))

for i in parts:
    comparePart(i, 'in')

part1 = 0
for i in acceptedParts:
    part1 += sumTotal(i)
print('Part 1 Total : ', part1)


# Part 2
p2Total = 0
startCombinationsRanges = [[1, 4000], [1, 4000], [1, 4000], [1, 4000]]


def sumCombinations(partValueRanges):
    rangeTotal = 1
    for i in partValueRanges:
        rangeTotal = rangeTotal * (i[1] - i[0] + 1)
    return rangeTotal

def adjustPartValue(partVal, condition):
    partValue = copy.deepcopy(partVal)
    condition = condition.split(':')
    operator = condition[0][1]
    prop = condition[0][0]
    if operator == '>':
        if partValue[partIndexValues[prop]][0] < int(condition[0][2:]):             #if required minimum is higher than current, update current minimum for the path
            partValue[partIndexValues[prop]][0] = int(condition[0][2:]) + 1
            return partValue
        elif partValue[partIndexValues[prop]][1] < int(condition[0][2:]):           #if required minimum is higher than current max, path ends
            return False
    elif operator == '<':

        if partValue[partIndexValues[prop]][1] > int(condition[0][2:]):             #if required max is higher than current, update current max for the path
            partValue[partIndexValues[prop]][1] = int(condition[0][2:]) - 1
            return partValue
        elif partValue[partIndexValues[prop]][0] > int(condition[0][2:]):           #if required maximum is higher than current min, path ends
            return False
    return False

def adjustPartValueInv(partVal, condition):                                         #gets range to fail condition check, before moving to next in sequence
    partValue = copy.deepcopy(partVal)
    condition = condition.split(':')
    operator = condition[0][1]
    prop = condition[0][0]
    if operator == '<':
        if partValue[partIndexValues[prop]][0] < int(condition[0][2:]):             
            partValue[partIndexValues[prop]][0] = int(condition[0][2:])
            return partValue
        elif partValue[partIndexValues[prop]][1] < int(condition[0][2:]):           
            return False
    elif operator == '>':
        if partValue[partIndexValues[prop]][1] > int(condition[0][2:]):             
            partValue[partIndexValues[prop]][1] = int(condition[0][2:])
            return partValue
        elif partValue[partIndexValues[prop]][0] > int(condition[0][2:]):           
            return False
    return False

def combination(partValues, curWorkflow):
    flowcond = workflows[curWorkflow]
    global p2Total
    for i in range(0, len(flowcond)):
        if i == len(flowcond) - 1:
            if flowcond[i] == 'R':
                next
            elif flowcond[i] == 'A':
                p2Total += sumCombinations(partValues)
            else:
                combination(partValues, flowcond[i])
        else:
            result = adjustPartValue(partValues, flowcond[i])
            if result != False:                                         #condition meet, get next step                                                                  
                step = flowcond[i].split(':')[1]
                if step == 'R':
                    next
                elif step == 'A':
                    p2Total += sumCombinations(result)
                else:
                    combination(result, step)
                partValues =  adjustPartValueInv(partValues, flowcond[i])
    return None


combination(startCombinationsRanges, 'in')
print('Part 2 Total : ', p2Total)