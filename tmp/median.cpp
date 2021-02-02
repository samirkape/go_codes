#include <map>
#include <list>
#include <cmath>
#include <ctime>
#include <deque>
#include <queue>
#include <stack>
#include <bitset>
#include <cstdio>
#include <limits>
#include <vector>
#include <cstdlib>
#include <numeric>
#include <sstream>
#include <iostream>
#include <algorithm>
#define NIL -0.5
using namespace std;
/* Head ends here */
typedef struct{
    long double fout;
    long int iout;
}mout; 

mout median(list<long int> &v)
{
    mout out;
    out.fout = NIL;
    long double tmp = 0.0;
    size_t n = v.size() / 2;
    size_t nv =  v.size();
    if( n!=0 && (( nv % 2 == 0 ) || nv == 2 )){
        if( !( v[nv-1] < v[n] ) && nv != 2 )
            sort(v.begin(),v.end());
        if(nv == 2 && (v[0] == v[1])){ out.iout = v[0]; return out; }
        tmp = (v[n-1] + v[n]) / 2.0;
        out.fout = tmp;
        out.iout = tmp;
        return out;
    }
    else 
    {   
        nth_element(v.begin(), v.begin()+n, v.end());
        out.fout = v[n];
        out.iout = v[n];
    }
    return out;
}

void format_out( mout out ){
    int isIntF = floor(out.fout);
    int isIntC = ceil(out.fout);
    if( out.fout != NIL && out.fout != (long double)0.0 && isIntC != isIntF ) printf("%.1Lf\n",out.fout);
    else cout<<out.iout<<endl;
    fflush(stdout);
}

void driver(vector<char> s,vector<int> X) {
    int size = X.size();
    vector<long int> xc;
    list<long int> xl;
    int j = 0;
    vector<long int> :: iterator it;
    mout out; 
    int nfl = 0;

    for(auto i = 0; i < size; i++ ){
        if(s[i] == 'a'){
            xl.push_back(X[i]);
            out = median(xl);
            format_out(out);
            fflush(stdout);
        }
        else if(s[i] == 'r'){
            if(xc.empty() || xc.size() == 1 ){ cout<<"Wrong!"<<endl; continue; }
            else{
                for(it = xc.begin(); it != xc.end(); it++ ){
                    if(xc[j] == X[i]){
                        xc.erase(it);
                        out = median(xc);
                        format_out(out);
                        nfl = 1;
                        break;
                    }
                    j++;
                } 
                //xc.erase(j);
            }
            if(!nfl) cout<<"Wrong!"<<endl;
            fflush(stdout);
            j=0;
            nfl = 0;
        }
    } 
    
}
int main(void){

//Helpers for input and output

    int N;
    cin >> N;
    
    vector<char> s;
    vector<int> X;
    char temp;
    int tempint;
    for(int i = 0; i < N; i++){
        cin >> temp >> tempint;
        s.push_back(temp);
        X.push_back(tempint);
    }
    
    driver(s,X);
    return 0;
}