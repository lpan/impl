class BinTree:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


"""
          0
        /   \
       1     2
      / \   / \
     3   4  5  6
"""
def default_tree():
    return BinTree(0,
                   BinTree(1,
                           BinTree(3),
                           BinTree(4)),
                   BinTree(2,
                           BinTree(5),
                           BinTree(6)))

def traverser(traverse_fn):
    visited = []
    visit_fn = lambda node : visited.append(node.val)

    def wrapper(tree):
        traverse_fn(tree, visit_fn)
        return visited

    return wrapper

@traverser
def preorder(tree, visit):
    """
    visit node before push children on to the stack
    """

    def traverse(node):
        if node:
            visit(node)
            traverse(node.left)
            traverse(node.right)

    traverse(tree)

@traverser
def preorder_iter(tree, visit):
    """
    1. visit node
    2. push children on the stack (left on top of right tho)
    3. pop stack and go back to #1 until stack is empty
    """

    stack = [tree]
    while stack:
        node = stack.pop()
        if node:
            visit(node)
            stack.append(node.right)
            stack.append(node.left)

@traverser
def inorder(tree, visit):
    """
    inorder visits the left-most leaf node first, then root, then right sub-tree

    visit the node after its left child's stack frame is popped
    """

    def traverse(node):
        if node:
            traverse(node.left)
            visit(node)
            traverse(node.right)

    traverse(tree)

@traverser
def inorder_iter(tree, visit):
    """
    1. reach the left-most leaf of the root, push nodes onto the stack along the
    way
    2. pop the stack, visit it
    3. go to step #1, treat the right child of the node as the root
    4. when stack is empty, return
    """

    stack = []
    root = tree

    while True:
        while root:
            stack.append(root)
            root = root.left

        if not stack:
            return

        node = stack.pop()
        visit(node)

        root = node.right

@traverser
def postorder(tree, visit):
    """
    postorder means visiting the root last.
    visit left, right, root

    visit node when the stack frame of the node is about to be popped
    """

    def traverse(node):
        if node:
            traverse(node.left)
            traverse(node.right)
            visit(node)

    traverse(tree)

@traverser
def postorder_iter(tree, visit):
    """
    https://www.geeksforgeeks.org/iterative-postorder-traversal/

    basically it is a reverse of preorder...
    """

    stack = [tree]
    result = []

    while stack:
        node = stack.pop()
        if node:
            stack.append(node.left)
            stack.append(node.right)
            result.append(node)

    while result:
        visit(result.pop())

if __name__ == "__main__":
    tree = default_tree()
    print("preorder:", preorder(tree))
    print("preorder:", preorder_iter(tree))
    print("inorder:", inorder(tree))
    print("inorder:", inorder_iter(tree))
    print("postorder:", postorder(tree))
    print("postorder:", postorder_iter(tree))
