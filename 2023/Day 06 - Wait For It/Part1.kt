import kotlin.math.ceil
import kotlin.math.floor
import kotlin.math.pow
import kotlin.math.sqrt

fun main() {
	val times = readLine()?.split(":")?.get(1)?.split(" ")?.filter { it.toIntOrNull() != null }?.map { it.toInt() }
			?: emptyList()
	val distances = readLine()?.split(":")?.get(1)?.split(" ")?.filter { it.toIntOrNull() != null }?.map { it.toInt() }
			?: emptyList()
	var ans = 1
	times.forEachIndexed { i, t ->
		val distance = distances[i]
		val l = floor((t - sqrt(t.toDouble().pow(2.toDouble()) -4*distance)) /2).toInt()
		val r = ceil((t + sqrt(t.toDouble().pow(2.toDouble())-4*distance)) /2).toInt()
		println("$l $r")
		println((r - l) - 1)
		ans *= (r - l) - 1
	}
	println(ans)
}