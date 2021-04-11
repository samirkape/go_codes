#include <bits/stdc++.h>

using namespace std;

#define NO_OF_CHARS 10

vector<string> split_string(string);

void badChar( vector<string> str, int size,  
                        int * badchar)  
{  
    int i;  
    for (i = 0; i < 10; i++)  
        badchar[i] = -1;  
    for (i = 0; i < size; i++)  
        badchar[stoi(str[i])] = i;  
}  

typedef struct{
    int row;
    int col;
    int offset;
    int not_found;
}s_first_ocr_meta;

typedef struct{
    int row;
    int result;
}s_rest_ocr_meta;

typedef struct{
    vector<string> txt;
    vector<string> patt;
    int p_idx;
}s_pre_proc;


s_rest_ocr_meta rest_pattern( s_first_ocr_meta cordinates,vector<string> G, vector<string> P ){
    int max_txt_rows = G.size();
    int max_patt_rows = P.size(); 
    int txt_row_start = cordinates.row + 1; // +1 cause we want to skip matched row
    int txt_col_start = cordinates.col;
    int patlength = P[0].size();
    int txtlen = G[0].size();
    int patt_start = 1;
    int counter = patt_start;
    s_rest_ocr_meta meta;

    for( int i = txt_row_start, j = patt_start; i < max_txt_rows && j < max_patt_rows; i++, j++ ){
        string patt = P[j]; 
        auto txt = G[i].begin() + txt_col_start;
        int row_number = i % txtlen;
        string new_txt(txt , txt + patlength);

        if ( new_txt == patt ){
            counter++;
        }
        else{
            meta.result = 0;
            meta.row = row_number;
            return meta;            
        }
        if( counter == max_patt_rows ){
            meta.result = 1;
            return meta;
        }
    }
    return meta;
}


s_first_ocr_meta first_ocr( vector<string> patt, vector<string> txt, int p_idx, int * badchar, int txt_size ){
    int shift = 0;
    int new_shift = 0;
    s_first_ocr_meta meta;
    while(1){   
        while( (txt[p_idx + shift] == patt[p_idx]) && p_idx != -1 ) // check match
            p_idx--;
        if( p_idx > -1 ){      // no match
            //new_shift = search( txt[p_idx + shift] , patt );
            new_shift = max(1, p_idx - badchar[stoi(txt[shift + p_idx])]);
            shift += new_shift;
            p_idx = patt.size() - 1; 
        }
        else
        {
            int column = ( shift ) % txt_size;
            int row = ( shift ) / txt_size ;
            meta.row = row;
            meta.col = column;
            meta.not_found = 0;
            return meta;
        }
        
        if( p_idx + shift >= txt.size() ){
            meta.not_found = 1;
            return meta;
        }
    }
    return meta;
}

s_pre_proc preproc( vector<string> G, vector<string> P ){
    string tx;
    string pt;
    int pi = 0;
    int j = 0;
    s_pre_proc store;

    for( auto i : G ) tx.append(i);
    vector<string> txt(tx.size()); 
    for ( auto i : tx ) txt[j++] = i;

    j = 0;
    for( auto i : P ) pt.append(i);
    vector<string> patt(P[0].size()); 
    for ( auto i : pt ){
        if( j == P[0].size() ) break;
        patt[j++] = i;
    }

    store.patt = patt;
    store.txt = txt;
    store.p_idx = patt.size() - 1;

    return store; 
}
// Complete the gridSearch function below.
string gridSearch(vector<string> G, vector<string> P) {
    int psize = ( P.size() * P[0].size());
    int gsize = ( G.size() * G[0].size());
    int txt_size = G[0].size();
    int pat_size = P[0].size();
    int txt_rows = G.size(); 
    string res ;

    int shift = 0;
    int ofset = 0;
    int flg = 0;
    int p_idx = 0;
    s_first_ocr_meta first_occur_meta;
    s_rest_ocr_meta rest_occur_meta;
    first_occur_meta.offset = 0;
    //pre_process
    s_pre_proc store = preproc( G, P );
    vector<string> txt = store.txt;
    vector<string> patt = store.patt; 
    p_idx = store.p_idx;
    int bad_char[10];

    // call bad char
    badChar( patt, p_idx + 1, bad_char );

    for ( int i = 0; i + pat_size <  gsize; i++){
        vector <string> new_txt( txt.begin() + i, txt.end() );
        first_occur_meta = first_ocr( patt, new_txt, p_idx, bad_char, txt_size);
        if(i)first_occur_meta.row = ( i / txt_rows ) + first_occur_meta.row;
        if( !first_occur_meta.not_found ){
            rest_occur_meta = rest_pattern( first_occur_meta, G, P );
            if( !rest_occur_meta.result ){
                i = (rest_occur_meta.row * txt_size) - 1;
            }
            else{
                return "YES";
            }
        }
        else{
            return "NO";
            i += txt_size ;
        }
    }
    
    //first_ocr();

    // design loop to iterate through number of rows in G

    // for( int i = 0; (i + psize) < gsize; i++ ){   
    //     while( (txt[j + shift] == patt[j])) // check match
    //         j--;
    //     if( j > -1 ){      // no match
    //         s = search( txt[j + shift] , patt );
    //         shift += s;
    //         j = pi;
    //     }
    //     else{      // match
    //         int column = shift % G[0].size();
    //         int row = ( shift / G[0].size() ) + 1; // +1 cause we want to skip the 1st matched row
    //         for( int iter_p = 1; iter_p < P.size(); iter_p++, row++ ){
    //             string spatt = P[iter_p]; 
    //             string new_txt(G[row].begin() + column  , G[row].begin() + column + pi + 1); 
    //             int res = search_rest( new_txt, spatt );
    //             if( !res ){
    //                 flg = 0;
    //                 shift = G[0].size();
    //                 j = P[0].size() - 1;
    //                 break;
    //             }
    //             else{
    //                 flg++;
    //                 if( flg == P.size() - 1 ) return "YES";
    //             }
    //         }
    //     }
    //     i = shift;
    // }
    // return "NO";
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int t;
    cin >> t;

    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    for (int t_itr = 0; t_itr < t; t_itr++) {
        string RC_temp;
        getline(cin, RC_temp);

        vector<string> RC = split_string(RC_temp);

        int R = stoi(RC[0]);

        int C = stoi(RC[1]);

        vector<string> G(R);

        for (int i = 0; i < R; i++) {
            string G_item;
            getline(cin, G_item);

            G[i] = G_item;
        }

        string rc_temp;
        getline(cin, rc_temp);

        vector<string> rc = split_string(rc_temp);

        int r = stoi(rc[0]);

        int c = stoi(rc[1]);

        vector<string> P(r);

        for (int i = 0; i < r; i++) {
            string P_item;
            getline(cin, P_item);

            P[i] = P_item;
        }

        string result = gridSearch(G, P);

        fout << result << "\n";
    }

    fout.close();

    return 0;
}

vector<string> split_string(string input_string) {
    string::iterator new_end = unique(input_string.begin(), input_string.end(), [] (const char &x, const char &y) {
        return x == y and x == ' ';
    });

    input_string.erase(new_end, input_string.end());

    while (input_string[input_string.length() - 1] == ' ') {
        input_string.pop_back();
    }

    vector<string> splits;
    char delimiter = ' ';

    size_t i = 0;
    size_t pos = input_string.find(delimiter);

    while (pos != string::npos) {
        splits.push_back(input_string.substr(i, pos - i));

        i = pos + 1;
        pos = input_string.find(delimiter, i);
    }

    splits.push_back(input_string.substr(i, min(pos, input_string.length()) - i + 1));

    return splits;
}
