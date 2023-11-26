open OUnit2
open Datastructures.Stack

let stack_tests () =
  let test_create _ =
    let stack = create () in
    assert_equal None (peek stack);
    assert_equal 0 (size stack)
  in

  let test_push _ =
    let stack = create () in
    push stack 1;
    push stack 2;
    push stack 3;
    assert_equal (Some 3) (pop stack);
    assert_equal (Some 2) (pop stack);
    assert_equal (Some 1) (pop stack);
    assert_equal None (pop stack);
    assert_equal None (peek stack);
    assert_equal 0 (size stack)
  in

  let test_pop_empty _ =
    let stack = create () in
    assert_equal None (pop stack);
    assert_equal None (peek stack);
    assert_equal 0 (size stack)
  in

  "Lib.Stack_tests"
  >::: [
         "test_create" >:: test_create;
         "test_push" >:: test_push;
         "test_pop_empty" >:: test_pop_empty;
       ]

let () = run_test_tt_main (stack_tests ())
