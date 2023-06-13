type 'a node = { value: 'a; next: 'a node option }

type 'a t = {
  mutable head: 'a node option;
  mutable size : int;
}
let create () =
  { head = None; size = 0 }

let push stack value =
  stack.size <- stack.size + 1;
  let new_node = Some({ value; next = stack.head }) in
  stack.head <- new_node

let pop stack =
  match stack.head with
    | None -> None
    | Some(node) ->
      stack.size <- stack.size - 1;
      stack.head <- node.next;
      Some(node.value)

let peek stack =
  match stack.head with
    | None -> None
    | Some(node) -> Some(node.value)

let size stack = stack.size

let to_list stack = 
  let rec append list node =
    match node.next with
    | None -> node.value :: list
    | Some(next) -> node.value :: append list next in
  match stack.head with
  | None -> []
  | Some(node) -> append [] node
