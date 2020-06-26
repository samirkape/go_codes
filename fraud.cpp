#include <bits/stdc++.h>

using namespace std;

vector<string> split_string(string);

#define RANGE 201

float even( int * count,int *arr, int n ){
    int mid = (n / 2) - 1;
    int tmp1 = 0;
    int tmp2 = 0;
    int m1 = 0;
    int m2 = 0;
    int i = 0;
    for(i = 0; i < RANGE; ++i){
        tmp1 += count[i];
        if(tmp1-1 >= mid){
            if( tmp1 - 1 > mid ){
                m1 = m2 = i;
                break;
            }
            else{
                m1 = i;
                while(1){
                    if(count[++i]){
                        m2 = i;
                        return ( m1 + m2 ) / 2.0;
                    }
                }
            }
        }
    }
    return ( m1 + m2 ) / 2.0;
}
float odd( int * count,int * arr, int mid ){
    int tmp1 = 0;
    int tmp2 = 0;
    int i = 0;
    for(i = 0; i < RANGE; ++i){
        if(tmp1 > mid) break;
        if( count[i] ) { 
            tmp1 += count[i];
        }
    }
    return i-1;
}

float med(int * arr, int * count, int d) {

    int size = d;
    int mid = (d / 2);
    float out = 0.0;

    if( (char)size & 1 )
        out = odd( count, arr, mid );
    else
        out = even( count, arr, size );
    return out;
}
// Complete the activityNotifications function below.
int activityNotifications(vector<int> expenditure, int d) {

    int * arr = &expenditure[0];
    int ct = 0;
    int count[RANGE + 1] = {0};

    for(int i = 0; i < d; ++i)  
        ++count[expenditure[i]]; 

    for( int i = 0; (i + d) < expenditure.size(); i++ ){
        float median = med( arr+i, count, d );
        if ( ( expenditure[d+i] ) >= (2 * median) )
            ct+=1;
        count[arr[i]] -= 1;
        count[arr[i+d]] += 1;
    }
    return ct; 
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string nd_temp;



    ifstream myfile ("input01.txt");
    //getline (myfile,nd_temp);
    getline (cin,nd_temp);


    vector<string> nd = split_string(nd_temp);

    int n = stoi(nd[0]);

    int d = stoi(nd[1]);

    string expenditure_temp_temp;
    getline(cin,expenditure_temp_temp);
    //getline(myfile, expenditure_temp_temp);

    vector<string> expenditure_temp = split_string(expenditure_temp_temp);

    vector<int> expenditure(n);

    for (int i = 0; i < n; i++) {
        int expenditure_item = stoi(expenditure_temp[i]);

        expenditure[i] = expenditure_item;
    }
    int result = activityNotifications(expenditure, d);

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
