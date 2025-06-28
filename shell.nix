{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = with pkgs; [
		gcc
		libGL
		xorg.libX11
		xorg.libXcursor
		xorg.libXext
		xorg.libXi
		xorg.libXinerama
		xorg.libXrandr
		xorg.libXxf86vm
	];

	nativeBuildInputs = with pkgs; [
		go
	];

	shellHook = with pkgs; ''
		export LD_LIBRARY_PATH=${lib.getLib libGL}/lib:${lib.getLib libGL}/lib:$LD_LIBRARY_PATH
	'';
}
