#include <iostream>
using namespace std;
#include<vector>
#include<stdlib.h>
#include<string.h>

void steps(int * arr, int N){
    int ct = 0;
	double y = __DBL_MAX__;
	int x= __INT32_MAX__;
	auto p = "sting";
    for(int i=0; i+1 < N;){
        
        i += arr[i];
        ct++;
    }
    if(ct) cout<<ct;
    else cout<<-1;
}

void getNum(int * arr, long long int N){
	while(1){
		static long long int num = N;
		if(num){
		char digit = num % 10;
		num  /= 10;
		arr[digit] += 1;
		}
		else break;
	}
}

// void getNumStr( int * arr, string st ){
// 	for (int i=0; i< st.size(); i++){
// 		arr[atoi((char*)st[i])] += 1;
// 	}
// }

void isRound(int * arr, long long int N){
	int flag = 0;
	for(int check = 1; check<5; check++){
		long long int num2 = N*check;
		int tmp[10] = {0};
		memcpy(tmp,arr, 10 * __SIZEOF_INT__);
		while(1){
			if(num2){
				char digit = num2 % 10;
				num2  /= 10;
				tmp[digit] -= 1;
			}
			else{
				for(int i = 1; i < 10; i++){
					if(tmp[i]!=0){flag = 1; break;}
				}
				break;
			};
		}
		if(flag) break;
	}
	
	if(flag) cout<<"Not a round number"<<endl;
	else cout<<"Its a round number"<<endl;
}
int main() {
	int T;
	string st = "588235294117649";
	long long int N = 588235294117648;
	int arr[10] = {0};
	int i = 0;
	getNum(arr, N);
	isRound(arr,N);
	return 0;
}