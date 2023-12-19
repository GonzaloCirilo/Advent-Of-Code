import java.util.LinkedList
import java.util.Queue
import kotlin.math.max

enum class Direction {
	NORTH,
	SOUTH,
	WEST,
	EAST
}

data class State(
	val pos: Pair<Int, Int>,
	val depth: Int
)

fun main() {
	val input = generateSequence(::readLine).toList()
	var coord = Pair(0, 0)
	input.forEachIndexed { i, line ->
		if (line.contains("S")) {
			coord = Pair(i, line.indexOf("S"))
		}
	}

	val q: Queue<State> = LinkedList()
	val visited = input.map { it.map { _ -> -1 }.toMutableList() }
	visited[coord.first][coord.second] = 0
	q.add(State(coord, 0))
	var maxe = -1
	while (q.isNotEmpty()) {
		val n = q.remove()
		maxe = max(maxe, n.depth)
		val pipe = input[n.pos.first][n.pos.second]
		for (d in Direction.values()) {
			val newCoord: Pair<Int, Int> = when (d) {
				Direction.NORTH -> {
					if (n.pos.first - 1 < 0) continue
					if (pipe == 'F' || pipe == '7'|| pipe == '-') continue
					val newPipe = input[n.pos.first - 1][n.pos.second]
					if (newPipe == 'L' || newPipe == 'J' || newPipe == '-' || newPipe == '.') continue
					n.pos.copy(first = n.pos.first - 1)
				}
				Direction.SOUTH -> {
					if (n.pos.first + 1 >=  input.size) continue
					if (pipe == 'L' || pipe == 'J'|| pipe == '-') continue
					val newPipe = input[n.pos.first + 1][n.pos.second]
					if (newPipe == 'F' || newPipe == '7' || newPipe == '-' || newPipe == '.') continue
					n.pos.copy(first = n.pos.first + 1)
				}
				Direction.WEST -> {
					if (n.pos.second - 1 < 0) continue
					if (pipe == 'L' || pipe == 'F'|| pipe == '|') continue
					val newPipe = input[n.pos.first][n.pos.second - 1]
					if (newPipe == 'J' || newPipe == '7' || newPipe == '|' || newPipe == '.') continue
					n.pos.copy(second = n.pos.second - 1)
				}
				Direction.EAST -> {
					if (n.pos.second + 1 >= input[0].length) continue
					if (pipe == 'J' || pipe == '7'|| pipe == '|') continue
					val newPipe = input[n.pos.first][n.pos.second + 1]
					if (newPipe == 'F' || newPipe == 'L' || newPipe == '|' || newPipe == '.') continue
					n.pos.copy(second = n.pos.second + 1)}
			}
			if (visited[newCoord.first][newCoord.second] == -1) {
				visited[newCoord.first][newCoord.second] = n.depth + 1
				q.add(State(newCoord, n.depth + 1))
			}
		}
	}
	var ans = 0
	val canTop = coord.first - 1 >= 0 && ((input[coord.first - 1][coord.second] == '|') || (input[coord.first - 1][coord.second] == 'F') || (input[coord.first - 1][coord.second] == '7'))
	val canBottom = coord.first + 1 < input.size && ((input[coord.first + 1][coord.second] == '|') || (input[coord.first + 1][coord.second] == 'J') || (input[coord.first + 1][coord.second] == 'L'))
	val canLeft = coord.second - 1 >= 0 && ((input[coord.first][coord.second - 1] == '-') || (input[coord.first][coord.second - 1] == 'F') || (input[coord.first][coord.second - 1] == 'L'))
	val canRight = coord.second + 1 < input[0].length && ((input[coord.first][coord.second + 1] == '-') || (input[coord.first][coord.second + 1] == 'J') || (input[coord.first][coord.second + 1] == '7'))

	val sToken = when {
		canBottom && canTop -> '|'
		canRight && canLeft -> '-'
		canTop && canRight -> 'L'
		canTop && canLeft -> 'J'
		canBottom && canLeft -> '7'
		canBottom && canRight -> 'F'
		else -> 'S'
	}
	visited.forEachIndexed { i, row ->
		row.forEachIndexed row@ { j, _ ->
			if (visited[i][j] == -1) {
				var count = 0
				var prev = '_'
				for (nj in 0 until j) {
					if (visited[i][nj] == -1) continue
					val newToken = if (i == coord.first && nj == coord.second) {
						sToken
					} else {
						input[i][nj]
					}
					if (newToken == 'F' || newToken == 'L')
						prev = newToken
					if (visited[i][nj] != -1 && ((newToken == '|') ||  (newToken == 'J' && prev == 'F') || (newToken == '7' && prev == 'L'))){
						count++
						prev = '_'
					} 
				}
				if (count % 2 == 0) {
					//print(".")
					return@row
				}
				prev = '_'
				count = 0
				for (nj in (j + 1) until visited[0].size) {
					if (visited[i][nj] == -1) continue
					val newToken = if (i == coord.first && nj == coord.second) {
						sToken
					} else {
						input[i][nj]
					}
					
					if (newToken == 'F' || newToken == 'L')
						prev = newToken
					if (visited[i][nj] != -1 && ((newToken == '|') ||  (newToken == 'J' && prev == 'F') || (newToken == '7' && prev == 'L'))){
						count++
						prev = '_'
					}
				}
				if (count % 2 == 0) {
					//print(".")
					return@row
				}
				//print("I")
				ans++
			}
			else {
				//print("#")
			}
		}
		//print("\n")
	}
	println(ans)
}