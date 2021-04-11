#include<iostream>

using namespace std;

void rearrange(int * arr, int size){
    int out[size];
    for(int i=0; i < size; i++){
        out[i] = -1;
    }
    for(int i=0; i < size; i++){
        if(arr[i]>=0){
            out[arr[i]] = arr[i];
        }
    }
    for(int i=0; i < size; i++){
        cout<<out[i]<<" ";
    }
}

int main()
{
   int arr[] = {-1, -1, 6, 1, 9, 3, 2, -1, 4, -1};
   int size = sizeof(arr)/__SIZEOF_INT__;
   rearrange(arr,size);
}
