var a = 5;
var b = 1;
function nwd(a,b) {
    while (a != 0 && b != 0) {
        if (a > b) {
            a = a % b;
        } else {
            b = b % a;
        }
    }
    return (a + b);
}



console.log(nwd(a,b));
