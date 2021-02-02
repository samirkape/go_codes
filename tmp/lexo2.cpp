// Find the next lexicographically 
// greater permutation of a word 

#include <bits/stdc++.h>
#include <iostream> 

using namespace std; 

void swap(char* a, char* b) 
{ 
	if (*a == *b) 
		return; 
	*a ^= *b; 
	*b ^= *a; 
	*a ^= *b; 
} 
void rev(string& s, int l, int r) 
{ 
	while (l < r) 
		swap(&s[l++], &s[r--]); 
} 

int bsearch(string& s, int l, int r, int key) 
{ 
	int index = -1; 
	while (l <= r) { 
		int mid = l + (r - l) / 2; 
		if (s[mid] <= key) 
			r = mid - 1; 
		else { 
			l = mid + 1; 
			if (index == -1 || s[index] >= s[mid]) 
				index = mid; 
		} 
	} 
	return index; 
} 

bool nextpermutation(string& s) 
{ 
	int len = s.length(), i = len - 2; 
	while (i >= 0 && s[i] >= s[i + 1]) 
		--i; 
	if (i < 0) 
		return false; 
	else { 
		int index = bsearch(s, i + 1, len - 1, s[i]); 
		swap(&s[i], &s[index]); 
		rev(s, i + 1, len - 1); 
		return true; 
	} 
} 

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int T;
    cin >> T;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    for (int T_itr = 0; T_itr < T; T_itr++) {
        string w;
        getline(cin, w);

        bool result = nextpermutation(w);

        fout << result << "\n";
    }

    fout.close();

    return 0;
}