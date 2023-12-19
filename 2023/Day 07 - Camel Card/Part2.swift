import Foundation

enum HandType: Comparable {
	case fiveOfAKind
	case fourOfAKind
	case fullHouse
	case threeOfAKind
	case twoPair
	case onePair
	case highCard

	private var sortOrder: Int {
		switch self {
		case .highCard:
			return 0
		case .onePair:
			return 1
		case .twoPair:
			return 2
		case .threeOfAKind:
			return 3
		case .fullHouse:
			return 4
		case .fourOfAKind:
			return 5
		case .fiveOfAKind:
			return 6
		}
	}

	static func ==(l: HandType, r: HandType) -> Bool {
		return l.sortOrder == r.sortOrder
	}

	static func <(l: HandType, r: HandType) -> Bool {
		return l.sortOrder < r.sortOrder
	}
}

func evaluateHand(_ s: String) -> HandType {
	var h: [Character: Int] = [:]

	var jCount = 0
	s.forEach { c in
		if c == "J" {
			jCount += 1
		} else {
			h.updateValue((h[c] ?? 0) + 1, forKey: c)
		}
	}
	var maxe = -1
	var maxK: Character = "."
	h.forEach { k, v in
		if (maxe < v) {
			maxe = v
			maxK = k
		}
	}

	var aux = 1
	h.forEach { k, v in
		if k == maxK {
			aux *= (v + jCount)
		} else {
			aux *= v
		}
	}
	if jCount == 5 {
		aux = 5
	}
	switch aux {
	case 1:
		return .highCard
	case 2:
		return .onePair
	case 3:
		return .threeOfAKind
	case 4:
		if (h.count == 2) {
			return .fourOfAKind
		} else {
			return .twoPair
		}
	case 6:
		return .fullHouse
	case 5:
		return .fiveOfAKind
	default:
		return .highCard
	}
}

func handValue(_ s: String) -> [Int] {
	return s.map { c in
		switch c {
	case "A":
		return 10000
		case "K":
			return 100
		case "Q":
			return 99
		case "J":
			return -1
		case "T":
			return 10
		default:
			return c.wholeNumberValue ?? 0
	}
	}
}

struct Hand {
	var hand: String
	var handType: HandType
	var bid : Int
}

func solve() {
	var lines: [String] = []
	while let line = readLine() {
		lines.append(line)
	}

	var hands = lines.map { s in
		let tokens = s.split(separator: " ")
		return Hand(hand: String(tokens[0]), handType: evaluateHand(String(tokens[0])), bid: Int(String(tokens[1])) ?? 0)
	}
	hands.sort { a, b in
		if a.handType == b.handType {
			let bs = handValue(b.hand)
			for (i, x) in handValue(a.hand).enumerated() {
				if x == bs[i] {
					continue
				}
				return x < bs[i]
			}
			return false
		}
		return a.handType < b.handType
	}
	var ans = 0
	for (r, h) in hands.enumerated() {
		print(r + 1, h.bid, h.hand, h.handType)
		ans += (r + 1) * h.bid
	}

	print(ans)

}

solve()
