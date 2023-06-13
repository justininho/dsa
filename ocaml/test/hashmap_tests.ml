open OUnit2
open Dsa.Hashmap

let test_insert ctxt =
  let hash = Hashtbl.hash in
  let map = create hash 10 in
  insert "key1" "value1" map;
  assert_equal ~ctxt (Some "value1") (find "key1" map)

let test_find ctxt =
  let hash = Hashtbl.hash in
  let map = create hash 10 in
  assert_equal ~ctxt None (find "key1" map);
  insert "key1" "value1" map;
  assert_equal ~ctxt (Some "value1") (find "key1" map)

let test_remove ctxt =
  let hash = Hashtbl.hash in
  let map = create hash 10 in
  insert "key1" "value1" map;
  remove "key1" map;
  assert_equal ~ctxt None (find "key1" map)

let test_create ctxt =
  let hash = Hashtbl.hash in
  let map = create hash 10 in
  assert_equal ~ctxt [] (bindings map)

let test_bindings ctxt =
  let hash = Hashtbl.hash in
  let map = create hash 10 in
  insert "key1" "value1" map;
  insert "key2" "value2" map;
  assert_equal ~ctxt [ ("key1", "value1"); ("key2", "value2") ] (bindings map)

let test_of_list ctxt =
  let hash = Hashtbl.hash in
  let list = [ ("key1", "value1"); ("key2", "value2") ] in
  let map = of_list hash list in
  assert_equal ~ctxt list (bindings map)

let suite =
  "Mapmodule Tests"
  >::: [
         "test_insert" >:: test_insert;
         "test_find" >:: test_find;
         "test_remove" >:: test_remove;
         "test_create" >:: test_create;
         "test_bindings" >:: test_bindings;
         "test_of_list" >:: test_of_list;
       ]

let _ = run_test_tt_main suite
