// Note that by default C++ creates a max-heap 
// for priority queue 
#include <iostream> 
#include <queue> 
#include<set>
using namespace std; 
  
void showpq(priority_queue <int> gq) 
{ 
    priority_queue <int> g = gq;
    multiset <int, greater <int> > gquiz1; 
    gquiz1.insert(40); 
    gquiz1.insert(30); 
    gquiz1.insert(60); 
    gquiz1.insert(20); 
    gquiz1.insert(50); 
    gquiz1.insert(50); // 50 will be added again to the multiset unlike set 
    gquiz1.insert(10);
    vector <int> ss;
    
    multiset <int, greater <int> > :: iterator itr; 
    cout << "\nThe multiset gquiz1 is : "; 
    for (itr = gquiz1.begin(); itr != gquiz1.end(); ++itr) 
    { 
        cout << '\t' << *itr; 
    } 
    cout << '\n'; 
} 
  
int main () 
{ 
    priority_queue <int> gquiz; 
    gquiz.push(10); 
    gquiz.push(30); 
    gquiz.push(20); 
    gquiz.push(5); 
    gquiz.push(1); 
  
    cout << "The priority queue gquiz is : "; 
    showpq(gquiz); 
  
    cout << "\ngquiz.size() : " << gquiz.size(); 
    cout << "\ngquiz.top() : " << gquiz.top(); 
  
  
    cout << "\ngquiz.pop() : "; 
    gquiz.pop(); 
    showpq(gquiz); 
  
    return 0; 
}