fun arrange(array: List<String>, memory: List<String>? = null): List<List<String>> {
    val results = mutableListOf<List<String>>()
    val tempMemory = memory ?: emptyList()
    for (i in array.indices) {
        val current = listOf(array[i])
        if (array.size == 1) {
            results.add(tempMemory + current)
        } else {
            val subarray = array.filterIndexed { index, _ -> index != i }
            val newMemory = tempMemory + current
            results.addAll(arrange(subarray, newMemory))
        }
    }
    return results
}

fun main() {
    val candidates = listOf("Baswedan", "Subianto", "Maharani","Ganjar")
    var chairs = arrange(candidates)
    for (chair in chairs) {
        println(chair)
    }
    
}
