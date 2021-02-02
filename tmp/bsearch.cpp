#include<iostream>
#include<vector>
#include<string>
#include<cmath>

using namespace std;

int binarySearch(string w, int key){
    int left = 0;
    int right = w.size();
    int mid = 0;
    while( left <= right && (left != w.size()) ){
        mid = ((left + right-1)/2);
        if(key > w[mid]) 
            right = mid - 1;
        else if( key < w[mid] ) 
            left = mid + 1;
        else return mid;
    }
    return mid;
}
int main(void) 
{ 
    string w = "biehzcmjckznhwrfgglverzyxuqpj"; 
    int size = w.size();
    char x = w[size-2]; 
    
    int result = binarySearch(w.substr(size-6,size), 'r'); 
    (result == -1) ? cout << "Element is not present in array"
                   : cout << "Element is present at index " << result; 
    return 0; 
} 