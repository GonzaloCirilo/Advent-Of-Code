import kotlin.math.abs

fun main() {
	val input = generateSequence(::readLine).toList()
	var space = input.map { it.toMutableList() }.toMutableList()
	input.reversed().forEachIndexed { i, line ->
		for (j in line.indices.reversed()) {
			if (input[input.size - 1 - i][j] != '.') {
				return@forEachIndexed
			}
		}
		space.add(input.size - 1 - i, line.toMutableList())
	}

	for(j in input[0].indices.reversed()) { 
		var cleanCol = true
		for(i in input.indices.reversed()) {
			if (input[i][j] != '.') {
				cleanCol = false
				break
			}
		}
		if (!cleanCol) continue
		for(i in space.indices.reversed()) {
			space[i].add(j, '.')
		}
	}

	val galaxies = mutableListOf<Pair<Int, Int>>()
	for (i in space.indices) {
		for (j in space[i].indices) {
			if (space[i][j] == '#') {
				galaxies.add(Pair(i,j))
			}
		}
	}
	var ans = 0
	//println(galaxies.size)
	for (i in galaxies.indices) {
		for (j in (i+1) until galaxies.size) {
			//println("${i + 1} ${j + 1}: ${abs(galaxies[i].first - galaxies[j].first) + abs(galaxies[i].second - galaxies[j].second)}")
			ans += abs(galaxies[i].first - galaxies[j].first) + abs(galaxies[i].second - galaxies[j].second)
		}
	}
	println(ans)
}