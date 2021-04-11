#include<iostream>
#include <cmath>
#include<string.h>

using namespace std;
int * grp = nullptr;

int find_df(int * arr){
    int sum = 0;
    sum = arr[0] - arr[1];
    return abs(sum);
}

int * select_group(int * arr, int diff, int size){
    static int min = diff;
    static int * group = (int*)malloc(sizeof(int)*size);
    if( diff < min ){
        memcpy(group,arr,size*__SIZEOF_INT__);
    }
    return group;
}

void print(int * arr, int size){
    int diff = 0;
    for( int j = 0; j < size-1; j++ ){
        diff += find_df(arr+j);
        fflush(stdout);
    }
    grp = select_group(arr, diff, size);
    //cout<<"sum : "<<diff;
    cout<<"\n";
}

void power_set(int * arr, int size, int K){
    int k = 0, sum = 0, idx = 0;
    int pwr = pow(size,2);
    int out[size];
    int fg = 0;
    int ctr = 0;

    for(int i = 0; i <= pwr; i++ )
    {
        k = i;
        for( int j = 0; j < size; j++ )
        {
            if( (char)k & 1 )
            {
                sum = abs(sum) - arr[j];
                out[idx] = arr[j];
                idx++;
                fg = 1;
            }
            fflush(stdout);
            k = k>>1;
        }
        if(fg == 1) ctr++;
        if(idx == K) print(out, idx);
        sum = 0;
        idx = 0;
        fg = 0;
    }
    for(int i = 0 ;i < K; i++)
        cout<<grp[i]<<" ";
}

int main()
{
   //char arr[] = {'a', 'b','c'};
   //int arr[] = {1, 2, 3, 4}, K = 3;
   //int arr[] = {5, 10, 12, 13, 15, 18}, K = 30;
   //int arr[] = {7,3,2,5,8}, K = 14;
   int arr[] = {7, 3, 8, 10, 14, 18 , 13, 12}, K = 5;
   int size = sizeof(arr)/sizeof(int);
   power_set(arr,size,K);
}
