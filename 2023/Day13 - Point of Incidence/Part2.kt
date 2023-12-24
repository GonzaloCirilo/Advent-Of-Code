fun invertMap(map: List<String>): List<String> {
    println("inverting")
    return map.let {
        val aux = MutableList(map[0].length) { "" }
        for (i in it.indices) {
			for (j in it[i].indices.reversed()) {
				aux[j] = aux[j] + it[i][j]
			}
		}
		aux
	}
}

fun stringDifference(a: String, b: String): Int {
	var delta = 0
	for (i in a.indices) {
		if (a[i] != b[i]) {
			delta++
		}
	}
	return delta
}

fun findMirror(map: List<String>): Int {
	for (j in 0 until (map.size - 1)) {
		var lPivot = j
		var rPivot = lPivot + 1
		var hasChangedSmudge = false
		while (map[lPivot] == map[rPivot] || (stringDifference(map[lPivot], map[rPivot]) == 1 && !hasChangedSmudge)) {
			if (stringDifference(map[lPivot], map[rPivot]) == 1) {
				hasChangedSmudge = true
			}
			lPivot--
			rPivot++
			if (lPivot < 0 || rPivot >= map.size) {
				if (hasChangedSmudge) {
					return j + 1
				} else {
					break
				}
			}
		}
	}
	println("not_found")
	return 0
}

fun main() {
	var ans = 0L
	val map = emptyList<String>().toMutableList()
	generateSequence(::readLine).toList().forEach {
		if (it.isEmpty()) {
			// horizontal * 100 + vertical. Either of each will return a values, the other one should be zero
			ans += findMirror(map) * 100 + findMirror(invertMap(map))
			println(ans)
			map.clear()
			return@forEach
		}
		map.add(it)
	}
	println(ans)
}