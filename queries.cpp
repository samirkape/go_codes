#include <bits/stdc++.h>

using namespace std;

string ltrim(const string &);
string rtrim(const string &);
vector<string> split(const string &);

/*
 * Complete the 'getMaxCharCount' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. 2D_INTEGER_ARRAY queries
 */

int findMax(string s){
    int ct = 0;
    char * sarr = &s[0];
    char cmax = s[0];
    while(*sarr){
        if(*sarr!=cmax)
            break;
        ct++;
        sarr++;
    }
    return ct;
}

int sort_(char * s, int end){
    string str(s,end+1);
    transform(str.begin(), str.end(), str.begin(),
    [](unsigned char c){ return std::tolower(c); });
    sort(str.begin(),str.end(),greater<int>());
    return findMax(str);
}

vector<int> getMaxCharCount(string s, vector<vector<int>> queries) {
    // queries is a n x 2 array where queries[i][0] and queries[i][1] represents x[i] and y[i] for the ith query.
    char * sarr = &s[0];
    for(int i=0; i<queries.size(); i++){
        for(int j=0; j<queries[i].size(); j+=2 ){
            int op = sort_(sarr+queries[i][j],queries[i][j+1] - queries[i][j]);
            cout<<op<<" ";
        }
         cout<<"\n";
    }
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string n_temp;
    getline(cin, n_temp);

    int n = stoi(ltrim(rtrim(n_temp)));

    string s;
    getline(cin, s);
    n = s.size();

    string q_temp;
    getline(cin, q_temp);

    int q = stoi(ltrim(rtrim(q_temp)));

    vector<vector<int>> query(q);

    for (int i = 0; i < q; i++) {
        query[i].resize(2);

        string query_row_temp_temp;
        getline(cin, query_row_temp_temp);

        vector<string> query_row_temp = split(rtrim(query_row_temp_temp));

        for (int j = 0; j < 2; j++) {
            int query_row_item = stoi(query_row_temp[j]);

            query[i][j] = query_row_item;
        }
    }

    vector<int> ans = getMaxCharCount(s, query);

    for (int i = 0; i < (int)ans.size(); i++) {
        fout << ans[i];

        if (i != (int)ans.size() - 1) {
            fout << "\n";
        }
    }

    fout << "\n";

    fout.close();

    return 0;
}

string ltrim(const string &str) {
    string s(str);

    s.erase(
        s.begin(),
        find_if(s.begin(), s.end(), not1(ptr_fun<int, int>(isspace)))
    );

    return s;
}

string rtrim(const string &str) {
    string s(str);

    s.erase(
        find_if(s.rbegin(), s.rend(), not1(ptr_fun<int, int>(isspace))).base(),
        s.end()
    );

    return s;
}

vector<string> split(const string &str) {
    vector<string> tokens;

    string::size_type start = 0;
    string::size_type end = 0;

    while ((end = str.find(" ", start)) != string::npos) {
        tokens.push_back(str.substr(start, end - start));

        start = end + 1;
    }

    tokens.push_back(str.substr(start));

    return tokens;
}
