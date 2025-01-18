/**
 * You are given the head of a linked list, which contains a series of integers separated by 0's. The beginning and end of the linked list will have Node.val == 0.
 * For every two consecutive 0's, merge all the nodes lying in between them into a single node whose value is the sum of all the merged nodes. The modified list should not contain any 0's.
 * Return the head of the modified linked list.
 */

/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val === undefined ? 0 : val)
        this.next = (next === undefined ? null : next)
    }
}

// [0,3,1,0,4,5,2,0]
function mergeNodes(head: ListNode | null): ListNode | null {
    if(head === null) return null;
    let newList: ListNode | null = null;
    let newListHead: ListNode | null = newList;
    while(head?.next !== null) {
        if(head.val === 0) {
            if(newList === null) {
                newList = new ListNode();
                newListHead = newList;
            }
            else {
                newList.next = new ListNode();
                newList = newList.next;
            }
        } else {
            newList!.val += head.val;
        }
        head = head?.next;
    }
    return newListHead;
};