#include <iostream>
#include <bits/stdc++.h>
#include <vector>
#include <set>
#include <iterator> 
using namespace std;
vector<string> split_string(string);
int inc_count(int *brr, int d, int mid, int val)
{
    int m = 0, v1 = 0, v2 = 0, count = 0;
    
    for(int i = 0; i < 201; ++i)
    {
        if(brr[i])
        {
            count += brr[i];
        }
        if(count-1 >= mid)
        {
            v1 = i;
            if(count-1 > mid)
            {
                v2 = i;
                break;
            }
            else
            {
                for(int j = i+1; j < 201; ++j)
                {
                    if(brr[j])
                    {
                        v2 = j;
                        break;
                    }
                }
            }
            break;
        }
    }

    if(d % 2)
    {
        if(v2 * 2 <= val)
            return 1;
    }
    else
    {
        m = v1 + v2;
        if(m <= val)
            return 1;
    }

    return 0;
}

int notification_count(vector<int> arr, int n, int d)
{
    int count = 0, mid = (d / 2 - 1);
    int brr[201] = {0};

    for(int i = 0; i < d; ++i)
    {
        brr[arr[i]] += 1;
    }

    for(int i = d; i < n; ++i)
    {
        count += inc_count(brr, d, mid, arr[i]);

        brr[arr[i-d]] -= 1;
        brr[arr[i]] += 1;
    }

    return count;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string nd_temp;



    ifstream myfile ("input01.txt");
    getline (myfile,nd_temp);
    //getline (cin,nd_temp);


    vector<string> nd = split_string(nd_temp);

    int n = stoi(nd[0]);

    int d = stoi(nd[1]);

    string expenditure_temp_temp;
    //getline(cin,expenditure_temp_temp);
    getline(myfile, expenditure_temp_temp);

    vector<string> expenditure_temp = split_string(expenditure_temp_temp);

    vector<int> expenditure(n);

    for (int i = 0; i < n; i++) {
        int expenditure_item = stoi(expenditure_temp[i]);

        expenditure[i] = expenditure_item;
    }
    int result = notification_count(expenditure, expenditure.size(), d);

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
