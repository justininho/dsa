module type Sort = sig
  type t

  val sort : t -> t
end

module ArraySort : Sort with type t = int array = struct
  type t = int array

  let swap arr i j =
    let tmp = arr.(i) in
    arr.(i) <- arr.(j);
    arr.(j) <- tmp

  let partition arr lo hi =
    let pivot = arr.(hi) in
    let i = ref lo in
    for j = lo to hi - 1 do
      if arr.(j) <= pivot then (
        swap arr !i j;
        incr i)
    done;
    swap arr !i hi;
    !i

  let rec quicksort arr lo hi =
    if lo < hi then (
      let p = partition arr lo hi in
      quicksort arr lo (p - 1);
      quicksort arr (p + 1) hi)

  let sort arr =
    quicksort arr 0 (Array.length arr - 1);
    arr
end

module ListSort : Sort with type t = int list = struct
  type t = int list

  let rec quicksort = function
    | [] -> []
    | pivot :: rest ->
        let is_less x = x < pivot in
        let left, right = List.partition is_less rest in
        let sorted_left = quicksort left in
        let sorted_right = quicksort right in
        sorted_left @ [ pivot ] @ sorted_right

  let sort = quicksort
end
