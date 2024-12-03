f = open("Day16\input.txt","r")
lines = f.read().splitlines()
f.close()
 
map = {}
ycoord = 0
for y in lines:
    arr = list(y)
    xcoord = 0
    for x in arr:
        map[(xcoord,ycoord)] = x
        xcoord += 1
    ycoord += 1
 
energized = map.copy()
 
stack = [((0,0),(1,0))]
count = 0
 
while count != 100000:
    count += 1
    val = ()
    try:
        val = stack.pop()
    except:
        break
    curr = val[0]
    dir = val[1]
    exit = True
    for x in curr:
        if x < 0 or x > 109:
            exit = False
    if exit:
        energized[curr] = "#"
        if map[curr] == '.':
            curr = (curr[0] + dir[0], curr[1] + dir[1])
            stack.append((curr,dir))
        elif map[curr] == '|':
            if dir[1] != 1:
                dir1 = (0,-1)
                curr1 = (curr[0] + dir1[0], curr[1] + dir1[1])
                stack.append((curr1,dir1))
            if dir[1] != -1:
                dir2 = (0,1)
                curr2 = (curr[0] + dir2[0], curr[1] + dir2[1])
                stack.append((curr2,dir2))
            map[curr] = '#'
        elif map[curr] == '-':
            if dir[0] != 1:
                dir1 = (-1,0)
                curr1 = (curr[0] + dir1[0], curr[1] + dir1[1])
                stack.append((curr1,dir1))
            if dir[0] != -1:
                dir2 = (1,0)
                curr2 = (curr[0] + dir2[0], curr[1] + dir2[1])
                stack.append((curr2,dir2))
            map[curr] = '#'
        elif map[curr] == '/':
            if dir == (1,0):
                dir = (0,-1)
                curr = (curr[0] + dir[0], curr[1] + dir[1])
                stack.append((curr,dir))
            elif dir == (-1,0):
                dir = (0,1)
                curr = (curr[0] + dir[0], curr[1] + dir[1])
                stack.append((curr,dir))
            elif dir == (0,1):
                dir = (-1,0)
                curr = (curr[0] + dir[0], curr[1] + dir[1])
                stack.append((curr,dir))
            else:
                dir = (1,0)
                curr = (curr[0] + dir[0], curr[1] + dir[1])
                stack.append((curr,dir))
        elif map[curr] == '\\':
            if dir == (1,0):
                dir = (0,1)
                curr = (curr[0] + dir[0], curr[1] + dir[1])
                stack.append((curr,dir))
            elif dir == (-1,0):
                dir = (0,-1)
                curr = (curr[0] + dir[0], curr[1] + dir[1])
                stack.append((curr,dir))
            elif dir == (0,1):
                dir = (1,0)
                curr = (curr[0] + dir[0], curr[1] + dir[1])
                stack.append((curr,dir))
            else:
                dir = (-1,0)
                curr = (curr[0] + dir[0], curr[1] + dir[1])
                stack.append((curr,dir))
 
vals = list(energized.values())
total = 0
for x in vals:
    if x == '#':
        total += 1
print(total)