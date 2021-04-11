#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
#include<string>
using namespace std;

void form_stack( vector<long int> input ){
    vector<long int> out;
    long int is = input.size();
    long int max = -INT64_MAX;
    for(int i=0; i<is;){
        if(input[i] == 1){
               if(input[i+1] > max) max = input[i+1]; 
               out.push_back(input[i+1]); 
               i+=2;
        }   
        else if (input[i] == 2){
            int pop = out[out.size() - 1];
            out.pop_back();
            if(pop >= max) {
                if(out.empty())
                    max = -INT64_MAX;
                else max = *max_element(out.begin(), out.end());
            }
            i++;
        }
        else if (input[i] == 3){
            cout<<max<<endl;
            i++;
        }
    }
}

int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */   
    int N;
    cin>>N;
    long int x = 0, y = 0;
    vector<long int> input;

    for(int i = 0; i < N; i++){
        cin>>x;
        input.push_back(x);
        if(x == 1){ cin>>y; input.push_back(y); }
    }
    form_stack(input, N);
    return 0;
}
