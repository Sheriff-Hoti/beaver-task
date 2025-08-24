{
  description = "ToDo TUI APP written in Go";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-25.05";
  };

  outputs =
    {
      self,
      nixpkgs,
    }:
    let
      # System types to support
      supportedSystems = [
        "x86_64-linux"
        "aarch64-linux"
      ];

      # Helper function to generate an attrset '{ x86_64-linux = f "x86_64-linux"; ... }'
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

      # Nixpkgs instantiated for supported system types
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });

      version = "0.1.0";
      pname = "beaver-task";
    in
    {
      packages = forAllSystems (
        system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.buildGoModule {
            inherit pname version;
            src = ./.;
            vendorHash = "sha256-NjpU3luCHoy8kZ1KioFYH4yRvXPl39KtZXvEqj8o2KA=";
          };
        }
      );

      defaultPackage = forAllSystems (system: self.packages.${system}.default);

      devShell = forAllSystems (
        system:
        let
          pkgs = nixpkgsFor.${system};
        in
        with pkgs;
        mkShell {
          buildInputs = [
            go
            gopls
          ];
        }
      );

      nixosModule = forAllSystems (
        system:
        {
          config,
          lib,
          pkgs,
          ...
        }:
        with lib;
        {
          options.programs.beaver-task = {
            enable = mkEnableOption "Enable beaver-task";

            package = mkOption {
              type = types.package;
              default = self.packages.${system}.default;
              description = "beaver-task package to use";
            };
          };
        }
      );

      homeModule = forAllSystems (
        system:
        {
          config,
          lib,
          pkgs,
          ...
        }:
        with lib;
        let
          jsonFormat = pkgs.formats.json { };
        in
        {
          options.programs.beaver-task = {
            enable = mkEnableOption "Enable beaver-task";

            package = mkOption {
              type = types.package;
              default = self.packages.${system}.default;
              description = "beaver-task package to use";
            };

            settings = mkOption {
              type = lib.types.nullOr jsonFormat.type;
              default = null;
              description = "JSON configuration settings for beaver-task";
              example = {
                test = "whatever string";
                data_dir = "/home/user/.local/share/beaver-task/data.db";
              };
            };
          };

          config = mkIf config.programs.beaver-task.enable (
            let
              cfg = config.programs.beaver-task;
            in
            {
              home.packages = [ cfg.package ];

              xdg.configFile = lib.mkIf (cfg.settings != null) {
                "beaver-task/config.json".source = jsonFormat.generate "beaver-task.json" cfg.settings;
              };
            }
          );
        }
      );

      checks = forAllSystems (system: {
        beaver-task = self.packages.${system}.default;
      });
    };
}
