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

int N = 5;
// Check if the bit at index pos is on in mask.
bool on(int mask, int pos) { return (mask & (1 << pos)) > 0; }
// Set the bit at index pos in mask.
int set(int mask, int pos) { return mask | (1 << pos); }

void subsets(int * L) {
    long int x = INTMAX_MAX;
    x = (char)x | 1<<63;
// mask will iterate through all 2^N subsets.
for(int mask = 0; mask < (1 << N); mask++) {
// Do something problem-specific with the subset. Here I’ll just
// print it.
for(int k = 0; k < N; k++) {
if(on(mask, k)) {
    cout<<L[k]<<" ";
}
}
cout<<"\n";
}
}

int main()
{
   //char arr[] = {'a', 'b','c'};
   //int arr[] = {1, 2, 3, 4}, K = 3;
   //int arr[] = {5, 10, 12, 13, 15, 18}, K = 30;
   //int arr[] = {7,3,2,5,8}, K = 14;
   int arr[] = {1,2,3,4,5}, K = 5;
   int size = sizeof(arr)/sizeof(int);
   //int x = GCD(8,3);
   //cout<<"GDC : "<<x<<endl;
   subsets(arr);
}