# Data Structures Demo
# Demonstrates various data structures implemented in El

(let
    # Stack implementation
    _ (print "=== Stack Implementation ===")
    make-stack (lambda [] [])
    stack-push (lambda stack item [item *stack])
    stack-pop (lambda stack (match (len stack)
        0 [nil []]
        [*slice stack [0] (slice stack (range 1 (len stack)))]
    ))
    stack-peek (lambda stack (match (len stack)
        0 nil
        *slice stack [0]
    ))
    stack-empty (lambda stack (eq (len stack) 0))
    
    # Stack operations
    stack (make-stack)
    stack (stack-push stack 1)
    stack (stack-push stack 2)
    stack (stack-push stack 3)
    _ (print "stack after pushes:" stack)
    
    top (stack-peek stack)
    _ (print "stack top:" top)
    
    result (stack-pop stack)
    popped (slice result [0])
    stack (slice result [1])
    _ (print "popped:" popped)
    _ (print "stack after pop:" stack)
    
    # Queue implementation
    _ (print "\n=== Queue Implementation ===")
    make-queue (lambda [] [])
    queue-enqueue (lambda queue item [*queue item])
    queue-dequeue (lambda queue (match (len queue)
        0 [nil []]
        [*slice queue [0] (slice queue (range 1 (len queue)))]
    ))
    queue-front (lambda queue (match (len queue)
        0 nil
        *slice queue [0]
    ))
    queue-empty (lambda queue (eq (len queue) 0))
    
    # Queue operations
    queue (make-queue)
    queue (queue-enqueue queue "first")
    queue (queue-enqueue queue "second")
    queue (queue-enqueue queue "third")
    _ (print "queue after enqueues:" queue)
    
    front (queue-front queue)
    _ (print "queue front:" front)
    
    result (queue-dequeue queue)
    dequeued (slice result [0])
    queue (slice result [1])
    _ (print "dequeued:" dequeued)
    _ (print "queue after dequeue:" queue)
    
    # Tree representation
    _ (print "\n=== Tree Representation ===")
    # Tree as [value left-child right-child]
    make-tree (lambda value left right [value left right])
    tree-value (lambda tree *slice tree [0])
    tree-left (lambda tree (slice tree [1]))
    tree-right (lambda tree (slice tree [2])
    tree-empty (lambda tree (eq (len tree) 0))
    
    # Create a simple tree
    leaf1 (make-tree 1 [] [])
    leaf2 (make-tree 3 [] [])
    leaf3 (make-tree 5 [] [])
    node1 (make-tree 2 leaf1 leaf2)
    root (make-tree 4 node1 leaf3)
    _ (print "tree structure:" root)
    
    # Tree traversal (preorder)
    preorder (lambda tree (match (tree-empty tree)
        true []
        [*slice tree [0] *preorder (tree-left tree) *preorder (tree-right tree)]
    ))
    _ (print "preorder traversal:" (preorder root))
    
    # Binary search tree operations
    _ (print "\n=== Binary Search Tree Operations ===")
    bst-insert (lambda tree value (match (tree-empty tree)
        true (make-tree value [] [])
        (let
            current-value (tree-value tree)
            (match (lt value current-value)
                true (make-tree current-value (bst-insert (tree-left tree) value) (tree-right tree))
                (make-tree current-value (tree-left tree) (bst-insert (tree-right tree) value))
            )
        )
    ))
    
    bst-search (lambda tree value (match (tree-empty tree)
        true false
        (let
            current-value (tree-value tree)
            (match (eq value current-value)
                true true
                (match (lt value current-value)
                    true (bst-search (tree-left tree) value)
                    (bst-search (tree-right tree) value)
                )
            )
        )
    ))
    
    # Build a BST
    bst []
    bst (bst-insert bst 5)
    bst (bst-insert bst 3)
    bst (bst-insert bst 7)
    bst (bst-insert bst 1)
    bst (bst-insert bst 9)
    _ (print "BST structure:" bst)
    _ (print "BST search for 3:" (bst-search bst 3))
    _ (print "BST search for 6:" (bst-search bst 6))
    
    # Hash table simulation (using lists)
    _ (print "\n=== Hash Table Simulation ===")
    make-hash-table (lambda [] [])
    hash-put (lambda table key value (let
        # Simple hash function simulation
        hash-key (mod key 10)
        # Would need more complex implementation for real hash table
        [key value *table]
    ))
    hash-get (lambda table key (let
        # Simple linear search for demonstration
        search-helper (lambda table key (match (len table)
            0 nil
            (match (eq *slice table [0] key)
                true *slice table [1]
                (search-helper (slice table (range 2 (len table))) key)
            )
        ))
        (search-helper table key)
    ))
    
    # Hash table operations
    hash-table (make-hash-table)
    hash-table (hash-put hash-table 1 "one")
    hash-table (hash-put hash-table 2 "two")
    hash-table (hash-put hash-table 3 "three")
    _ (print "hash table:" hash-table)
    _ (print "hash-get(2):" (hash-get hash-table 2))
    _ (print "hash-get(4):" (hash-get hash-table 4))
    
    # Set implementation (using lists)
    _ (print "\n=== Set Implementation ===")
    make-set (lambda [] [])
    set-add (lambda set item (let
        contains (lambda set item (match (len set)
            0 false
            (match (eq *slice set [0] item)
                true true
                (contains (slice set (range 1 (len set))) item)
            )
        ))
        (match (contains set item)
            true set
            [item *set]
        )
    ))
    set-contains (lambda set item (match (len set)
        0 false
        (match (eq *slice set [0] item)
            true true
            (set-contains (slice set (range 1 (len set))) item)
        )
    ))
    set-union (lambda set1 set2 (let
        union-helper (lambda set1 set2 (match (len set2)
            0 set1
            (let
                item *slice set2 [0]
                rest (slice set2 (range 1 (len set2)))
                new-set1 (set-add set1 item)
                (union-helper new-set1 rest)
            )
        ))
        (union-helper set1 set2)
    ))
    
    # Set operations
    set1 (make-set)
    set1 (set-add set1 1)
    set1 (set-add set1 2)
    set1 (set-add set1 3)
    _ (print "set1:" set1)
    
    set2 (make-set)
    set2 (set-add set2 3)
    set2 (set-add set2 4)
    set2 (set-add set2 5)
    _ (print "set2:" set2)
    
    union-set (set-union set1 set2)
    _ (print "union of set1 and set2:" union-set)
    
    _ (print "set1 contains 2:" (set-contains set1 2))
    _ (print "set1 contains 6:" (set-contains set1 6))
    
    # Linked list representation
    _ (print "\n=== Linked List Representation ===")
    # Linked list as [value next]
    make-node (lambda value next [value next])
    node-value (lambda node *slice node [0])
    node-next (lambda node (slice node [1]))
    list-empty (lambda node (eq (len node) 0))
    
    # Create a linked list
    list3 (make-node 3 [])
    list2 (make-node 2 list3)
    list1 (make-node 1 list2)
    _ (print "linked list:" list1)
    
    # Linked list traversal
    list-length (lambda list (match (list-empty list)
        true 0
        {1 + list-length (node-next list)}
    ))
    _ (print "linked list length:" (list-length list1))
    
    # List to array conversion
    list-to-array (lambda list (match (list-empty list)
        true []
        [*slice list [0] *list-to-array (node-next list)]
    ))
    _ (print "list to array:" (list-to-array list1))
    
    nil
)
