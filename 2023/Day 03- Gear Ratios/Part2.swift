import Foundation

struct Pair<T: Hashable, U: Hashable>: Hashable {
  let first: T
  let second: U
}

func solve() {
    var lines: [String] = []
    while let line = readLine() {
        var aux = line
        aux.append(".")
        lines.append(aux)
    }
    var currentNumber = 0
    var isCurrentPar = false
    var pos: Pair<Int, Int>? = nil
    var ans = 0
    var h = [Pair<Int, Int> : [Int]]()
    for i in 0 ..< lines.count {
        for j in 0 ..< lines[i].count {
            let c = Array(lines[i])[j]
            if (!c.isNumber) {
                if let p = pos {
                    print(p.first, p.second)
                }
                
                if (isCurrentPar) {
                    if let aux = pos {
                        (h[aux, default: []]).append(currentNumber)
                    }
                }
                isCurrentPar = false
                currentNumber = 0
                pos = nil
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
                            if nc == "*" {
                                pos = Pair(first: ni, second: nj)
                            }
                        }
                    }
                }
            }
        }
    }
    for (_, nums) in h {
        if nums.count == 2 {
            var aux = 1
            for x in nums {
                aux *= x
            }
            ans += aux
        }
    }
    print(ans)
}

solve()
