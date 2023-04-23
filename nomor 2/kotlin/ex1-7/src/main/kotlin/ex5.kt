import kotlin.system.measureTimeMillis
import kotlin.time.measureTime

val address1 = "DKI Jakarta"

val addresses1 = MutableList(1000){ address1 }

fun findAddress1(addresses: MutableList<String>) {
    val default = addresses.get(0)
    println("The default Address is $default")
}
fun main() {
    val time = measureTimeMillis {
        findAddress1(addresses1)
    }
    println("Time taken to process code is $time ms")
}