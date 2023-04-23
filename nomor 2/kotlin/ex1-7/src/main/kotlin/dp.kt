var counterPlainRecursive = 0
var counterDynamicProgramming = 0
fun fibonacciPlainRecursive(n: Int): Int {
    counterPlainRecursive++
    return if (n < 2) {
        n
    } else {
        fibonacciPlainRecursive(n - 1) + fibonacciPlainRecursive(n - 2)
    }
}
fun fibonacciDynamicProgramming(): (Int) -> Int {
    val cache = mutableMapOf<Int, Int>()
    fun fib(n: Int): Int {
        counterDynamicProgramming++
        if (n in cache) {
            return cache[n]!!
        } else {
            val result = if (n < 2) {
                n
            } else {
                fib(n - 1) + fib(n - 2)
            }
            cache[n] = result
            return cache[n]!!
        }
    }
    return ::fib
}

fun main(){
    fibonacciPlainRecursive(20)
    val fasterFibonacci = fibonacciDynamicProgramming()
    fasterFibonacci(20)
    println("we did $counterPlainRecursive calculations for Plain Recursive")
    println("we did $counterDynamicProgramming calculations for Dynamic Programming")
}