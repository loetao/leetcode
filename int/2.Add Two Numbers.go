You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	current := &ListNode{}
    res:=current
    addTwoNumber(l1,l2,res)
    
	return res
}

func addTwoNumber(l1 *ListNode, l2 *ListNode,l3 *ListNode)  {
    if l1 == nil && l2==nil{
		return 
	}
	
     temp:=l3
    var g,t int
    if l1==nil{
        l1=&ListNode{Val: 0}
    }
    if l2==nil{
       l2=&ListNode{Val: 0}
    }
   
    
    g=(l1.Val+l2.Val+l3.Val)%10
    t=(l1.Val+l2.Val+l3.Val)/10
    temp.Val = g
    
    l1=l1.Next
    l2=l2.Next
    if l1==nil && l2==nil && t==0{  
        
    }else{
              
        
        temp.Next=&ListNode{Val: t}
    }
    l3=temp.Next
    addTwoNumber(l1,l2,l3)
}
