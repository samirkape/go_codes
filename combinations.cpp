#include<iostream>
#include <cmath>

using namespace std;

void print(int * arr, int size){
    for( int j = 0; j < size; j++ ){
        cout<<arr[j]<<" ";
        fflush(stdout);
    }
    cout<<"\n";
}

int swap(int *xp, int *yp) 
{ 
    int temp = *xp; 
    *xp = *yp; 
    *yp = temp; 
} 
  
int  sort_(int arr[], int n) 
{ 
   int i, j; 
   for (i = 0; i < n-1; i++)       
       for (j = 0; j < n-i-1; j++)  
           if (arr[j] > arr[j+1]) 
              swap(&arr[j], &arr[j+1]); 
} 

void power_set(int * arr, int size, int K){
    int k = 0, sum = 0, idx = 0;
    int pwr = pow(size,2);
    int out[size];
    int fg = 0;
    int ctr = 0;

    for(int i = 0; i < pwr; i++ ){
        k = i;
        for( int j = 0; j < size; j++ ){
            if( (char)k & 1 )
            {
                sum += arr[j];
                out[idx] = arr[j];
                idx++;
                fg = 1;
            }
            fflush(stdout);
            k = k>>1;
        }
        if(fg == 1) ctr++;
        if(sum == K) print(out, idx);
        sum = 0;
        idx = 0;
        fg = 0;
    }
    cout<<"combs"<<ctr;
}

int main()
{
   //char arr[] = {'a', 'b','c'};
   //int arr[] = {1, 2, 3, 4}, K = 3;
   int arr[] = {5, 10, 12, 13, 15, 18}, K = 15;
   int size = sizeof(arr)/sizeof(int);
   sort_(arr,size);
   power_set(arr,size,K);
}
