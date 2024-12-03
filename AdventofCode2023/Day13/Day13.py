import numpy

f = open('Day13\Input.txt', 'r')
Data = f.readlines()
f.close

row = []
total = 0
totalSmudge = 0

def offByOne(l1, l2):
    mismatch = 0
    index = 0
    for i in range(0, len(l1)):
        if l1[i] != l2[i]:
            mismatch += 1
            index = i          #position of mismatch  
    if mismatch != 1:
        index = -1              #if multiple mismatches are found, reset index value to default
    return index
    
def charswap(cha):
    newchar = '#'
    if cha == '#':
        newchar = '.'
    return newchar


def countMirLin (grid):
    mirLine = 0           
    blkmirLine = 0        
    mirrored = True
    for i in range(1, len(grid)):
        if grid[i] == grid[i - 1]:
            offset = 1
            mirrored = True
            if (i - 1) == 0 or i == 0:
                mirLine = i               
            else:
                while (i - 1 - offset) > -1 and (i + offset) < len(grid):
                    if grid[i + offset] == grid[i -1 - offset]:
                        offset += 1
                    else:
                        mirrored = False
                        break              
                if mirrored == False:
                    mirLine = 0
                else:
                    mirLine = i
        else:
            mirLine = 0
        blkmirLine += mirLine
    return blkmirLine


def countMirLinSmudge (grid, normmirline):               
    mirLine = 0           
    blkmirLine = 0        
    mirrored = True
    remSmud = False
    reset = ''
    resetloc = []
    for i in range(1, len(grid)):
        if grid[i] == grid[i - 1] or offByOne(grid[i], grid[i - 1]) != -1:
            if offByOne(grid[i], grid[i - 1]) != -1:
                offIndex = offByOne(grid[i], grid[i - 1])
                reset = grid[i][offIndex]
                grid[i][offIndex] = charswap(grid[i][offIndex])               
                resetloc = [i, offIndex]
                remSmud = True    
            offset = 1
            mirrored = True
            if (i - 1) == 0 or i == 0:
                mirLine = i               
            else:
                while (i - 1 - offset) > -1 and (i + offset) < len(grid):
                    if grid[i + offset] == grid[i -1 - offset]:
                        offset += 1
                    else:
                        offIndex = offByOne(grid[i + offset], grid[i - 1 - offset])
                        if offIndex == -1 or remSmud == True:
                            mirrored = False
                            if remSmud == True:
                                grid[resetloc[0]][resetloc[1]] = reset
                                remSmud = False
                            break
                        else:
                            reset = grid[i + offset][offIndex]
                            grid[i + offset][offIndex] = charswap(grid[i + offset][offIndex])                           
                            resetloc = [i + offset, offIndex]
                            remSmud = True            
                if mirrored == False:
                    mirLine = 0
                else:
                    mirLine = i
        else:
            mirLine = 0
        if mirLine != normmirline:
            blkmirLine += mirLine
    return blkmirLine


def processBlock(rows):
    columns = []
    blockTotal_smudge = 0
    columnTemp = numpy.transpose(rows)
    for i in columnTemp:
        columns.append(list(i))   
    blockTotal_row = countMirLin(rows)
    blockTotal_col = countMirLin(columns) 
    blockTotal = (blockTotal_row * 100) + blockTotal_col
    blockTotal_smudge = (countMirLinSmudge(rows, blockTotal_row) * 100) + countMirLinSmudge(columns, blockTotal_col)
    return blockTotal, blockTotal_smudge  

for i in Data:   
    if i != '\n':
        i = list(i)
        i.pop()
        row.append(i)
    elif i == '\n':             #Stop and process current data, then reset
        totals = processBlock(row)
        total += totals[0]
        totalSmudge += totals[1]
        row = []
        #print('break')
    

print('P1 Total: ', total)
print('p2 Total: ', totalSmudge)