fun generateData(n: Int): List<String> {
    val baseNumber = "882"
    val customers = mutableListOf<String>()
    for (i in 0 until n) {
        val mobileNumber = "$baseNumber${"0".repeat(9)}$i".takeLast(10)
        customers.add(mobileNumber)
    }
    return customers
}

fun sendPromoDiscount(input: List<String>) {
    for (mobileNumber in input) {
        println("Sending Promo to $mobileNumber")
    }
    println("Its Finished to send Promo to ${input.size} Customers")
    for (mobileNumber in input) {
        println("Sending Discount to $mobileNumber")
    }
    println("Its Finished to send Discount to ${input.size} Customers")
}

fun main() {
    val data = generateData(10)
    sendPromoDiscount(data)
}
