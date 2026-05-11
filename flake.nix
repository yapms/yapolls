{
	description = "YAPolls";

	inputs = {
		nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
		utils.url = "github:numtide/flake-utils";
	};

	outputs = { nixpkgs, utils, ... }: utils.lib.eachDefaultSystem(
		system: let
			pkgs = nixpkgs.legacyPackages.${system};
		in {
			devShells.default = pkgs.mkShellNoCC {
				buildInputs = [
					pkgs.nodejs_25

					pkgs.go
					pkgs.sqlc

					pkgs.turso
					pkgs.turso-cli
					pkgs.sqld
					pkgs.gcc

					pkgs.just
				];
				shellHook = ''
					echo "Welcome to the YAPolls development environment!"
				'';
			};
		}
	);
}
