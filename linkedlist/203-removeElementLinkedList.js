var removeElements = function (head, val) {
  while (head !== null && head.val === val) {
    head = head.next;
  }

  let curr = head;
  while (curr !== null && curr.next !== null) {
    if (curr.next.val === val) {
      curr.next = curr.next.next;
    } else {
      curr = curr.next;
    }
  }

  return head;
};
