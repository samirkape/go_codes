#include<iostream>

using namespace std;

int swap(int *xp, int *yp) 
{ 
    int temp = *xp; 
    *xp = *yp; 
    *yp = temp; 
} 
  
int  sort_features_array(int arr[], int n) 
{ 
   int i, j; 
   for (i = 0; i < n-1; i++)       
       for (j = 0; j < n-i-1; j++)  
           if (arr[j] > arr[j+1]) 
              swap(&arr[j], &arr[j+1]); 
} 

int main()
{
   int arr[] = {2, 5, 17, -1};
   int sum = 7;
   int size = sizeof(arr)/__SIZEOF_INT__;
   sort_features_array(arr,size);
   
}
