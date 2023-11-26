open OUnit2
open Datastructures.Bst

let test_empty _ =
  let t = empty () in
  assert_bool "empty tree" (is_empty t)

let test_insert _ =
  let t = insert (empty ()) 5 in
  assert_bool "non-empty tree" (not (is_empty t));
  assert_equal (root t) (Some 5);
  assert_bool "tree contains inserted element" (find t 5)

let test_delete _ =
  let t = List.fold_left insert (empty ()) [ 5; 3; 7; 1; 6 ] in
  let t' = delete t 3 in
  assert_bool "tree does not contain deleted element" (not (find t' 3))

let test_suite =
  "BST Test Suite"
  >::: [
         "test_empty" >:: test_empty;
         "test_insert" >:: test_insert;
         "test_delete" >:: test_delete;
       ]

let _ = run_test_tt_main test_suite
