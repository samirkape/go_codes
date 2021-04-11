#include <bits/stdc++.h>
#include<unordered_map>
using namespace std;

int find_pattern(unordered_map <int,int> set){
    int max_ = 0;
    int index_ = 0;
    unordered_map<int, int>:: iterator p;
    for (p = set.begin(); p != set.end(); p++) 
         if (max_ < (p->first * p->second)){
            max_ = p->first * p->second;
            index_ = p->first;
         }
    return index_;
}

int compute(unordered_map<int, int> umap, int key){
    unordered_map<int, int>:: iterator p;
    int flag = 0;
    int count = 0;
    int bal = 1;
    for (p = umap.begin(); p != umap.end(); p++){
        if(p->first!=key){
            if((p->first * p->second)-bal == key || (p->first * p->second)-bal == 0)
                flag = 1;
            bal--;
        }
    }
    if(bal == 0 && flag == 1) return 1;
    else return 0;
}
int oddPlaces( unordered_map< int , int > umap, int key ){
    int count = 0;
    unordered_map<int, int>:: iterator p;
    for (p = umap.begin(); p != umap.end(); p++){
        if(p->first!=key)
            count++;
    }
    return count;
}
string isValid(string s) {
    unordered_map<int, int> umap;
    vector<int> :: iterator ip;

    int arr[26] = {0};
    for ( int i=0; i<s.size(); i++){
        if(s[i]!=0){
            int alpha = (int)s[i]-97;
            arr[alpha] += 1;
        }
    }
        
    for(int j = 0; j < 26; j++) if(arr[j]) umap[arr[j]] += 1;
    if ( umap.size() == 1 || s.size() == 1) return "YES"; 
    int patt = find_pattern(umap);
    if(compute(umap, patt)) return "YES";
    return "NO";
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string s;
    getline(cin, s);

    string result = isValid(s);

    fout << result << "\n";

    fout.close();

    return 0;
}
