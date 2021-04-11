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

void find_sum_pair(int * arr, int size, int sum ){
   for(int i=0; i< size; i++)
      for(int j=i; j< size; j++){
         if ( arr[i] + arr[j] == sum && arr[j]!=arr[i]){
            cout<<"found pair : "<<arr[i]<<" "<<arr[j];
            cout<<"\n";
         }
      }
}

int main()
{
   int arr[] = {1, 5, 7, -1, 5};
   int sum = 6;
   find_sum_pair(arr,sizeof(arr)/__SIZEOF_INT__,sum);
}
