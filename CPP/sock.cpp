#include<iostream>
#include <cmath>

using namespace std;



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

int is_pair(int * arr){
    if(arr[0]==arr[1]) return 1;
    else return 0;
}

void find_sock_pair(int * arr, int size){
    int counter = 0;
    int on = 0;
    for(int i = 0; i+1 < size; i++ ){
        on = is_pair(arr+i);
        if(on){
            i++;
            counter++;
        }
    }
    cout<<counter;
}

int main()
{
   //char arr[] = {'a', 'b','c'};
   //int arr[] = {1, 2, 3, 4}, K = 3;
   //int arr[] = {5, 10, 12, 13, 15, 18}, K = 30;
   int arr[] = {50,49,38,49,78,36,25,96,10,67,78,58,98,8,53,1,4,7,29,6,59,93,74,3,67,47,12,85,84,40,81,85,89,70,33,66,6,9,13,67,75,42,24,73,49,28,25,5,86,53,10,44,45,35,47,11,81,10,47,16,49,79,52,89,100,36,6,57,96,18,23,71,11,99,95,12,78,19,16,64,23,77,7,19,11,5,81,43,14,27,11,63,57,62,3,56,50,9,13,45}, K = 14;
   int size = sizeof(arr)/sizeof(int);
   sort_(arr,size);
   find_sock_pair(arr,size);
}
