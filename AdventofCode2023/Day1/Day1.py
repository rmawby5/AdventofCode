f = open('Day1\InputTest.txt', 'r')
Data = f.readlines()
f.close


alNumdict = {
    'one': 1,
    'two': 2,
    'three': 3,
    'four': 4,
    'five': 5,
    'six': 6,
    'seven': 7,
    'eight': 8,
    'nine': 9
}



def stringParseP1(stg):
    ls = list(stg.strip())
    Num = ''
    lstNum = ''
    for i in ls:
        if Num == '':
            try:
                int(i)
                Num = i
            except:
                None
        else:
            try:
                int(i)
                lstNum = i
            except:
                None
    if lstNum == '':
        lstNum = Num
    Num = int(Num + lstNum)
    return Num



def stringParseP2(stg):
    buffer = ''
    f_Num = 0
    l_Num = 0
    
    for i in stg:
        if f_Num == 0:
            print(i)
            try:
                f_Num = int(i)
            except:
                
                buffer = buffer + i
                
                try:
                    print(buffer)
                    f_Num = alNumdict[buffer]
                except:
                    None
    print('found : ', f_Num)
    buffer = ''

    for i in range(len(stg) - 1, 0, -1):
        if l_Num == 0:
            j = stg[i]
            #print(j)
            try:
                l_Num = int(j)
            except:
                buffer = j + buffer
                
                try:
                    print(buffer)
                    l_Num = alNumdict[buffer]
                except:
                    None
    print('found : ', l_Num)

    return int(str(f_Num) + str(l_Num))

totalP1 = 0
totalP2 = 0
for i in Data:
    
    totalP1 += stringParseP1(i)
    totalP2 += stringParseP2(i)

print('Part 1 total : ', totalP1)
print('Part 2 total : ', totalP2)