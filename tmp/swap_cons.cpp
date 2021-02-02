#include <iostream>
#include "dll..cpp"
#include<stdio.h>
#include <map>
#include <vector>

using namespace std;
struct Node{
int data;
struct Node * next;
};

struct Node * create_node()
{
    struct Node * node = nullptr;
    return node = new Node;
}

/*---------------------------------------------------------------*/
int * insert_at_next(struct Node ** node, int data, int pos ){
    struct Node * new_node = new Node;
    struct Node * tmp = *node;
    new_node->data = data;
    
    if ( *node == NULL){
        new_node->next = nullptr;
        *node = new_node;
    }
    else
    {
        while( tmp->next != NULL )
        {
            tmp = tmp->next;
        }
        new_node->next = tmp->next;
        tmp->next = new_node;
    }
    return (int*)new_node;
}

void insert_at_front(struct Node ** head, int data)
{
    struct Node * tmp = *head;    
    struct Node * new_node  = create_node();
    new_node->data = data;
    if(tmp==NULL){
        new_node->next = nullptr;
        *head = new_node;
    }
    else{
        new_node->next = *head;
        *head = new_node;
    }
}
/*--------------------------------------------------------------*/

void trace_list_fwd(struct Node * head){
    system("clear");
    struct Node * temp = head;
    while( temp != NULL){
        cout<<temp->data;
        cout<<"->";
        temp = temp->next;
    }
}

/*
void trace_list_bck(struct Node * head){
    system("clear");
    struct Node * temp = head;
    while( temp->next != NULL)
    {
        temp = temp->next;
    }
    do
    {
        cout<<temp->data;
        cout<<"->";
        temp = temp->prev;
        fflush(stdout);
    }while (temp!=NULL);
}
*/
typedef struct Node Node;

void swap_blocks(Node ** prev, Node ** curr){
    Node * tmp  = NULL;
    tmp = *curr;
    (*prev)->next->next = *prev;
    (*prev)->next = tmp->next;
    //curr->next = prev;
    //prev->next = tmp->next;
}

void alter_nodes(Node ** head){
    Node * prev = NULL;
    Node * curr = NULL;
    Node * tmp = NULL;
    Node * tmp_head = NULL;
    prev = *head;
    curr = (*head)->next;

    while(curr!=NULL){
        swap_blocks(&prev,&curr);
        tmp_head = prev;
        /*tmp = prev->next;
        curr = tmp->next;
        prev = tmp;*/
    }
    *head = tmp_head;
}
/*--------------------------------------------------------*/

// void delete_node_at_pos(struct Node ** head, int pos){
//     struct Node * tmp  = *head;
//     int keep_count = 1;
//     if(*head == NULL){
//         perror("Empty list");
//         return;
//     }
//     if(pos == 1){
//         if((*head)->next != NULL){
//         *head = (*head)->next;
//         (*head)->prev = NULL;
//         }
//         delete(tmp);
//         return;
//     }
//     while(keep_count != pos){
//         tmp = tmp->next;
//         keep_count++;
//     }
//     tmp->prev->next = tmp->next;
//     delete(tmp);
// }



void delete_node_at_pos( Node ** prev, Node ** curr){
    Node * tmp = *curr;
    (*prev)->next = (*curr)->next;
    delete(tmp);
}

// void delete_node_at_pos( Node ** curr ){
//     Node * tmp = *curr;
//     (*curr)->prev->next = (*curr)->next;
//     delete(tmp);
// }


void remove_sorted_dups(Node ** head){    
    Node * curr = NULL;
    Node * prev = NULL;
    curr = (*head)->next;
    prev = *head;
    

    while (curr != NULL)
    {

       if(prev->data == curr->data){
           delete_node_at_pos(&prev,&curr);
           curr = prev->next;
       }else
       {
           prev = curr;
           curr = curr->next;
       }
    }
}

void remove_dups(Node ** head){    
    Node * curr = NULL;
    Node * tmp = NULL;
    curr = (*head);
    std::map<int,int> store;
    int counter = 1;
  
    // check if key is present    
    while (curr != NULL)
    {
       if(store.find(curr->data) != store.end())
       {
           //delete_node_at_pos(&prev,&curr);
           tmp = curr;
           cout << "map contains key %d!\n",curr->data;
           //delete_node_at_pos(&curr);
           curr = tmp->next;
       }else
       {
           store[curr->data];
           curr = curr->next;
       }
    }
}

void del_by_node(Node * node)
{
   Node * tmp = node->next;
   node->data = node->next->data;
   node->next = node->next->next;
   delete(tmp);
}

int main()
{
    struct Node * head = nullptr;
    int get_choice = 1;
    int data = 500;
    int pos = 5;
    int i = 0;
    int x = 0;
    vector<int *> nd_ar;
    int ct = 0;
    int nd = 0;
    while(1)
    {

        cout<<"\n";
        cout<<"at front = 1"<<endl;
        cout<<"at last  = 2"<<endl;
        cout<<"delete at pos = 3"<<endl;
        cout<<"remove duplicates = 4"<<endl;
        cout<<"display fwd  = 5"<<endl;
        cout<<"display bck  = 6"<<endl;
        cin>>get_choice;
        if( get_choice == 1 ){
            cin>>data;
            insert_at_front(&head,data);
            trace_list_fwd(head);
            fflush(stdout);
            i++;
        }
        else if ( get_choice == 2 ){
            cin>>data;
            nd_ar.push_back(insert_at_next(&head, data, pos));
            trace_list_fwd(head);
        }
        else if ( get_choice == 3 ){
            cout<<" enter node number "<<endl;
            cin>>nd;
            Node * tmp = (Node *)nd_ar[nd-1];
            del_by_node(tmp);
        }
        else if ( get_choice == 4 )
        {
            trace_list_fwd(head);
        }

        else break;
    }

    return 0;
}