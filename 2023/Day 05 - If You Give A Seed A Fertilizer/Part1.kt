import kotlin.math.min

fun getNewTargets(ranges: List<Triple<Long, Long, Long>>, targetList: List<Long>): List<Long> {
	return targetList.map { t ->
		ranges.forEach { (start, end, v) ->
			if (t in start..end) {
				return@map (t - start) + v
			}
		}
		t
	}
}

fun main() {
	val input = generateSequence(::readLine).toList()
	var targetList =
		input[0].split(":")[1].trim().split(" ").filter { it.isNotEmpty() }.map { it.toLong() }
	var ranges = mutableListOf<Triple<Long, Long, Long>>()
	input.subList(2, input.size).forEach { s ->
		if (s.isEmpty()) {
			targetList = getNewTargets(ranges, targetList)
			ranges = mutableListOf()
			return@forEach
		}
		if (!s[0].isDigit()) return@forEach
		val (v, start, l) = s.split(" ").map { it.toLong() }
		ranges.add(Triple(start, start + l - 1, v))
	}
	targetList = getNewTargets(ranges, targetList)
	println(targetList.min())
}