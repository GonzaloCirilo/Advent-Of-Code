fun hash(s: String): Int {
	var value = 0
	for (c in s) {
		value += c.toInt()
		value *= 17
		value %= 256
	}
	return value
}

fun main() {
	var ans = 0
	var seq = readln().split(",")
	var dict = mutableMapOf<Int, MutableList<Pair<String, Int>>>()
	for (s in seq) {
		val op = if (s.contains('=')) {
			s.split("=")
		} else {
			s.split("-")
		}
		val box = hash(op[0])
		if (s.contains('=')) {
			val lens = Pair(op[0],op[1].toInt())
			if (!dict.contains(box)) {
				dict[box] = mutableListOf(lens)
			} else {
				val isInList = dict[box]?.map { it.first }?.indexOf(op[0]) ?: -1
				if (isInList != -1){
					dict[box]?.set(isInList, lens)
				} else {
					dict[box]?.add(lens)
				}
			}
		} else {
			if (dict.contains(box)) {
				dict[box]?.removeIf { it.first == op[0] }
			}
		}
	}
	for ((k,v) in dict) {
		v.forEachIndexed { index, pair ->
			//println("${(k + 1)} ${(index + 1)} ${pair.second}")
			ans += (k + 1) * (index + 1) * pair.second
		}
	}
	println(ans)
}