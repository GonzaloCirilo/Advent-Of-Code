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

	visited.forEach {
		//println(it)
	}
	println(maxe)

}