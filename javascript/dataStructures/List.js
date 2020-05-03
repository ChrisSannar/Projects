class Node {

  constructor(val) {
    this.val = val;
    this.next = null;
  }
}

class List {

  constructor() {
    this.head = null;
    this.tail = null;
    this.length = 0;
  }

  add = function (val) {
    if (this.tail) {
      this.tail.next = new Node(val);
      this.tail = this.tail.next;
    }
    else if (this.head) {
      this.tail = new Node(val);
      this.head.next = this.tail;
    }
    else {
      this.head = this.tail = new Node(val);
    }
    this.length++;
  }

  find = function (func) {
    let current = this.head;
    while (current != null) {
      if (func(current.val)) {
        return current.val;
      }
      current = current.next;
    }
    return null;
  }

  clear = function () {
    this.head = null;
    this.tail = null;
    this.length = 0;
  }

  toString = function () {
    let result = '[ '
    let current = this.head;
    if (current) {
      while (current.next != null) {
        result += `${current.val} -> `;
        current = current.next;
      }
      result += `${current.val} ]`;
      return result;
    } else {
      return result + ']';
    }
  }
}

module.exports = { Node, List }