open OUnit2
open Algorithms.Binary_search

let test_found _ =
  let arr = [| 0; 1; 2; 3; 4; 5 |] in
  let value = 2 in
  let result = binary_search value arr in
  assert_equal result 2

let test_not_found _ =
  let arr = [| 0; 1; 2; 3; 4; 5 |] in
  let value = 6 in
  assert_raises Not_found (fun () -> binary_search value arr)

let test_suite =
  "Binary Search Test Suite"
  >::: [ "test_found" >:: test_found; "test_not_found" >:: test_not_found ]

let _ = run_test_tt_main test_suite
