import Foundation

func solve() {
    var lines: [String] = []
    while let line = readLine() {
        var aux = line
        aux.append(".")
        lines.append(aux)
    }
    var currentNumber = 0
    var isCurrentPar = false
    var ans = 0
    for i in 0 ..< lines.count {
        for j in 0 ..< lines[i].count {
            let c = Array(lines[i])[j]
            if (!c.isNumber) {
                if (isCurrentPar) {
                    ans += currentNumber
                }
                isCurrentPar = false
                currentNumber = 0
                continue
            }
            if let myNum = c.wholeNumberValue {
                currentNumber *= 10
                currentNumber += myNum
            }
            
            for k in -1 ... 1 {
                for l in -1 ... 1 {
                    let ni = i + k
                    let nj = j + l
                    if nj < 0 || ni < 0 || ni >= lines.count || nj >= lines[ni].count {
                        continue
                    }
                    let nc = Array(lines[ni])[nj]
                    if nc != "." {
                        if !nc.isNumber {
                            isCurrentPar = true
                        }
                    }
                }
            }
        }
    }
    print(ans)
}

solve()
