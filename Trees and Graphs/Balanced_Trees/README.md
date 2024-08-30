# **Balanced Trees**

Balanced trees are data structures that automatically maintain a balanced height to ensure efficient operations. Two popular types of balanced trees are AVL Trees and Red-Black Trees. 

## **AVL Tree**

### **Overview:**
An AVL Tree is a self-balancing binary search tree where the heights of two child subtrees of any node differ by no more than one. This ensures that operations such as insertion, deletion, and lookup are efficient.

### **Operations:**
- **Insertion:** Ensures balance after insertion using rotations.
- **Deletion:** Ensures balance after deletion using rotations.

### **Python Implementation:**

Here's a detailed implementation of an AVL Tree, including insertion and rotations:

```python
class AVLNode:
    def __init__(self, key):
        self.key = key
        self.left = None
        self.right = None
        self.height = 1

class AVLTree:
    def insert(self, root, key):
        # Step 1: Perform normal BST insert
        if not root:
            return AVLNode(key)
        elif key < root.key:
            root.left = self.insert(root.left, key)
        else:
            root.right = self.insert(root.right, key)
        
        # Step 2: Update height of this ancestor node
        root.height = 1 + max(self.get_height(root.left), self.get_height(root.right))
        
        # Step 3: Get the balance factor
        balance = self.get_balance(root)
        
        # Step 4: Balance the tree if needed
        # Left Left Case
        if balance > 1 and key < root.left.key:
            return self.right_rotate(root)
        # Right Right Case
        if balance < -1 and key > root.right.key:
            return self.left_rotate(root)
        # Left Right Case
        if balance > 1 and key > root.left.key:
            root.left = self.left_rotate(root.left)
            return self.right_rotate(root)
        # Right Left Case
        if balance < -1 and key < root.right.key:
            root.right = self.right_rotate(root.right)
            return self.left_rotate(root)
        
        return root

    def left_rotate(self, z):
        y = z.right
        T2 = y.left
        y.left = z
        z.right = T2
        z.height = 1 + max(self.get_height(z.left), self.get_height(z.right))
        y.height = 1 + max(self.get_height(y.left), self.get_height(y.right))
        return y

    def right_rotate(self, z):
        y = z.left
        T3 = y.right
        y.right = z
        z.left = T3
        z.height = 1 + max(self.get_height(z.left), self.get_height(z.right))
        y.height = 1 + max(self.get_height(y.left), self.get_height(y.right))
        return y

    def get_height(self, root):
        if not root:
            return 0
        return root.height

    def get_balance(self, root):
        if not root:
            return 0
        return self.get_height(root.left) - self.get_height(root.right)

    def pre_order(self, root):
        result = []
        if root:
            result.append(root.key)
            result.extend(self.pre_order(root.left))
            result.extend(self.pre_order(root.right))
        return result

# Example usage
avl = AVLTree()
root = None
keys = [10, 20, 30, 15, 25]
for key in keys:
    root = avl.insert(root, key)
print("Pre-order traversal of the AVL Tree:", avl.pre_order(root))
```

#### **Explanation:**
- **Insertion**: Inserts a new key while ensuring the tree remains balanced by performing rotations as necessary.
- **Rotations**: `left_rotate` and `right_rotate` help in balancing the tree when it becomes unbalanced.

---

## **Red-Black Tree**

**Overview:**
A Red-Black Tree is a binary search tree with an additional property of balancing based on colors of the nodes. It ensures that the tree remains approximately balanced, providing good performance for insertion, deletion, and search operations.

### **Operations:**
- **Insertion:** Ensures balance after insertion using recoloring and rotations.
- **Deletion:** Ensures balance after deletion using recoloring and rotations.

### **Python Implementation:**

Here’s a detailed implementation of a Red-Black Tree, including insertion and rotations:

```python
class RedBlackNode:
    def __init__(self, key, color="RED"):
        self.key = key
        self.left = None
        self.right = None
        self.color = color
        self.parent = None

class RedBlackTree:
    def __init__(self):
        self.NIL = RedBlackNode(key=None, color="BLACK")
        self.root = self.NIL

    def insert(self, key):
        node = RedBlackNode(key)
        node.left = self.NIL
        node.right = self.NIL
        parent = None
        current = self.root

        while current != self.NIL:
            parent = current
            if node.key < current.key:
                current = current.left
            else:
                current = current.right

        node.parent = parent
        if parent is None:
            self.root = node
        elif node.key < parent.key:
            parent.left = node
        else:
            parent.right = node
        
        node.color = "RED"
        self.insert_fixup(node)

    def insert_fixup(self, node):
        while node != self.root and node.parent.color == "RED":
            if node.parent == node.parent.parent.left:
                uncle = node.parent.parent.right
                if uncle.color == "RED":
                    node.parent.color = "BLACK"
                    uncle.color = "BLACK"
                    node.parent.parent.color = "RED"
                    node = node.parent.parent
                else:
                    if node == node.parent.right:
                        node = node.parent
                        self.left_rotate(node)
                    node.parent.color = "BLACK"
                    node.parent.parent.color = "RED"
                    self.right_rotate(node.parent.parent)
            else:
                uncle = node.parent.parent.left
                if uncle.color == "RED":
                    node.parent.color = "BLACK"
                    uncle.color = "BLACK"
                    node.parent.parent.color = "RED"
                    node = node.parent.parent
                else:
                    if node == node.parent.left:
                        node = node.parent
                        self.right_rotate(node)
                    node.parent.color = "BLACK"
                    node.parent.parent.color = "RED"
                    self.left_rotate(node.parent.parent)
        
        self.root.color = "BLACK"

    def left_rotate(self, x):
        y = x.right
        x.right = y.left
        if y.left != self.NIL:
            y.left.parent = x
        y.parent = x.parent
        if x.parent is None:
            self.root = y
        elif x == x.parent.left:
            x.parent.left = y
        else:
            x.parent.right = y
        y.left = x
        x.parent = y

    def right_rotate(self, y):
        x = y.left
        y.left = x.right
        if x.right != self.NIL:
            x.right.parent = y
        x.parent = y.parent
        if y.parent is None:
            self.root = x
        elif y == y.parent.right:
            y.parent.right = x
        else:
            y.parent.left = x
        x.right = y
        y.parent = x

    def pre_order(self, node):
        result = []
        if node != self.NIL:
            result.append((node.key, node.color))
            result.extend(self.pre_order(node.left))
            result.extend(self.pre_order(node.right))
        return result

# Example usage
rbt = RedBlackTree()
keys = [10, 20, 30, 15, 25]
for key in keys:
    rbt.insert(key)
print("Pre-order traversal of the Red-Black Tree:", rbt.pre_order(rbt.root))
```

#### **Explanation:**
- **Insertion**: Adds a new key while ensuring the tree maintains Red-Black properties.
- **Fixup**: `insert_fixup` corrects violations of Red-Black Tree properties using rotations and recoloring.
- **Rotations**: `left_rotate` and `right_rotate` are used to maintain the balance of the tree.
