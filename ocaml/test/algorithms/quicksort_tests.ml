open OUnit2
open Algorithms.Quicksort

let test_quicksort_array _ =
  let arr = [| 5; 3; 2; 4; 1 |] in
  let result = ArraySort.sort arr in
  assert_equal [| 1; 2; 3; 4; 5 |] result

let test_quicksort_list _ =
  let list = [ 5; 3; 2; 4; 1 ] in
  let result = ListSort.sort list in
  assert_equal [ 1; 2; 3; 4; 5 ] result

let test_suite =
  "Quicksort Test Suite"
  >::: [
         "test_quicksort_array" >:: test_quicksort_array;
         "test_quicksort_list" >:: test_quicksort_list;
       ]

let _ = run_test_tt_main test_suite
