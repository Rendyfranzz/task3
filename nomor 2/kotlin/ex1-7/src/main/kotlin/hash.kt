class HashTable(size: Int) {
    private val data = arrayOfNulls<Array<Pair<String, Any>>>(size)

    private fun hash(key: String): Int {
        var hash = 0
        for (i in key.indices) {
            hash = (hash + key[i].code * i) % data.size
        }
        return hash
    }

    operator fun set(key: String, value: Any) {
        val address = hash(key)
        if (data[address] == null) {
            data[address] = emptyArray()
        }
        data[address] = data[address]!! + arrayOf(key to value)
    }

    operator fun get(key: String): Any? {
        val address = hash(key)
        val currentBucket = data[address]
        if (currentBucket != null) {
            for (i in currentBucket.indices) {
                if (currentBucket[i].first == key) {
                    return currentBucket[i].second
                }
            }
        }
        return null
    }
}

fun main(){
    val myHashTable = HashTable(100)
    myHashTable["082124606606"] = "Rony Setyawan"
    println(myHashTable["082124606606"])
}