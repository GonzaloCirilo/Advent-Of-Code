fun isAllZeroes(arr: List<Long>): Boolean {
    return arr.reduce { acc, x -> acc + x} == 0L
}

fun solve(arr: List<Long>): Long {
    if (isAllZeroes(arr)) return arr[0]
    val nArr = mutableListOf<Long>()
    arr.dropLast(1).forEachIndexed { i, _ ->
        nArr.add(arr[i + 1] - arr[i])
    }
    return solve(nArr) + arr.last()
}


fun main() {
    val ans = generateSequence(::readLine).toList().map { it.split(" ").map { x -> x.toLong() } }.map {  series ->
        solve(series)
    }.reduce { acc, s -> acc + s}
    println(ans)
}