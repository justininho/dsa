type 'a bst = 
| Leaf
| Node of 'a bst * 'a  * 'a bst

let empty () = Leaf

let is_empty = function
| Leaf -> true
| Node _ -> false

let root = function
| Leaf -> None
| Node (_, k, _) -> Some k


let rec insert bst key =
  match bst with
    | Leaf -> Node(Leaf, key, Leaf)
    | Node(l, k, r) -> 
      if key > k then insert r key
      else if key < k then insert l key
      else Node(l, k, r)

let rec delete bst key =
  match bst with
    | Leaf -> bst
    | Node(l, k, r) ->
      if key > k then delete r key (* look in the right tree *)
      else if key < k then delete l key (* look in the left tree *)
      else match (l, r) with (* found our node *)
      | (Leaf, Leaf) -> Leaf (* no children. node becomes leaf *)
      | (Leaf, r) -> r (* one child on the right. node becomes right child *)
      | (l, Leaf) -> l (* one child on the left. node becomes left child *)
      | (l, r) ->
      (* Node has both left and right children, find the successor in the right subbst *)
      let successor = minimum r in
      Node (l, successor, delete r successor)
  and minimum bst =
    match bst with
    | Leaf -> failwith "no successor"
    | Node (Leaf, k, _) -> k  (* Found the successor (minimum key) *)
    | Node (l, _, _) -> minimum l (* Continue searching in the left subbst *) 

let rec find bst key = 
  match bst with 
    | Leaf -> false
    | Node(l, k, r) -> 
      if key > k then find r key
      else if key < k then find l key
      else true 

let rec traverse bst printer =
  match bst with
    | Leaf -> print_string "leaf"
    | Node(l, k, r) -> 
      traverse l printer;
      printer k;
      traverse r printer