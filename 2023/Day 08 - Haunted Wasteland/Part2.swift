import Foundation

struct Dest {
    var rigth: String
    var left: String
}

func gcd(_ x: Int, _ y: Int) -> Int {
    var a = 0
    var b = max(x, y)
    var r = min(x, y)
    
    while r != 0 {
        a = b
        b = r
        r = a % b
    }
    return b
}

func lcm(_ x: Int, _ y: Int) -> Int {
    return x / gcd(x, y) * y
}

func hasEnded(_ token: String) -> Bool {
    let ss = token.map { c in c }
    if ss[2] != "Z" {
        return false
    }
    return true
}

func solve() {
    var lines: [String] = []
    var dict: [String: Dest] = [:]
    let inst = readLine()
    readLine()
    while let line = readLine() {
        lines.append(String(line.prefix(3)))
        let startIndexL = line.index(line.startIndex, offsetBy: 7)
        let endIndexL = line.index(line.endIndex, offsetBy: -6)
        let rangeL = startIndexL..<endIndexL
        let startIndexR = line.index(line.startIndex, offsetBy: 12)
        let endIndexR = line.index(line.endIndex, offsetBy: -1)
        let rangeR = startIndexR..<endIndexR
        dict[String(line.prefix(3))] = Dest(rigth: String(line[rangeR]), left: String(line[rangeL]))
        //print(dict)
    }
    
    var currents: [String] = []
    lines.forEach { s in
        let ss = s.map { c in
            c
        }
        if (ss[2] == "A") {
            currents.append(s)
        }
    }
    //print(currents)
    let answers = currents.map({ auxC in
        var current = auxC
        var steps = 0
        var instC = 0
        let input = inst!.map { c in
            c
        }
        while !hasEnded(current) {
            let c = input[instC]
            //print(c, current)
            if c == "L" {
                current = dict[current]?.left ?? "break"
            } else if c == "R" {
                current = dict[current]?.rigth ?? "break"
            }
            
            steps += 1
            instC = (instC + 1) % input.count
            if (steps == 10000000) {
                break
            }
        }
        return steps
    })
    let ans = answers.reduce(1) { partialResult, x in
        return lcm(partialResult, x)
    }
    print(ans)
}

solve()
