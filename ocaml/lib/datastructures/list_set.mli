type 'a t
(** ['a t] is the type of sets whos elements have type ['a] *)

val empty : 'a t
(** [empty] is the empty set *)

val mem : 'a -> 'a t -> bool
(** [mem k s] returns whether [k] is a member of [s]  *)

val insert : 'a -> 'a t -> 'a t
(** [insert k s] is the set contains [k] and all the elements
	 of [s] *)
