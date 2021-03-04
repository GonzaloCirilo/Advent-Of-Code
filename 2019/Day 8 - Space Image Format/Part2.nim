import terminal

let layerBound = 25 * 6

proc solve() =
    var 
        message: array[25 * 6, int]
        cont = 0

    for i in 0 ..< layerBound:
        message[i] = 2

    while true:
        let c = getch()
        if c < '0' or c > '9':
            break

        let ind = cont mod layerBound

        if message[ind] == 2 and c != '2':
            message[ind] = (case c
                of '1': 1    
                of '0': 0
                else: 2)
        
        cont = cont + 1
    
    for i in 0 ..< layerBound:
        if i mod 25 == 0 and i != 0:
            write(stdout, '\n')
        write(stdout, message[i])
        
solve()
