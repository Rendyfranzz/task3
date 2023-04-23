import kotlin.system.measureTimeMillis
import kotlin.time.measureTime

val company2 = "telkom"

val largeCompanyName2 = MutableList(100000){ company2 }

fun findCompany2(array: MutableList<String>) {
    for (item in array) {
        if (item == "telkom") {
            println("Company Found!")
        }
    }
}
fun main() {
    val time = measureTimeMillis {
        findCompany2(largeCompanyName2)
    }
    findCompany2(largeCompanyName2)
    println("Time taken to process code is $time ms")
}