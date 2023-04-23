import kotlin.random.Random

val telco = listOf("Telkom", "Indosat", "XL", "Verizon", "AT&T", "Nippon", "Vodafone", "Orange", "KDDI", "Telefonica", "T-Mobile")

fun findCompany(array: List<String>, input: Int) {
    for (i in array.indices) {
        if (array[i] == array[input]) {
            println("Company Found ${array[input]}")
            break
        }
        println("Searching Company...${i+1}")
    }
}

fun main() {
    val company = Random.nextInt(telco.size)
    findCompany(telco, company)
}
