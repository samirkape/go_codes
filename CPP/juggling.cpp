
#include<iostream>
#include <cmath>
#include<string.h>


using namespace std;
int GCD(int A, int B) 
{ 
    if (B == 0) 
        return A; 
    else
        return GCD(B, A % B); 
} 

void rotate(int * arr, int size, int d){

    int sz = GCD(size, d);
    int j = 0;

    if (sz==1){
        for(int i=0; i<d; i++){
            int tmp = arr[0];
            for(j=0; j+sz < size; j+=sz ){
                arr[j] = arr[j+sz];
            }
            arr[j] = tmp;
        }
    }
    else{
        for(int i=0; i<sz; i++){
            int tmp = arr[i];
            for(j=i; j+d < size; j+=d ){
                arr[j] = arr[j+d];
            }
            arr[j] = tmp;
        }
    }
    for(int i=0; i<size; i++)
        cout<<arr[i]<<" ";
}

int main()
{
   //char arr[] = {'a', 'b','c'};
   //int arr[] = {1, 2, 3, 4}, K = 3;
   //int arr[] = {5, 10, 12, 13, 15, 18}, K = 30;
   //int arr[] = {7,3,2,5,8}, K = 14;
   int arr[] = {1,2,3,4,5}, K = 3;
   int size = sizeof(arr)/sizeof(int);
   rotate(arr,size,K);
}