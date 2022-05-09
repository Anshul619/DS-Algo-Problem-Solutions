package leetCodeProblems.LinkedList;

/**
 * LeetCode - https://leetcode.com/problems/merge-two-sorted-lists/
 * InterviewBit - https://www.interviewbit.com/problems/merge-two-sorted-lists/
 * 
 * @author anshul.agrawal
 *
 */
public class MergeTwoSortedLists21 {
	
	static class ListNode {
		
		int val;
		ListNode next;
		
		ListNode(int val1) {
			val = val1;
		}
	}
	
	ListNode mergedList = null;
    ListNode lastPointer = null;
    
    public void appendNodeToMergedList(int value) {
        ListNode newNode = new ListNode(value);
                
        if (mergedList == null) {
            mergedList = newNode;
            lastPointer = mergedList;
        }
        else {
            ListNode tempPointer = lastPointer;
            tempPointer.next = newNode;
            
            lastPointer = newNode;
        }
        
        return;
    }
    
    public ListNode incrementNextPointer(ListNode node) {
        if (node.next != null) {
            node = node.next;
        }
        else {
            node = null;
        }
        
        return node;
    }
    
    public void print() {
        
        ListNode mergedListPointer = mergedList;
        
        while(mergedListPointer != null) {
            System.out.println(mergedListPointer.val);
            mergedListPointer = incrementNextPointer(mergedListPointer);
         }
    }
    
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        
        ListNode list1Pointer = list1;
        ListNode list2Pointer = list2;
        
        while(list1Pointer != null && list2Pointer != null) {
            
            if (list1Pointer.val == list2Pointer.val) { 
                
                //System.out.println("Case1 - list2Pointer is equal to list1Pointer");
                
                appendNodeToMergedList(list1Pointer.val);
                appendNodeToMergedList(list2Pointer.val);
                
                list1Pointer = incrementNextPointer(list1Pointer);
                list2Pointer = incrementNextPointer(list2Pointer);
            
            }
            else if (list1Pointer.val < list2Pointer.val) {
                
                //System.out.println("Case2 - list2Pointer is greater than list1Pointer");
                
                appendNodeToMergedList(list1Pointer.val);
                list1Pointer = incrementNextPointer(list1Pointer);

            }
            else {
                
                //System.out.println("Case3 - list1Pointer is greater than list2Pointer");
                
                appendNodeToMergedList(list2Pointer.val);
                list2Pointer = incrementNextPointer(list2Pointer);
            }
            
        }
        
        while(list1Pointer != null) {
             appendNodeToMergedList(list1Pointer.val);
             list1Pointer = incrementNextPointer(list1Pointer);
        }
        
        while(list2Pointer != null) {
             appendNodeToMergedList(list2Pointer.val);
             list2Pointer = incrementNextPointer(list2Pointer);
        }
        
        return mergedList;
    }
    
    public static void main(String[] args) {
 
        ListNode list1 = new ListNode(5);
        list1.next = new ListNode(15);
        list1.next.next = new ListNode(25);
        
        ListNode list2 = new ListNode(5);
        list2.next = new ListNode(10);
        list2.next.next = new ListNode(20);
        
        MergeTwoSortedLists21 obj = new MergeTwoSortedLists21();
        obj.mergeTwoLists(list1, list2);
        
        obj.print();
        
    }

}
