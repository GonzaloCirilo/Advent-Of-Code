fun invertMap(map: List<String>): List<String> {
	//println("inverting")
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

fun findMirror(map: List<String>): Int {
	for (j in 0 until (map.size - 1)) {
		var lPivot = j
		var rPivot = lPivot + 1
		while (map[lPivot] == map[rPivot]) {
			lPivot--
			rPivot++
			if (lPivot < 0 || rPivot >= map.size) {
				return j + 1
			}
		}
	}
	//println("not_found")
	return 0
}

fun main() {
	var ans = 0L
	val map = emptyList<String>().toMutableList()
	// add two lines to the input to work
	generateSequence(::readLine).toList().forEach {
		if (it.isEmpty()) {
			// horizontal * 100 + vertical. Either of each will return a values, the other one should be zero
			ans += findMirror(map) * 100 + findMirror(invertMap(map))
			//println(ans)
			map.clear()
			return@forEach
		}
		map.add(it)
	}
	println(ans)
}