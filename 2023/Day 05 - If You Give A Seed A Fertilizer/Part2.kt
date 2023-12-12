import kotlin.math.min
import kotlin.math.max
import kotlin.collections.*

fun main() {
	val input = generateSequence(::readLine).toList()
	var targetList =
			input[0].split(":")[1].trim().split(" ").filter { it.isNotEmpty() }.map { it.toLong() }
					.chunked(2).map {
						Pair(it[0], it[0] + it[1])
					}.toMutableList()
	var ranges = mutableListOf<Triple<Long, Long, Long>>()
	input.drop(2).forEach { line ->
		if (line.isEmpty()) {
			var newTargetList = emptyList<Pair<Long, Long>>().toMutableList()
			ranges = ranges
			while (targetList.isNotEmpty()) {
				val (s, e) = targetList.last()
				targetList.removeAt(targetList.size - 1)
				var inRange = false
				for ((rs, re, v) in ranges) {
					var auxS = max(s, rs)
					var auxE = min(e, re)
					if (auxS < auxE) {
						newTargetList.add(Pair(auxS - rs + v, auxE - rs + v))
						if (auxS > s) {
							newTargetList.add(Pair(s, auxS))
						}
						if (e > auxE) {
							newTargetList.add(Pair(auxE, e))
						}
						inRange = true
						break
					}
				}
				if (!inRange) {
					newTargetList.add(Pair(s, e))
				}
			}
			targetList = newTargetList
			ranges = mutableListOf()
			return@forEach
		}
		if (!line[0].isDigit()) return@forEach
		val (v, start, l) = line.split(" ").map { it.toLong() }
		ranges.add(Triple(start, start + l, v))
	}
	println(targetList.filter { it.first != 0L }.sortedBy { it.first }[0].first)

}