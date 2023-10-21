#include <iostream>
using namespace std;
int nwd(int a, int b) {
    while (a != 0 && b != 0) {
        if (a > b) {
            a = a % b;
        } else {
            b = b % a;
        }
    }
    return (a + b);
}
int a = 5;
int b = 5;
int main() {
    cout << nwd(a,b);
    return 0;
}