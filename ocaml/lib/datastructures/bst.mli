type 'a bst
val empty : unit -> 'a bst
val is_empty : 'a bst -> bool
val root : 'a bst -> 'a option
val insert : 'a bst -> 'a -> 'a bst
val delete : 'a bst -> 'a -> 'a bst
val find : 'a bst -> 'a -> bool
val traverse : 'a bst -> 'a list
(* val traverse : 'a bst -> ('a -> unit) -> unit *)