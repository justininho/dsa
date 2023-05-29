type 'a queue = { front: 'a list; back: 'a list }
  
exception Empty

let empty = { front=[]; back=[] }

let is_empty {front; _} =
  match front with
  | [] -> true
  | _ -> false

let peek {front; _} =
  match front with
  | [] -> raise Empty
  | h :: _ -> h

let enqueue value queue =
  match queue with
  | {front=[]; _;} -> {queue with front=[value]}
  | {front=_; back} -> {queue with back=value::back}

let dequeue queue =
  match queue with
  | {front=[]; _} -> raise Empty
  | {front=[_]; back} -> {front=(List.rev back); back=[]}  
  | {front=_::t; _;} -> {queue with front=t}

let size {front; back} = 
  List.length front + List.length back

let to_list {front; back} =
  front  @ List.rev back