#include<vector>
#include<iostream>

using namespace std;


int b_search(vector<int> brr, int value, vector<int>::iterator lower, vector<int>::iterator upper)
{
    while(lower <= upper)
    {
        vector<int>::iterator mid = lower + distance(lower, upper)/2;
        if(*mid == value)
        {
            *mid = -1;
            return 1;
        }
        else if(*mid < value)
        {
            lower = mid + 1;
        }
        else
        {
            upper = mid - 1;
        }
    }
    return 0;
}

int main(void) 
{ 
    vector<int> w = {1,2,3,4,5,6,7}; 
    int size = w.size();
    auto st =  w.begin();
    auto en = w.end();
    int result = b_search(w, 6,st,en ); 
    (result == -1) ? cout << "Element is not present in array"
                   : cout << "Element is present at index " << result; 
    return 0; 
} 