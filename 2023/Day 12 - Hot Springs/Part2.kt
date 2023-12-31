import kotlin.math.min
var hashMap: HashMap<String, Int> = hashMapOf()
var minLength = 0

const val UNDEFINED = (-1e9).toInt()
fun main() {
	var ans = 0L
	generateSequence(::readLine).toList().forEach {
		val tokens = it.split(" ")
		val initialSpring = tokens[0]
		val configuration = mutableListOf<Int>()
		var springs = initialSpring
		repeat(5){
			tokens[1].split(",").map { s -> s.toInt() }.forEach { x ->
				configuration.add(UNDEFINED)
				configuration.add(x)
			}
		}
		configuration.add(UNDEFINED)
		repeat(4) {
			springs += "?$initialSpring"
		}
		springs = ".$springs."
		//println("$springs $configuration")
		val n: Int = springs.length
		val m = configuration.size
		val dp = Array(n + 1) { LongArray(m + 1) }
		dp[n][m] = 1

		for (i in n - 1 downTo 0) {
			for (j in m - 1 downTo 0) {
				dp[i][j] = if (configuration[j] > 0) {
					val end: Int = i + configuration[j]
					if (!springs.subSequence(i, min(end, springs.length)).contains('.'))
						dp[end][j + 1]
					else
						dp[i][j]
				} else {
					if (springs[i] != '#')
						dp[i + 1][j + 1] + dp[i + 1][j]
					else
						dp[i][j]
				}
			}
		}
		//println(dp[0][0])
		ans += dp[0][0]
	}
	println(ans)
}