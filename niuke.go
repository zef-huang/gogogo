package main

import (
	"fmt"
)


 type ListNode struct{
   Val int
   Next *ListNode
 }
 

/**
 * 
 * @param pHead1 ListNode类 
 * @param pHead2 ListNode类 
 * @return ListNode类
*/
func FindFirstCommonNode( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
    // write code here
    p1 := pHead1
    p2 := pHead2
    
    len1 := 0
    len2 := 0
    
    for{
        if p1 == nil {
            break
        }
        len1 = len1 + 1
        p1 = p1.Next
    }
    
    for{
        if p2 == nil {
            break
        }
        len2 = len2 + 1
        p2 = p2.Next
    }
	
    for {
        if len1 == len2 {
			fmt.Println("pHead1.Val", pHead1.Val)
			fmt.Println("pHead2.Val", pHead2.Val)
            if pHead1.Val != pHead2.Val {
                pHead1 = pHead1.Next
				pHead2 = pHead2.Next
				len1 = len1 - 1
				len2 = len2 - 1			
			} 
			print("pHead1.Val", pHead1.Val)
            if pHead1.Val == pHead2.Val {
                break
            }
            
        }
        if len1 > len2 {
			len1 = len1 - 1
            pHead1 = pHead1.Next
        }
        if len1 < len2 {
			len2 = len2 - 1
            pHead2 = pHead2.Next
        }
    }
    
    return pHead1
}


func main() {
	var p1 ListNode
	p1.Val = 12

	var p2 ListNode
	p2.Val = 12
	FindFirstCommonNode(&p1, &p2)
}