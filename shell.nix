{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.cmake
    pkgs.go
    pkgs.openssl
    pkgs.pipx
    pkgs.poetry
    pkgs.python3
    pkgs.rustup
    pkgs.sqlite
    pkgs.zlib
  ];

  shellHook = ''
    export OPENSSL_DIR=${pkgs.openssl.dev}
    export OPENSSL_LIB_DIR=${pkgs.openssl.dev}/lib
  '';
}

