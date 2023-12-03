import java.math.BigInteger
import kotlin.math.max

fun main() {
	val input = generateSequence(::readLine)
	var ans: BigInteger = BigInteger.ZERO
	input.toList().forEach{ s ->
		val gameTries = s.split(":")[1].split(";")
		var maxr = 0L
		var maxg = 0L
		var maxb = 0L
		gameTries.forEach {  gTry->
			gTry.split(",").forEach { c ->
				val cube = c.trim()
				val aux = cube.split(" ")
				val cubeCount = aux[0].toLong()
				when(aux[1]) {
					"green" -> maxg = max(maxg, cubeCount)
					"red" -> maxr = max(maxr, cubeCount)
					"blue" -> maxb = max(maxb, cubeCount)
					else -> 0
				}
			}
		}
		ans += BigInteger.valueOf(maxr) * BigInteger.valueOf(maxb) * BigInteger.valueOf(maxg)
	}
	print(ans)
}