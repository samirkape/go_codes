#include <bits/stdc++.h>

using namespace std;

// Complete the findDigits function below.
vector<int> split_int(int n){
    vector<int> digits;
    int i = 0;
    while(n){
        //digits[i] =  n % 10;
        digits.push_back( n % 10 );
        n /= 10;
    }
    return digits;
}

int findDigits(int n) {
    vector<int> digit_bank = split_int(n);
    int count = 0;
    for( int i = 0; i < digit_bank.size(); i++ ){
        if(digit_bank[i] == 0) continue;
        if( n % digit_bank[i] == 0 ) count++;
    }
    return count;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int t;
    cin >> t;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    for (int t_itr = 0; t_itr < t; t_itr++) {
        int n;
        cin >> n;
        cin.ignore(numeric_limits<streamsize>::max(), '\n');

        int result = findDigits(n);

        fout << result << "\n";
    }

    fout.close();

    return 0;
}
