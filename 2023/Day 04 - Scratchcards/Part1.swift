import Foundation


func solve() {
	var sum = 0
	while let line = readLine() {
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
			sum += Int(pow(2, match-1) as NSDecimalNumber)
		}
	}
	print(sum)
}

solve()
