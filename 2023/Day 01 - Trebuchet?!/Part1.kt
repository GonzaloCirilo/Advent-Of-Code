
fun main() {
    val input = generateSequence(::readLine)
    val ans = input.toList().map { line ->
        val d = line.first{ it.isDigit() }.digitToInt()
        val u = line.last { it.isDigit() }.digitToInt()
        d * 10 + u
    }.reduce { acc, i ->
        acc + i
    }
    println(ans)
}