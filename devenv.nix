{ pkgs, ... }:

{
  # https://devenv.sh/basics/
  env.GREET = "Datastructures and Algorithms Shell";

  # https://devenv.sh/packages/
  packages = [ pkgs.vscode ];

  # https://devenv.sh/languages/
  languages.ocaml.enable = true;
  languages.rust.enable = true;
  languages.swift.enable = true;
  languages.zig.enable = true;
  languages.c.enable = true;
  languages.go.enable = true;


  # See full reference at https://devenv.sh/reference/options/
}
