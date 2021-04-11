#include <bits/stdc++.h>
#include<map>

using namespace std;

vector<char> border;
//vector<char> cvt;
map<int,char> cvt;
int idxf = 0;

void get_edges( string grid, int size ){
    for(int i=0; i<size; i++){
        border.push_back(grid[i]);
    }
}
void get_border( string grid, int size , int idx){
    for(int i = 0; i < size; i++ )
        if(i != 0 && i != size-1 ){
            idxf = size * idx + i;
            cvt[idxf] = grid[i];
        }
}
vector<string> fr_cavity(vector<string> &grid, int size ){
    char up = 0;
    char down = 0;
    char prev = 0;
    char next = 0;
    map <int, int> cvt_;
    vector<string> cgrid = grid;
    for(auto i=cvt.begin(); i!=cvt.end(); i++){
        char val = i -> second;
        up = grid[(i->first / size)-1][(i->first - size)%size];
        down = grid[(i->first / size)+1][(i->first)%size];
        prev = grid[(i->first / size)][((i->first)%size) - 1];
        next = grid[(i->first / size)][((i->first)%size) + 1];
        if( val > up && val > down && val > prev && val > next ){
                cgrid[(i->first)/size][i->first%size] = 'X';
        }
    }
    return cgrid;
}
// Complete the cavityMap function below.
vector<string> cavityMap(vector<string> grid) {
    int size = grid.size();
    vector<string> out;
    for(int i = 0; i < size; i++ ){
        if(i == 0 || i == size-1 ) get_edges(grid[i], size);
        else get_border(grid[i], size, i);
    }
    out = fr_cavity(grid, size);
    return out;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    vector<string> grid(n);

    for (int i = 0; i < n; i++) {
        string grid_item;
        getline(cin, grid_item);

        grid[i] = grid_item;
    }

    vector<string> result = cavityMap(grid);

    for (int i = 0; i < result.size(); i++) {
        fout << result[i];

        if (i != result.size() - 1) {
            fout << "\n";
        }
    }

    fout << "\n";

    fout.close();

    return 0;
}
