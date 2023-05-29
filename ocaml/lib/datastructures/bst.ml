type 'a bst = 
| Leaf
| Node of 'a bst * 'a  * 'a bst

let empty () = Leaf

let is_empty = function
| Leaf -> true
| Node _ -> false

let root = function
| Leaf -> None
| Node (_, v, _) -> Some v

let rec insert bst value =
  match bst with
    | Leaf -> Node(Leaf, value, Leaf)
    | Node(l, v, r) -> 
      if value > v then insert r value
      else if value < v then insert l value
      else Node(l, v, r)
  
let rec delete bst value =
  (* Helper function to find the successor (minimum value) in a subbst *)
  let rec find_successor bst =
    match bst with
      | Leaf -> failwith "No successor found!"
      | Node (Leaf, v, _) -> v  (* Found the successor (minimum value) *)
      | Node (l, _, _) -> find_successor l in  (* Continue searching in the left subbst *) 
  match bst with
    | Leaf -> bst
    | Node(l, v, r) ->
        if value > v then delete r value
        else if value < v then delete l value
        else match (l, r) with
          | (Leaf, Leaf) -> Leaf (* found, no children, delete *)
          | (Leaf, r) -> r (* found one child, right *)
          | (l, Leaf) -> l (* found one child, left *)
          | (l, r) ->
          (* Node has both left and right children, find the successor in the right subbst *)
          let successor = find_successor r in
          Node (l, successor, delete r successor)

let rec find bst value = 
  match bst with 
    | Leaf -> false
    | Node(l, v, r) -> 
      if value > v then find r value
      else if value < v then find l value
      else true 

let rec traverse bst printer =
  match bst with
    | Leaf -> print_string "leaf"
    | Node(l, v, r) -> 
      traverse l printer;
      printer v;
      traverse r printer