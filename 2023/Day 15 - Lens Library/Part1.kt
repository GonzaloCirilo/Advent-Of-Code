fun hash(s: String): Int {
	var value = 0
	for (c in s) {
		value += c.toInt()
		value *= 17
		value %= 256
	}
	return value
}

fun main() { 
	println(readln().split(",").map{ hash(it) }.reduce { acc, i -> acc + i })
}