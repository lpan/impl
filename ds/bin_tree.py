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

"""
visit node before push children on to the stack
"""
@traverser
def preorder(tree, visit):
    def traverse(node):
        if node:
            visit(node)
            traverse(node.left)
            traverse(node.right)

    traverse(tree)

"""
1. visit node
2. push children on the stack (left on top of right tho)
3. pop stack and go back to #1 until stack is empty
"""
@traverser
def preorder_iter(tree, visit):
    stack = [tree]
    while stack:
        node = stack.pop()
        if node:
            visit(node)
            stack.append(node.right)
            stack.append(node.left)

"""
visit node right before stack starts to unwind
1. reach to the deepest left node
2. visit the node as stack frames being popped
3. push right nodes onto the stack
"""
@traverser
def inorder(tree, visit):
    def traverse(node):
        if node:
            traverse(node.left)
            visit(node)
            traverse(node.right)

    traverse(tree)

"""
1. reach the left-most leaf of the root, push nodes onto the stack along the
way
2. pop the stack, visit it
3. go to step #1, treat the right child of the node as the root
4. when stack is empty, return
"""
@traverser
def inorder_iter(tree, visit):
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

if __name__ == "__main__":
    tree = default_tree()
    print("preorder:", preorder(tree))
    print("preorder:", preorder_iter(tree))
    print("inorder:", inorder(tree))
    print("inorder:", inorder_iter(tree))
