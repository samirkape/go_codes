#include <bits/stdc++.h>
#include<cmath>

using namespace std;

// Complete the encryption function below.
typedef struct{
    int s_floor;
    int s_ceil;
}s_dim;

s_dim checkSize(int fl, int cl, int size){
    int s_fl = 0;
    s_dim new_dim;
    if(size > ( fl * cl )){
        s_fl = ceil( size / (float)cl );
        new_dim.s_floor = s_fl;
        new_dim.s_ceil = cl;
    }
    else{
        new_dim.s_floor = fl;
        new_dim.s_ceil = cl;       
    }
    return new_dim; 
}

void trimStr( string &s ){
    s.erase(remove(s.begin(), s.end(), ' '),s.end());
}

string encryption(string s) {
    s_dim new_dim;
    trimStr( s );
    int fl = floor(sqrt(s.size()));
    int cl = ceil(sqrt(s.size()));
    new_dim = checkSize(fl, cl, s.size());
    int row = 0;
    int newfl = new_dim.s_floor;
    int newcl = new_dim.s_ceil;
    string tmp;
    vector<char> vtmp;
    int idx = 0;
    
    for(int col = 0; col < newcl; col++){
        idx = col;
        for(row = 0; row < newfl, idx < s.size(); row++ ){
            vtmp.push_back(s[idx]);
            idx += newcl; 
        }
        vtmp.push_back(' ');
    }
    string new_s(vtmp.begin(), vtmp.end()-1);
    return new_s;
}   

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string s;
    getline(cin, s);

    string result = encryption(s);

    fout << result << "\n";

    fout.close();

    return 0;
}
