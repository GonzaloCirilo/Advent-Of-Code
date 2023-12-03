
fun main() {
    val lr = 12
    val lg = 13
    val lb = 14
	val input = generateSequence(::readLine)
	var ans = 0
	input.toList().forEachIndexed { i, s ->
		val gameTries = s.split(":")[1].split(";")

		gameTries.forEach {  gTry->
			gTry.split(",").forEach { c ->
				val cube = c.trim()
				val aux = cube.split(" ")
				val cubeCount = aux[0].toInt()
				val compareTo = when(aux[1]) {
					"green" -> lg
					"red" -> lr
					"blue" -> lb
					else -> 0
				}
				if (cubeCount > compareTo) return@forEachIndexed
			}
		}
		ans += (i + 1)
	}
	println(ans.toString())
}