/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Input: list1 = [1,2,4], list2 = [1,3,5]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return list1
	}
	if list1 == nil && list2 != nil{
		return list2
	}
	if list2 == nil && list1 != nil {
		return list1
	}

	if list1.Val <= list2.Val {
        list1.Next = mergeTwoLists(list1.Next, list2)
        return list1
    }

	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
