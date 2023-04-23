import kotlin.system.measureTimeMillis
import kotlin.time.measureTime

val foods = listOf<String> ("Sate", "Bakso", "Dimsum", "Rames")
val drinks = listOf<String> ("Jeruk", "Teh", "Kelapa", "Cendol")

fun logPairs(array1: List<String>, array2: List<String>, words: String) {
    var counter = 0
    for (item1 in array1) {
        for (item2 in array2) {
            counter++
            println("$words $counter $item1 dan $item2")
        }
    }
}

fun main() {
   logPairs(foods, drinks, "Menu")
}