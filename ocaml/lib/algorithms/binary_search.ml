(** [binary_search v arr] returns the index of [v] in [arr] if it exists,
    and raises [Not_found] otherwise. [arr] must be sorted in ascending order
    and contain no duplicates. *)
let binary_search v arr =
  let rec search v lo hi =
    let mid = lo + (hi - lo) / 2 in
    if lo > hi then
      raise Not_found
    else if v > arr.(mid) then
      search v (mid + 1) hi
    else if v < arr.(mid) then
      search v lo mid
    else mid
  in
  let lo = 0 in 
  let hi = Array.length arr - 1 in
  search v lo hi

