type ('k, 'v) t = {
  hash : 'k -> int;
  mutable size : int;
  mutable buckets : ('k * 'v) list array
}

(** Efficiency: O(n) *)
let create hash n =
  {hash; size = 0; buckets = Array.make n []}

(** [capacity map] is the number of buckets in [map].
    Efficiency: O(1) *)
let capacity {buckets; _;} =
  Array.length buckets

(** [index k map] is the index at which key [k] should be stored in the
      buckets of [map].
      Efficiency: O(1) *)
let index k map =
  (map.hash k) mod (capacity map)

(** [load_factor map] is the load factor of [map], i.e., the number of
    bindings divided by the number of buckets. *)
let load_factor map =
  float_of_int map.size /. float_of_int (capacity map)

(** [insert_no_resize k v map] inserts a binding from [k] to [v] in [map]
    and does not resize the table, regardless of what happens to the
    load factor.
    Efficiency: expected O(L) *)
let insert_no_resize key value map =
  let idx = index key map in
  let bucket = map.buckets.(idx) in
  map.buckets.(idx) <- (key, value) :: (List.remove_assoc key bucket);
  if not (List.mem_assoc key bucket) then
    map.size <- map.size + 1
  
 (** [rehash map new_capacity] replaces the buckets array of [map] with a new
      array of size [new_capacity], and re-inserts all the bindings of [map]
      into the new array.  The keys are re-hashed, so the bindings will
      likely land in different buckets.
      Efficiency: O(n), where n is the number of bindings. *)
let rehash map new_capacity =
  (* insert [(k, v)] into [map] *)
  let rehash_binding (k, v) =
    insert_no_resize k v map
  in
  let rehash_bucket bucket =
    List.iter rehash_binding bucket
  in
  let old_buckets = map.buckets in
  map.buckets <- Array.make new_capacity [];
  map.size <- 0;
  (* [rehash_binding] is called by [rehash_bucket] once for every binding *)
  Array.iter rehash_bucket old_buckets (* expected O(n) *)

(* [resize_if_needed map] resizes and rehashes [map] if the load factor
  is too big or too small.  Load factors are allowed to range from
  1/2 to 2. *)
let resize_if_needed map =
  let lf = load_factor map in
  if lf > 2.0 then
    rehash map (capacity map * 2)
  else if lf < 0.5 then
    rehash map (capacity map / 2)

(** Efficiency: O(n) *)
let insert k v map =
  insert_no_resize k v map; (* O(L) *)
  resize_if_needed map (* O(n) *)

(** [remove_no_resize k map] removes [k] from [map] and does not trigger
    a resize, regardless of what happens to the load factor.
    Efficiency: expected O(L) *)
let remove_no_resize k map =
  let idx = index k map in 
  let bucket = map.buckets.(idx) in
  map.buckets.(idx) <- List.remove_assoc k bucket;
  if List.mem_assoc k bucket then 
    map.size <- map.size - 1

(** Efficiency: O(n) *)
let remove k map =
  remove_no_resize k map; (* O(L) *)
  resize_if_needed map (* O(n) *)

(** Efficiency: expected O(L) *)
let find key map =
  List.assoc_opt key map.buckets.(index key map)

(* Manual implementation *)
(* let find key map =
  let rec search_bucket k bucket =
    match bucket with
      | [] -> None
      | (k', v) :: [] when k = k' -> Some(v)
      | (k', v) :: t -> 
        if k = k' then Some(v) else search_bucket k t in
  let idx = (map.hash key) mod map.size in
  let bucket = map.buckets.(idx) in
    match bucket with
    | [] -> None
    | _ -> search_bucket key bucket *)

(** [bindings m] is an association list containing the same bindings
  as [m]. *)
let bindings map =
  Array.fold_left
    (fun acc bucket ->
      List.fold_left 
        (fun acc (k, v) -> (k, v) :: acc)
        acc bucket)
    [] map.buckets

(** Efficiency: O(n^2) *)
let of_list hash lst =
  let m = create hash (List.length lst) in  (* O(n) *)
  List.iter (fun (k, v) -> insert k v m) lst; (* n * O(n) is O(n^2) *)
  m
