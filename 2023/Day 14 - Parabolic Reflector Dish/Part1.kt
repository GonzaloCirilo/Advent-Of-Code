fun main() {
	var ans = 0
	val map = generateSequence(::readLine).toMutableList()
	map.add(map[0].map { '#' }.joinToString(""))
	var lastBlock = 0
	var rockCount = 0
	for (j in map[0].indices) {
		lastBlock = 0
		for (i in map.indices) {
			if (map[i][j] == '#') {
				val points = (map.size - 1) - lastBlock
				var delta = points - rockCount
				ans += (points * (points + 1) / 2) - (delta * (delta + 1) / 2)
				rockCount = 0
				lastBlock = i + 1
			}
			if (map[i][j] == 'O') {
				rockCount++
			}
		}
	}
	println(ans)
}