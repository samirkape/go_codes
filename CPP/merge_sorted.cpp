#include<vector>
#include<algorithm> 
#include<limits.h>
#include<unordered_map>

using namespace std;

class Solution {
public:
    int get_pos_idx( vector<int> & arr, int arr_size, int key ){
        int l, r, mid; 
        l = r = mid = 0;
        
        r = arr_size - 1;
        l = 0;

        if( key < arr[l] ) return l;
        
        for( int i = 0; i < arr_size; i++ ){
            mid = ( r - l ) / 2 + l;
            int mid_e = arr[mid];
            
            if( key == mid_e || ( key <= arr[mid + 1] && key > mid_e ) )
                return mid + 1;
            
            if( key > mid_e ){
                l == mid ? l -= 1: l = mid + 1; 
            }
            else if( key < mid_e ){
                r == mid ? r -= 1: r = mid - 1; 
            }
        }
        
        return 0;
    }
    void merge(vector<int>& arr1, int m, vector<int>& arr2, int n) {
        int exit_flag = 0;
        int res1 = m;
        int res2 = n;
        
        for( int i = 0; i < n; i++ ){
            auto it = arr1.begin();
            if( m != 0 && arr2[i] < arr1[m-1] ){
                int pos = get_pos_idx( arr1, m, arr2[i] );
                arr1.insert( it + pos, arr2[i] );
                m += 1;
            }
            
            else{
                for( int j = m; i < n; j++ ){
                    arr1[j] = arr2[i++];
                }
                break;
            }
        }
        arr1.resize( res1+res2 ) ;
    }
};

int main(){
    Solution sol;
    vector<int> x = {0,0,3,0,0,0,0,0,0} ;
    vector<int> y = { -1,1,1,1,2,3 } ;
    sol.merge(x, 3, y, 6);
    int xc;
}