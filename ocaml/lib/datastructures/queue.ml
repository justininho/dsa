module type Queue = sig
	type 'a queue
	(** An ['a queue] is a queue whose elements have type ['a]. *)

	exception Empty
	(** Raised if [front] or [dequeue] is applied to the empty queue. *)

	val empty : 'a queue
	(** [empty] is the empty queue. *)

	val is_empty : 'a queue -> bool
	(** [is_empty q] is whether [q] is empty. *)

	val peek : 'a queue -> 'a
	(* [peek q] retreives the first item in the queue but does not remove it.
		Raises [Empty] if [q] is empty. *)

	val enqueue : 'a -> 'a queue -> 'a queue
	(** [enqueue x q] is the queue [q] with [x] added to the end. *)

	val dequeue : 'a queue -> 'a queue
	(** [dequeue q] is the queue containing all the elements of [q] except the
			front of [q]. Raises [Empty] is [q] is empty. *)

	val size : 'a queue -> int
	(** [size q] is the number of elements in [q]. *)
	
	val to_list : 'a queue -> 'a list
	(** [to_list q] is a list containing the elements of [q] in order from
			front to back. *)
end

module ListQueue : Queue = struct
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
end

module BatchedQueue : Queue = struct
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
end