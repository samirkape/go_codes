#include <bits/stdc++.h>

using namespace std;

vector<string> split_string(string);
void swap(int *xp, int *yp) 
{ 
    int temp = *xp; 
    *xp = *yp; 
    *yp = temp; 
} 
  
void  sort_(int T[],int L[], int n) 
{ 
   int i, j; 
   for (i = 0; i < n-1; i++)       
       for (j = 0; j < n-i-1; j+=2)  
           if (L[j] > L[j+1] && T[j]){
              swap(&L[j], &L[j+1]); 
              //swap(&T[j], &T[j+1]);
           }

} 


// Complete the luckBalance function below.
int luckBalance(int k, vector<vector<int>> contests) {
    vector<int> tmp;
    int sum = 0;
    for(long int i = 0; i < contests.size(); i++ ){
        for(long int j = 0; j < contests[i].size(); j+=2 ){
            if(contests[i][1])
                tmp.push_back(contests[i][0]);
                sum += contests[i][0];
        }
    }
    sort(tmp.begin(),tmp.end());
    if(tmp.size() < k ) return sum;
    for(long int i = 0; i < tmp.size()-k; i++ ){
        sum -= (tmp[i] + tmp[i]);
    }
    return sum;
    return 0;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string nk_temp;
    getline(cin, nk_temp);

    vector<string> nk = split_string(nk_temp);

    int n = stoi(nk[0]);

    int k = stoi(nk[1]);

    vector<vector<int>> contests(n);
    for (int i = 0; i < n; i++) {
        contests[i].resize(2);

        for (int j = 0; j < 2; j++) {
            cin >> contests[i][j];
        }

        cin.ignore(numeric_limits<streamsize>::max(), '\n');
    }

    int result = luckBalance(k, contests);

    fout << result << "\n";

    fout.close();

    return 0;
}

vector<string> split_string(string input_string) {
    string::iterator new_end = unique(input_string.begin(), input_string.end(), [] (const char &x, const char &y) {
        return x == y and x == ' ';
    });

    input_string.erase(new_end, input_string.end());

    while (input_string[input_string.length() - 1] == ' ') {
        input_string.pop_back();
    }

    vector<string> splits;
    char delimiter = ' ';

    size_t i = 0;
    size_t pos = input_string.find(delimiter);

    while (pos != string::npos) {
        splits.push_back(input_string.substr(i, pos - i));

        i = pos + 1;
        pos = input_string.find(delimiter, i);
    }

    splits.push_back(input_string.substr(i, min(pos, input_string.length()) - i + 1));

    return splits;
}
