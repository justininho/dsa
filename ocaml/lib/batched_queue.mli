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