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


int main()
{
   //char arr[] = {'a', 'b','c'};
   //int arr[] = {1, 2, 3, 4}, K = 3;
   //int arr[] = {5, 10, 12, 13, 15, 18}, K = 30;
   //int arr[] = {7,3,2,5,8}, K = 14;
   int arr[] = {7, 3, 8, 10, 14, 18 , 13, 12}, K = 5;
   int size = sizeof(arr)/sizeof(int);
   int x = GCD(8,3);
   cout<<"GDC : "<<x<<endl;
}