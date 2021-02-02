#include<iostream>
#include <cmath>
#include<string.h>

using namespace std;
int swap(int *xp, int *yp) 
{ 
    int temp = *xp; 
    *xp = *yp; 
    *yp = temp; 
} 
  
int sort_(int arr[], int n) 
{ 
   int i, j; 
   for (i = 0; i < n-1; i++)       
       for (j = 0; j < n-i-1; j++)  
           if (arr[j] > arr[j+1]) 
              swap(&arr[j], &arr[j+1]); 
} 
int check_triplet(int * arr){
    if(pow(arr[0],2) + pow(arr[1],2) == pow(arr[2],2)){
        //cout<<arr[0]<<" "<<arr[1]<<" "<<arr[2]<<endl;
        return 1;
    }
    else 0;
}

void calc_permutations(int * arr, int size){
    long int k   = 0; 
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
                if(idx>3) break;
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
   //int arr[] = {68,35,1,70,25,79,59,63,65,6,46,82,28,62,92,96,43,28,37,92,5,3,54,93,83,22,17,19,96,48,27,72,39,70,13,68,100,36,95,4,12,23,34,74};
   //int arr[] = {3,6,23,12,54,13,22,264,265};
   string k = "k";
   char s = 's';
   string sam = s + k;

   int arr[] = {42,12,54,69,48,45,63,58,38,60,24,42,30,79,17,36,91,43,89,7,41,43,65,49,47,6,91,30,71,51,7,2,94,49,30,24,85,55,57,41,67,77,32,9,45,40,27,24,38,39,19,83,30,42,34,16,40,59,5,31,78};
   int size = sizeof(arr)/sizeof(int);
   //sort_(arr, size);
   //calc_permutations(arr, size);
}
