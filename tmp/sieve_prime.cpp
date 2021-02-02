#include<iostream>
#include <cmath>
#include<string.h>

using namespace std;

void sieve(int N) 
{ 
    bool isPrime[N + 1]; 
    for (int i = 0; i<N; ++i) { 
        isPrime[i] = true; 
    } 
  
    isPrime[0] = false; 
    isPrime[1] = false; 
  
    for (int i = 2; i * i <= N; ++i) { 
  
        // Mark all the multiples of i as composite numbers 
        if (isPrime[i] == true) { 
            for (int j = i * i; j <= N; j += i) 
                isPrime[j] = false; 
        } 
    } 
} 


int main()
{
   //char arr[] = {'a', 'b','c'};
   //int arr[] = {1, 2, 3, 4}, K = 3;
   //int arr[] = {5, 10, 12, 13, 15, 18}, K = 30;
   //int arr[] = {7,3,2,5,8}, K = 14;
   int arr[] = {7, 3, 8, 10, 14, 18 , 13, 12}, K = 5;
   int size = sizeof(arr)/sizeof(int);
   int x = sieve(10);
   cout<<"GDC : "<<x<<endl;
}