import Foundation


func solve() {
    var lines: [String] = []
    var copies: [Int] = []
    while let line = readLine() {
        lines.append(line)
		copies.append(1)
	}
	for (i, line) in lines.enumerated() {
		var aux = line
		var card = line.split(separator: ":")[1].split(separator: "|")
		var winning = card[0].trimmingCharacters(in: .whitespaces).split(separator: " ")
		var current = card[1].trimmingCharacters(in: .whitespaces).split(separator: " ")
		var sWinning: Set<Int> = []
		for s in winning {
			var n = Int(s) ?? -1
			if n != -1 {
				sWinning.insert(n)
			}
		}
		var match = 0
		for s in current {
			var n = Int(s) ?? -1
			if n != -1 {
				if sWinning.contains(n) {
					match += 1
				}
			}
		}
		if match > 0 {
			for x in 1 ... match {
				if i + x < copies.count {
					copies[i + x] += copies[i]
				}
			}
		}
	}
	var sum = 0
	for c in copies {
		sum += c
	}
	print(sum)
}

solve()
