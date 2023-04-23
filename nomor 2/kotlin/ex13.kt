package main

import java.lang.StringBuilder
import kotlin.random.Random

fun generateData(n: Int): List<String> {
    val baseNumber = "882"
    val customers = mutableListOf<String>()
    val sb = StringBuilder()

    for (i in 0 until n) {
        sb.clear()
        sb.append(baseNumber)
        sb.append("0".repeat(9))
        sb.append(i)
        customers.add(sb.toString())
    }

    return customers
}

fun sendPromoDiscount(input1: List<String>, input2: List<String>) {
    for (i in input1.indices) {
        println("Sending Promo to ${input1[i]}")
    }
    println("Its Finished to send Promo to ${input1.size} Customers")
    for (i in input2.indices) {
        println("Sending Discount to ${input2[i]}")
    }
    println("Its Finished to send Discount to ${input2.size} Customers")
}

fun main() {
    val dataPromo = generateData(100000000)
    val dataDiscount = generateData(1000)
    sendPromoDiscount(dataPromo, dataDiscount)
}
