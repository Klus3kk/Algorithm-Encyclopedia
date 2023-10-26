bool Anagram(char str1, char str2) {
    char temp;
    int i, j;
    int n = strlen(str1);
    int n1 = strlen(str2);

    if (n != n1) {
        return 0;
    }
    str1.sort()
    str2.sort()
    for (i = 0; i < n; i++) {
        if(s1[i] == s2[i]) {
            return 1;
        }
    }    
    return 0;
}