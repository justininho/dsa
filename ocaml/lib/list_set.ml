type 'a t = 'a list

let empty = []

let mem = List.mem

let insert key set =
  if mem key set then set else  key :: set