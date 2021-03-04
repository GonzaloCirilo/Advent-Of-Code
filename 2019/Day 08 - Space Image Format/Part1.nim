import terminal

let layerBound = 25 * 6

proc solve() =
    var 
        min = 99999999
        z = 0
        o = 0
        t = 0
        cont = 0
        ans = 0
    while true:
        let c = getch()
        if c < '0' or c > '9':
            break

        if c == '0':
            z = z + 1
        elif c == '1':
            o = o + 1
        elif c == '2':
            t = t + 1
        
        if cont mod layerBound == 0 and cont != 0:
            if z < min:
                min = z
                ans = o * t
            z = 0
            o = 0
            t = 0
        cont = cont + 1
    
    write(stdout, ans)


solve()
