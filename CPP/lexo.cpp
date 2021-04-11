#include <bits/stdc++.h>

using namespace std;

int bsearch(string& s, int l, int r, int key) 
{ 
    int index = -1; 
    while (l <= r) { 
        int mid = l + (r - l) / 2; 
        if (s[mid] <= key) 
            r = mid - 1; 
        else { 
            l = mid + 1; 
            if (index == -1 || s[index] >= s[mid]) 
                index = mid; 
        } 
    } 
    return index; 
} 

void swap( char *a, char *b ){
    if( *a== *b ) return;
    char tmp = *a;
    *a = *b;
    *b = tmp;
}
void rev(string& s, int l, int r) 
{ 
    while (l < r) 
        swap(&s[l++], &s[r--]); 
} 
// Complete the biggerIsGreater function below.
string biggerIsGreater(string w) {

    int len = w.length();
    string cp = w;
    int tmplen = len;
    int i = len - 2;
    auto itb = w.begin();
    auto ite = w.end();
    int pos = 0;

    while( i >= 0 && ( w[i] > w[i+1] || w[i] == w[i+1] ))
        i--; 
    
    if( i < 0 ) return "no answer";
    
    pos = bsearch(w, i+1, len-1, w[i]);
    if( pos == -1 ) return "no answer";
    swap( &w[i], &w[pos] );
    //sort( itb + i + 1, ite );
    rev(w, i+1, len-1);
    if(cp == w) return "no answer";
    return w;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int T;
    cin >> T;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    for (int T_itr = 0; T_itr < T; T_itr++) {
        string w;
        getline(cin, w);

        string result = biggerIsGreater(w);

        fout << result << "\n";
    }

    fout.close();

    return 0;
}
