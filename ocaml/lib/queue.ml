type 'a queue = 'a list
exception Empty
let empty = []
let is_empty = function
  | [] -> true
  | _ -> false
let peek = function
  | [] -> raise Empty
  | h :: _ -> h
let enqueue value = function
  | [] -> [value]
  | queue -> queue @ [value]
let dequeue = function
  | [] -> raise Empty
  | _ :: t -> t
let size = List.length 
let to_list = Fun.id
