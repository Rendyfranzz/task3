import kotlin.system.measureTimeMillis
import kotlin.time.measureTime

val address = "DKI Jakarta"

val addresses = MutableList(10){ address }

fun findAddress(addresses: MutableList<String>) {
    val default = addresses.get(0)
    println("The default Address is $default")
}
fun main() {
    val time = measureTimeMillis {
        findAddress(addresses)
    }
    println("Time taken to process code is $time ms")
}