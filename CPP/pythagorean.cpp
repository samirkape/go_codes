#include<iostream>
#include <cmath>
#include<string.h>

using namespace std;

int check_triplet(int * arr){
    if(pow(arr[0],2) + pow(arr[1],2) == pow(arr[2],2)){
        cout<<arr[0]<<" "<<arr[1]<<" "<<arr[2]<<endl;
        return 1;
    }
    else 0;
}

void calc_permutations(int * arr, int size){
    int k   = 0; 
    int sum = 0;
    int idx = 0;
    int fg  = 0;
    int res = 0;
    int ctr = 0;
    int pwr = pow(size,3);
    int out[size];

    for(int i = 0; i <= pwr; i++ )
    {
        k = i;
        for( int j = 0; j < size; j++ )
        {
            if( (char)k & 1 )
            {
                out[idx] = arr[j];
                idx++;
                fg = 1;
            }
            fflush(stdout);
            k = k>>1;
        }
        if(fg == 1) ctr++;
        if(idx == 3) res = check_triplet(out);
        if(res) break;
        sum = 0;
        idx = 0;
        fg = 0;
    }
    if(res)
        cout<<"YES"<<endl;
    else
        cout<<"NO"<<endl;   
}

int main()
{
   //char arr[] = {'a', 'b','c'};
   //int arr[] = {1, 2, 3, 4}, K = 3;
   //int arr[] = {5, 10, 12, 13, 15, 18}, K = 30;
   //int arr[] = {7,3,2,5,8}, K = 14;
   int arr[] = {3,2,4,6,5};
   //int arr[] = {3,6,23,12,54,13,22,264,265};
   int size = sizeof(arr)/sizeof(int);
   calc_permutations(arr, size);
}
