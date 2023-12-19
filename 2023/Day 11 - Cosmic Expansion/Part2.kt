import kotlin.math.abs
import kotlin.math.min
import kotlin.math.max

fun main() {
	val input = generateSequence(::readLine).toList()
	val multiplier = 1000000
	val rowSpace = mutableListOf<Int>()
	val colSpace = mutableListOf<Int>()
	input.forEachIndexed { i, line ->
		for (j in line.indices) {
			if (input[i][j] != '.') {
				return@forEachIndexed
			}
		}
		rowSpace.add(i)
	}

	for(j in input[0].indices) {
		var cleanCol = true
		for(i in input.indices) {
			if (input[i][j] != '.') {
				cleanCol = false
				break
			}
		}
		if (!cleanCol) continue
		colSpace.add(j)
	}

	val galaxies = mutableListOf<Pair<Int, Int>>()
	for (i in input.indices) {
		for (j in input[i].indices) {
			if (input[i][j] == '#') {
				galaxies.add(Pair(i,j))
			}
		}
	}
	var ans: Long = 0L
	for (i in galaxies.indices) {
		for (j in (i+1) until galaxies.size) {
			val myRowRange = min(galaxies[i].first, galaxies[j].first)..max(galaxies[i].first, galaxies[j].first)
			var rowOccurrences = 0
			rowSpace.forEach { 
				if (myRowRange.contains(it)) {
					rowOccurrences++
				}
			}
			val myColRange = min(galaxies[i].second, galaxies[j].second)..max(galaxies[i].second, galaxies[j].second)
			var colOccurrences = 0
			colSpace.forEach { 
				if (myColRange.contains(it)) {
					colOccurrences++
				}
			}
			ans += abs(galaxies[i].first - galaxies[j].first) + (rowOccurrences * (multiplier-1)) + abs(galaxies[i].second - galaxies[j].second) + (colOccurrences * (multiplier-1))
		}
	}
	println(ans)
}