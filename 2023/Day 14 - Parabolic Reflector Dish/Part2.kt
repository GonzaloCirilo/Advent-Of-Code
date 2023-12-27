fun rotate(dish: MutableList<MutableList<Char>>): MutableList<MutableList<Char>> {
    val aux = mutableListOf<MutableList<Char>>()
    for (j in dish[0].indices) {
        aux.add(mutableListOf())
        for (i in dish.indices.reversed()) {
            aux[j].add(dish[i][j])
        }
    }
    return aux
}

fun bringRocksToTop(dish: MutableList<MutableList<Char>>): MutableList<MutableList<Char>> {
    var lastBlock = 0
    var rockCount = 0
    dish.add(dish[0].map { '#' }.toMutableList())
    for (j in dish[0].indices) {
        lastBlock = 0
        for (i in dish.indices) {
            if (dish[i][j] == '#') {
                for (k in lastBlock until i) {
                    dish[k][j] = if (rockCount > 0) {
                        'O'
                    } else {
                        '.'
                    }
                    rockCount--
                }
                rockCount = 0
                lastBlock = i + 1
            }
            if (dish[i][j] == 'O') {
                rockCount++
            }
        }
    }
    //dish.removeLast()
    dish.removeAt(dish.size - 1)
    return dish
}

fun cycle(dish: MutableList<MutableList<Char>>): MutableList<MutableList<Char>> {
    var aux = dish
    for (i in 0 until 4) {
        aux = bringRocksToTop(aux)
        aux = rotate(aux)
    }
    return aux
}

fun evaluateLoad(dish: MutableList<MutableList<Char>>): Int {
    var load = 0
    for (j in dish[0].indices) {
        for (i in dish.indices) {
            if (dish[i][j] == 'O') {
                load += (dish.size - i)
            }
        }
    }
    return load
}

fun main() {
    var ans = 0
    var dishOriginal =
        generateSequence(::readLine).toMutableList().map { it.toMutableList() }.toMutableList()
    var dish = dishOriginal
    val dict = mutableMapOf<String, Int>()
    for(x in 1..5000) {
        dish = cycle(dish)
        val key = dish.joinToString("") { it.joinToString("") }
        if (!dict.containsKey(key)) {
            dict[key] = 1
        } else {
            dict[key]?.let {
                dict[key] = it + 1
            }
        }
    }
    val threshold = dict.maxOf { it.value }
    var cycleSize = 0
    for ((k, v) in  dict) {
        //println("$v $k")
        if (v == threshold || v == threshold - 1) {
            cycleSize++
        }
    }
    val totalToEvaluate = ((1000000000 - (dict.size - cycleSize)) % cycleSize) + (dict.size - cycleSize)
    //println("cycle size: $cycleSize evaluate only $totalToEvaluate")

    for (x in 1..totalToEvaluate) {
        dishOriginal = cycle(dishOriginal)
    }
    ans = evaluateLoad(dishOriginal)
    println(ans)
}