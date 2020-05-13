class Node {

  constructor(val) {
    this.val = val;
    this.edges = [];
  }
}

class Edge {

  constructor(node1, node2, weight) {
    this.node1 = node1;
    this.node2 = node2;
    this.weight = weight;
  }
}

class NodeGraph {

  constructor() {
    this.nodes = [];
  }

  addNode = function (node) {
    this.nodes.push(node);
  }

}