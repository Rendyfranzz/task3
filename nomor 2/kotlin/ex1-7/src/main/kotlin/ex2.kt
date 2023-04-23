import kotlin.system.measureTimeMillis
import kotlin.time.measureTime

val company1 = "telkom"

val largeCompanyName1 = MutableList(1000){ company1 }

fun findCompany1(array: MutableList<String>) {
    for (item in array) {
        if (item == "telkom") {
            println("Company Found!")
        }
    }
}
fun main() {
    val time = measureTimeMillis {
        findCompany1(largeCompanyName1)
    }
    findCompany1(largeCompanyName1)
    println("Time taken to process code is $time ms")
}