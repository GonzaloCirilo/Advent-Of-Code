import Foundation


func solve() {
    var lines: [String] = []
    while let line = readLine() {
        lines.append(line)
    }
    var sol = 0
    for line in lines {
        let tokens = line.split(separator: " ")
        var map = String(tokens[0])
        let arragenment: [Int] = tokens[1].split(separator: ",").map { s in
            Int(s) ?? 0
        }
        var p = 0
        var indexMap: [Int: Int] = [:]
        for (j, t) in map.enumerated() {
            if t == "?" {
                indexMap[j] = p
                p += 1
                
            }
        }
        
        map += "."
        var ans = 0
        for mask in 0...((1 << (p))-1) {
            var aux: String = ""
            for bit in 0...(p) {
                if (mask >> bit & 1) == 1 {
                    aux += "#"
                } else {
                    aux += "."
                }
            }
            
            var countD = 0
            var isleInd = 0
            var valid = true
            var x = ""
            for (l, t) in map.enumerated() {
                var c: String = ""
                if t == "?" {
                    let ind = aux.index(aux.startIndex, offsetBy: indexMap[l] ?? 0)
                    c = String(aux[ind])
                } else {
                    let ind = map.index(map.startIndex, offsetBy: l)
                    c = String(map[ind])
                }
                if (c == "#") {
                    countD += 1
                } else if c == "." && countD > 0 {
                    if(isleInd < arragenment.count && countD != arragenment[isleInd]) {
                        valid = false
                        //break
                    }
                    countD = 0
                    isleInd += 1
                }
                x += c
            }
            //print("    " + x + " " + String(valid && isleInd == arragenment.count))
            if valid && isleInd == arragenment.count {
                ans += 1
            }

        }
        //print(ans)
        sol += ans
    }
    print(sol)
}

solve()
