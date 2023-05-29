type 'a node
type 'a t
val create : unit -> 'a t
val push : 'a t -> 'a -> unit
val pop : 'a t -> 'a option
val peek : 'a t -> 'a option
val size : 'a t -> int
val to_list : 'a t -> 'a list