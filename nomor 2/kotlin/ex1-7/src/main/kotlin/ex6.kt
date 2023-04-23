import kotlin.system.measureTimeMillis
import kotlin.time.measureTime

val address2 = "DKI Jakarta"

val addresses2 = MutableList(100000){ address2 }

fun findAddress2(addresses: MutableList<String>) {
    val default = addresses.get(0)
    println("The default Address is $default")
}
fun main() {
    val time = measureTimeMillis {
        findAddress2(addresses2)
    }
    println("Time taken to process code is $time ms")
}