fun main() {
    var a = 5
    var b = 10

    while (a != 0 && b != 0) {
        if (a > b) {
            a = a % b
        } else {
            b = b % a
        }
    }
    println(a+b)
}