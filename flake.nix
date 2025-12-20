{
  description = "wails + angular devshell";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
  let
    system = "x86_64-linux";
    pkgs = import nixpkgs { inherit system; };
  in {
    devShells.${system}.default = pkgs.mkShell {
      packages = [
        pkgs.go
        pkgs.nodejs_20

        # wails linux deps
        pkgs.pkg-config
        pkgs.gcc
        pkgs.gtk3
        pkgs.webkitgtk_4_1
        pkgs.webkitgtk_6_0
        pkgs.libsoup_3
        pkgs.glib
      ];

      shellHook = ''
        export PATH="$HOME/go/bin:$PATH"
        export PKG_CONFIG_PATH="${pkgs.webkitgtk_4_1.dev}/lib/pkgconfig:${pkgs.gtk3.dev}/lib/pkgconfig:${pkgs.glib.dev}/lib/pkgconfig";
        echo "devshell ready"
        echo "node: $(node -v)"
        echo "go:   $(go version)"
      '';
    };
  };
}

