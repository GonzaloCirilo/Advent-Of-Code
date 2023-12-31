import java.util.LinkedList
import java.util.Queue
import kotlin.math.max

enum class Direction {
					 UP, DOWN, LEFT, RIGHT
}

data class Position(val x: Int, val y: Int) {
	fun move(direction: Direction): Position {
		return when (direction) {
			Direction.UP -> Position(x, y - 1)
			Direction.DOWN -> Position(x, y + 1)
			Direction.LEFT -> Position(x - 1, y)
			Direction.RIGHT -> Position(x + 1, y)
		}
	}
}

fun isInBounds(grid: List<String>, position: Position): Boolean {
	return position.y >= 0 && position.y < grid.size && position.x >= 0 && position.x < grid[position.y].length
}

fun hasCompletelyVisited(grid: List<String>, visited: List<List<Pair<Int, Int>>>, newPos: Position): Boolean {
	return visited[newPos.y][newPos.x].first == 1 && visited[newPos.y][newPos.x].second == 1
}

fun evaluateVis(grid: List<String>, visited: List<List<Pair<Int, Int>>>, newPos: Position, dir: Direction): Pair<Int, Int> {
	return when (dir) {
		Direction.LEFT -> {
			when (grid[newPos.y][newPos.x]) {
				'\\' -> visited[newPos.y][newPos.x].copy(first = 1)
				'/' -> visited[newPos.y][newPos.x].copy(second = 1)
				else -> 1 to 1
			}
		}

		Direction.UP -> {
			when (grid[newPos.y][newPos.x]) {
				'\\' -> visited[newPos.y][newPos.x].copy(second = 1)
				'/' -> visited[newPos.y][newPos.x].copy(second = 1)
				else -> 1 to 1
			}
		}

		Direction.DOWN -> {
			when (grid[newPos.y][newPos.x]) {
				'\\' -> visited[newPos.y][newPos.x].copy(first = 1)
				'/' -> visited[newPos.y][newPos.x].copy(first = 1)
				else -> 1 to 1
			}
		}

		Direction.RIGHT -> {
			when (grid[newPos.y][newPos.x]) {
				'\\' -> visited[newPos.y][newPos.x].copy(second = 1)
				'/' -> visited[newPos.y][newPos.x].copy(first = 1)
				else -> 1 to 1
			}
		}
	}
}

fun getEnergizedTileCount(grid: List<String>, sPosition: Position, initialDirection: Direction): Int {
	val visited = mutableListOf<MutableList<Pair<Int, Int>>>()
	grid.forEach { s ->
		visited.add(s.map { 0 to 0 }.toMutableList())
	}
	visited[sPosition.y][sPosition.x] = evaluateVis(grid, visited, sPosition, initialDirection)
	val queue: Queue<Pair<Position, Direction>> = LinkedList()
	queue.add(sPosition to initialDirection)

	while (queue.isNotEmpty()) {
		val (pos, dir) = queue.remove()
		val currentGridType = grid[pos.y][pos.x]
		when (dir) {
			Direction.UP -> {
				when (currentGridType) {
					'\\' -> {
						val newPos = pos.move(Direction.LEFT)
						if (isInBounds(grid, newPos) && (!hasCompletelyVisited(grid, visited, newPos) || grid[newPos.y][newPos.x] == '.')) {
							visited[newPos.y][newPos.x] = evaluateVis(grid, visited, newPos, Direction.LEFT)
							queue.add(newPos to Direction.LEFT)
						}
					}

					'/' -> {
						val newPos = pos.move(Direction.RIGHT)
						if (isInBounds(grid, newPos) && (!hasCompletelyVisited(grid, visited, newPos) || grid[newPos.y][newPos.x] == '.')) {
							visited[newPos.y][newPos.x] = evaluateVis(grid, visited, newPos, Direction.RIGHT)
							queue.add(newPos to Direction.RIGHT)
						}
					}

					'-' -> {
						val newPosLeft = pos.move(Direction.LEFT)
						if (isInBounds(grid, newPosLeft) && (!hasCompletelyVisited(grid, visited, newPosLeft) || grid[newPosLeft.y][newPosLeft.x] == '.')) {
							visited[newPosLeft.y][newPosLeft.x] = evaluateVis(grid, visited, newPosLeft, Direction.LEFT)
							queue.add(newPosLeft to Direction.LEFT)
						}
						val newPosRight = pos.move(Direction.RIGHT)
						if (isInBounds(grid, newPosRight) && (!hasCompletelyVisited(grid, visited, newPosRight) || grid[newPosRight.y][newPosRight.x] == '.')) {
							visited[newPosRight.y][newPosRight.x] = evaluateVis(grid, visited, newPosRight, Direction.RIGHT)
							queue.add(newPosRight to Direction.RIGHT)
						}
					}

					else -> {
						val newPos = pos.move(Direction.UP)
						if (isInBounds(grid, newPos) && (!hasCompletelyVisited(grid, visited, newPos) || grid[newPos.y][newPos.x] == '.')) {
							visited[newPos.y][newPos.x] = evaluateVis(grid, visited, newPos, Direction.UP)
							queue.add(newPos to Direction.UP)
						}
					}
				}
			}

			Direction.DOWN -> {
				when (currentGridType) {
					'\\' -> {
						val newPos = pos.move(Direction.RIGHT)
						if (isInBounds(grid, newPos) && (!hasCompletelyVisited(grid, visited, newPos) || grid[newPos.y][newPos.x] == '.')) {
							visited[newPos.y][newPos.x] = evaluateVis(grid, visited, newPos, Direction.RIGHT)
							queue.add(newPos to Direction.RIGHT)
						}
					}

					'/' -> {
						val newPos = pos.move(Direction.LEFT)
						if (isInBounds(grid, newPos) && (!hasCompletelyVisited(grid, visited, newPos) || grid[newPos.y][newPos.x] == '.')) {
							visited[newPos.y][newPos.x] = evaluateVis(grid, visited, newPos, Direction.LEFT)
							queue.add(newPos to Direction.LEFT)
						}
					}

					'-' -> {
						val newPosLeft = pos.move(Direction.LEFT)
						if (isInBounds(grid, newPosLeft) && (!hasCompletelyVisited(grid, visited, newPosLeft) || grid[newPosLeft.y][newPosLeft.x] == '.')) {
							visited[newPosLeft.y][newPosLeft.x] = evaluateVis(grid, visited, newPosLeft, Direction.LEFT)
							queue.add(newPosLeft to Direction.LEFT)
						}
						val newPosRight = pos.move(Direction.RIGHT)
						if (isInBounds(grid, newPosRight) && (!hasCompletelyVisited(grid, visited, newPosRight) || grid[newPosRight.y][newPosRight.x] == '.')) {
							visited[newPosRight.y][newPosRight.x] = evaluateVis(grid, visited, newPosRight, Direction.RIGHT)
							queue.add(newPosRight to Direction.RIGHT)
						}
					}

					else -> {
						val newPos = pos.move(Direction.DOWN)
						if (isInBounds(grid, newPos) && (!hasCompletelyVisited(grid, visited, newPos) || grid[newPos.y][newPos.x] == '.')) {
							visited[newPos.y][newPos.x] = evaluateVis(grid, visited, newPos, Direction.DOWN)
							queue.add(newPos to Direction.DOWN)
						}
					}
				}
			}

			Direction.LEFT -> {
				when (currentGridType) {
					'\\' -> {
						val newPos = pos.move(Direction.UP)
						if (isInBounds(grid, newPos) && (!hasCompletelyVisited(grid, visited, newPos) || grid[newPos.y][newPos.x] == '.')) {
							visited[newPos.y][newPos.x] = evaluateVis(grid, visited, newPos, Direction.UP)
							queue.add(newPos to Direction.UP)
						}
					}

					'/' -> {
						val newPos = pos.move(Direction.DOWN)
						if (isInBounds(grid, newPos) && (!hasCompletelyVisited(grid, visited, newPos) || grid[newPos.y][newPos.x] == '.')) {
							visited[newPos.y][newPos.x] = evaluateVis(grid, visited, newPos, Direction.DOWN)
							queue.add(pos.move(Direction.DOWN) to Direction.DOWN)
						}
					}

					'|' -> {
						val newPosUp = pos.move(Direction.UP)
						if (isInBounds(grid, newPosUp) && (!hasCompletelyVisited(grid, visited, newPosUp) || grid[newPosUp.y][newPosUp.x] == '.')) {
							visited[newPosUp.y][newPosUp.x] = evaluateVis(grid, visited, newPosUp, Direction.UP)
							queue.add(newPosUp to Direction.UP)
						}
						val newPosDown = pos.move(Direction.DOWN)
						if (isInBounds(grid, newPosDown) && (!hasCompletelyVisited(grid, visited, newPosDown) || grid[newPosDown.y][newPosDown.x] == '.')) {
							visited[newPosDown.y][newPosDown.x] = evaluateVis(grid, visited, newPosDown, Direction.DOWN)
							queue.add(newPosDown to Direction.DOWN)
						}
					}

					else -> {
						val newPos = pos.move(Direction.LEFT)
						if (isInBounds(grid, newPos) && (!hasCompletelyVisited(grid, visited, newPos) || grid[newPos.y][newPos.x] == '.')) {
							visited[newPos.y][newPos.x] = evaluateVis(grid, visited, newPos, Direction.LEFT)
							queue.add(newPos to Direction.LEFT)
						}
					}
				}
			}

			Direction.RIGHT -> {
				when (currentGridType) {
					'\\' -> {
						val newPos = pos.move(Direction.DOWN)
						if (isInBounds(grid, newPos) && (!hasCompletelyVisited(grid, visited, newPos) || grid[newPos.y][newPos.x] == '.')) {
							visited[newPos.y][newPos.x] = evaluateVis(grid, visited, newPos, Direction.DOWN)
							queue.add(newPos to Direction.DOWN)
						}
					}

					'/' -> {
						val newPos = pos.move(Direction.UP)
						if (isInBounds(grid, newPos) && (!hasCompletelyVisited(grid, visited, newPos) || grid[newPos.y][newPos.x] == '.')) {
							visited[newPos.y][newPos.x] = evaluateVis(grid, visited, newPos, Direction.UP)
							queue.add(newPos to Direction.UP)
						}
					}

					'|' -> {
						val newPosUp = pos.move(Direction.UP)
						if (isInBounds(grid, newPosUp) && (!hasCompletelyVisited(grid, visited, newPosUp) || grid[newPosUp.y][newPosUp.x] == '.')) {
							visited[newPosUp.y][newPosUp.x] = evaluateVis(grid, visited, newPosUp, Direction.UP)
							queue.add(newPosUp to Direction.UP)
						}
						val newPosDown = pos.move(Direction.DOWN)
						if (isInBounds(grid, newPosDown) && (!hasCompletelyVisited(grid, visited, newPosDown) || grid[newPosDown.y][newPosDown.x] == '.')) {
							visited[newPosDown.y][newPosDown.x] = evaluateVis(grid, visited, newPosDown, Direction.DOWN)
							queue.add(newPosDown to Direction.DOWN)
						}
					}

					else -> {
						val newPos = pos.move(Direction.RIGHT)
						if (isInBounds(grid, newPos) && (!hasCompletelyVisited(grid, visited, newPos) || grid[newPos.y][newPos.x] == '.')) {
							visited[newPos.y][newPos.x] = evaluateVis(grid, visited, newPos, Direction.RIGHT)
							queue.add(newPos to Direction.RIGHT)
						}
					}
				}
			}
		}
	}
	var ans = 0
	visited.forEach { row ->
		row.forEach {
			ans += if (it.second == 1 || it.first == 1) 1 else 0
		}
		//println(row.map { if (it.second == 1 || it.first == 1) '#' else '.' }.joinToString(""))
	}
	//println("$ans ${sPosition.x} ${sPosition.y} ${grid[sPosition.y][sPosition.x]} ${initialDirection}")
	return ans
}

fun main() {
	val grid = generateSequence(::readLine).toList()
	println(getEnergizedTileCount(grid, Position(0, 0), Direction.RIGHT))
	var ans = -1
	for (x in grid[0].indices) {
		ans = max(ans, getEnergizedTileCount(grid, Position(x, 0), Direction.DOWN))
		ans = max(ans, getEnergizedTileCount(grid, Position(x, grid.size - 1), Direction.UP))
	}
	for (y in grid.indices) {
		ans = max(ans, getEnergizedTileCount(grid, Position(0, y), Direction.RIGHT))
		ans = max(ans, getEnergizedTileCount(grid, Position(grid[0].length - 1, y), Direction.LEFT))
	}
	println(ans)
}