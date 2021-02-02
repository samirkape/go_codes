#include <iostream>
#include "dll..cpp"
#include<stdio.h>

using namespace std;
struct Node{
int data;
struct Node * next;
struct Node * prev;
};

struct Node * create_node()
{
    struct Node * node = nullptr;
    return node = new Node;
}

/*---------------------------------------------------------------*/
void insert_at_next(struct Node ** node, int data, int pos ){
    struct Node * new_node = new Node;
    struct Node * tmp = *node;
    new_node->data = data;
    
    if ( *node == NULL){
        new_node->next = nullptr;
        new_node->prev = nullptr;
        *node = new_node;
    }
    else
    {
        while( tmp->next != NULL )
        {
            tmp = tmp->next;
        }
        new_node->next = tmp->next;
        new_node->prev = tmp;
        tmp->next = new_node;
    }
}

void insert_at_front(struct Node ** head, int data)
{
    struct Node * tmp = *head;    
    struct Node * new_node  = create_node();
    new_node->data = data;
    if(tmp==NULL){
        new_node->next = nullptr;
        new_node->prev = nullptr;
        *head = new_node;
    }
    else{
        new_node->next = *head;
        (*head)->prev = new_node; 
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
/*--------------------------------------------------------*/

void delete_node_at_pos(struct Node ** head, int pos){
    struct Node * tmp  = *head;
    int keep_count = 1;
    if(*head == NULL){
        perror("Empty list");
        return;
    }
    if(pos == 1){
        if((*head)->next != NULL){
        *head = (*head)->next;
        (*head)->prev = NULL;
        }
        delete(tmp);
        return;
    }
    while(keep_count != pos){
        tmp = tmp->next;
        keep_count++;
    }
    tmp->prev->next = tmp->next;
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


    while(1)
    {
        cout<<"\n";
        cout<<"at front = 1"<<endl;
        cout<<"at last  = 2"<<endl;
        cout<<"delete at pos = 3"<<endl;
        cout<<"display fwd  = 4"<<endl;
        cout<<"display bck  = 5"<<endl;
        cin>>get_choice;
        if( get_choice == 1 ){
            cin>>data;
            insert_at_front(&head,data);
            fflush(stdout);
            i++;
        }
        else if ( get_choice == 2 ){
            cin>>data;
            insert_at_next(&head, data, pos);
        }
        else if ( get_choice == 3 )
        {
            cout<<"enter pos :";
            cin>>pos;
            delete_node_at_pos(&head, pos);
        }
        else if ( get_choice == 4 )
        {
            trace_list_fwd(head);
        }
        else if ( get_choice == 5 )
        {
            trace_list_bck(head);
        }
        else break;
    }

    return 0;
}