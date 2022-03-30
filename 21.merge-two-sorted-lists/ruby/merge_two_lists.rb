# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val = 0, _next = nil)
#         @val = val
#         @next = _next
#     end
# end
# @param {ListNode} list1
# @param {ListNode} list2
# @return {ListNode}
def merge_two_lists(list1, list2)
  return list1 || list2 unless list1 && list2

  head, list1, list2 = pick_min_and_advance(list1, list2)
  current = head
  while list1 || list2
    current.next, list1, list2 = pick_min_and_advance(list1, list2)
    current = current.next
  end
  head
end

def pick_min_and_advance(list1, list2)
  return [list2, list1, list2.next] if !list1
  return [list1, list1.next, list2] if !list2 || list1.val <= list2.val

  [list2, list1, list2.next]
end
