{ pkgs, ... }:

{
  # https://devenv.sh/packages/
  packages = with pkgs; [ air sqlc go-migrate go-task ];
}
