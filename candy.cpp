#include <bits/stdc++.h>

using namespace std;

// Complete the candies function below.
long candies(int n, vector<int> grades) {
    long long int sum = 0;
    vector<int> candies(n, 1);
    for(int i = 0; i + 1 < n; i++ ){
        int curr = grades[ i + 1 ];
        int prev = grades[ i ];
        int curr_idx = i + 1;
        int prev_idx = i;
        if( curr > prev ){
            candies[ curr_idx ] = candies[ prev_idx ] + 1; 
        }else if( curr < prev ){
            continue;
        }
    }
    for( auto i = n - 1; i - 1 >= 0; i-- ){
        int curr = grades[ i - 1 ];
        int prev = grades[ i ];
        int curr_idx = i - 1;
        int prev_idx = i;
        if(  curr > prev ){
            if ( candies[ curr_idx ] <= candies[ prev_idx ] ){
                candies[ curr_idx ] = candies[ prev_idx ] + 1;
            }
        }
    }
    for(int i = 0; i < n; i++ ) sum += candies[i];
    return sum;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    vector<int> arr(n);

    for (int i = 0; i < n; i++) {
        int arr_item;
        cin >> arr_item;
        cin.ignore(numeric_limits<streamsize>::max(), '\n');

        arr[i] = arr_item;
    }

    long result = candies(n, arr);

    fout << result << "\n";

    fout.close();

    return 0;
}
