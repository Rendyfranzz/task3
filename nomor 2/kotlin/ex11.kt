import kotlin.math.log10
import kotlin.math.pow

fun generateData(): List<String> {
    val number = "0821234567"
    val customers = mutableListOf<String>()
    for (i in 0..99) {
        val mobileNumber = if (i < 10) number else "$number${i}"
        customers.add(mobileNumber)
    }
    return customers
}

fun sendPromoDiscount(array: List<String>) {
    for (mobileNumber in array) {
        println("Sending Promo to $mobileNumber")
    }
    for (i in 1..10) {
        val chosenCustomer = log10(i.toDouble() + 1).pow(3).toInt() % array.size
        println("Sending Discount to Chosen Customer ${chosenCustomer + 1}")
    }
}

fun main() {
    val customers = generateData()
    sendPromoDiscount(customers)
}
