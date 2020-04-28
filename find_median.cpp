#include<iostream>

using namespace std;

int swap(int *xp, int *yp) 
{ 
    int temp = *xp; 
    *xp = *yp; 
    *yp = temp; 
} 
  
int * sort_(int arr[], int n) 
{ 
   int i, j; 
   int * sorted = (int*)malloc(n*sizeof(int));
   for(int i = 0; i < n; i++ ){
       sorted[i] = arr[i];
   }
   for (i = 0; i < n-1; i++)       
       for (j = 0; j < n-i-1; j++)  
           if (sorted[j] > sorted[j+1]) 
              swap(&sorted[j], &sorted[j+1]); 
    return sorted;
} 

int find_med(int * arr, int slide, int size){
    int * sorted;
    if(slide%2==0){
        for(int i=0; i+slide <= size; i++){
            sorted = sort_(i+arr,slide);
            cout<<(float)(sorted[1]+sorted[2])/2<<" ";
            fflush(stdout);
            free(sorted);
        }
    }else
    {
        for(int i=0; i+slide <= size; i++){
            sorted = sort_(i+arr,slide);
            cout<<sorted[1]<<" ";
            fflush(stdout);
            free(sorted);
        }
    }   
}

int main()
{
   int arr[] = {-1, 5, 13, 8, 2, 3, 3, 1};
   int k = 3;
   int size = sizeof(arr)/__SIZEOF_INT__;
   find_med(arr,k,size);
   
}
