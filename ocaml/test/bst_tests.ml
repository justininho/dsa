open OUnit2
open Ds.Bst

let test_empty _ =
  let t = empty () in
  assert_bool "empty tree" (is_empty t)

let test_insert _ =
  let t = insert (empty ()) 5 in
  assert_bool "non-empty tree" (not (is_empty t));
  assert_equal (root t) (Some 5);
  assert_bool "tree contains inserted element" (find t 5)

(* (* let test_insert_multiple _ =
  let t = List.fold_left insert (empty ()) [5; 3; 7; 1; 6] in
  assert_bool "non-empty tree" (not (is_empty t));
  assert_equal (root t) (Some 5);
  assert_bool "tree contains 5" (find t 5);
  assert_bool "tree contains 3" (find t 3);
  assert_bool "tree contains 7" (find t 7);
  assert_bool "tree contains 1" (find t 1);
  assert_bool "tree contains 6" (find t 6)
   *)

let test_insert_multiple _ =
  let t = List.fold_left insert (empty ()) [5; 3; 7; 1; 6] in
  List.iter (fun x -> assert_bool (Printf.sprintf "tree contains %d" x) (find t x)) [5; 3; 7; 1; 6] *)

let test_delete _ =
  let t = List.fold_left insert (empty ()) [5; 3; 7; 1; 6] in
  let t' = delete t 3 in
  assert_bool "tree does not contain deleted element" (not (find t' 3))

(* let test_traverse _ =
  let t = List.fold_left insert (empty ()) [5; 3; 7; 1; 6] in
  let l = traverse ~printer:intt in
  assert_equal l [1; 3; 5; 6; 7] *)

(* let test_traverse _ =
  let t = List.fold_left insert (empty ()) [5; 3; 7; 1; 6] in
  assert_equal (traverse t) [1; 3; 5; 6; 7] *)

let test_suite =
  "BST Test Suite" >::: [
    "test_empty" >:: test_empty;
    "test_insert" >:: test_insert;
    (* "test_insert_multiple" >:: test_insert_multiple; *)
    "test_delete" >:: test_delete;
    (* "test_traverse" >:: test_traverse *)
  ]

let _ = run_test_tt_main test_suite
