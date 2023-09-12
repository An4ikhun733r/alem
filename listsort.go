package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	middle := getMiddle(l)
	nextToMiddle := middle.Next
	middle.Next = nil

	left := ListSort(l)
	right := ListSort(nextToMiddle)

	return merge(left, right)
}

func getMiddle(head *NodeI) *NodeI {
	if head == nil {
		return head
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(left, right *NodeI) *NodeI {
	result := &NodeI{}
	current := result

	for left != nil && right != nil {
		if left.Data <= right.Data {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		current = current.Next
	}

	if left != nil {
		current.Next = left
	} else if right != nil {
		current.Next = right
	}

	return result.Next
}
