
class Node {
  val: number
  children: Node[]
  constructor(val?: number) {
    this.val = (val === undefined ? 0 : val)
    this.children = []
  }
}


function preorder(root: Node | null): number[] {
  if (root === null) return []

  let result: number[] = [];

  let node_queue = [root];

  while (node_queue.length > 0) {
    let node = node_queue.shift() as Node;

    result.push(node.val);
    node_queue = node.children.concat(node_queue);
  }

  return result
};

export {}
