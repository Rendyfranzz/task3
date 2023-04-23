import kotlin.system.measureTimeMillis
import kotlin.time.measureTime

val company = "telkom"

val largeCompanyName = MutableList(100){ company }

fun findCompany(array: MutableList<String>) {
    for (item in array) {
        if (item == "telkom") {
            println("Company Found!")
        }
    }
}
fun main() {
    val time = measureTimeMillis {
        findCompany(largeCompanyName)
    }
    findCompany(largeCompanyName)
    println("Time taken to process code is $time ms")
}
