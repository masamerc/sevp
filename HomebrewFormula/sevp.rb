class Sevp < Formula
  desc "A lightweight TUI for seamlessly switching environment variable values."
  homepage "https://github.com/masamerc/sevp"
  version "1.0.0"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.0/sevp_1.0.0_darwin_arm64.tar.gz"
      sha256 "fd12a0aa8e02e5422e9eec8f5c6bd6e976c434eb3d48d204c42935bd8576a9bb"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.0/sevp_1.0.0_darwin_amd64.tar.gz"
      sha256 "87dd05b00b961dd2df70c9f80da69498315608e08034f231237227bd915a4fca"

      def install
        bin.install "sevp"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.0/sevp_1.0.0_linux_arm64.tar.gz"
      sha256 "2be557e88a80cd7e720ee3f390884c8c3797c66fbd09e72d0157609c132b1675"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.0/sevp_1.0.0_linux_amd64.tar.gz"
      sha256 "ab5b512fedb4c1891e01b307d767f5ad8c87951bdbd5c3aca9ec94cec31c08eb"

      def install
        bin.install "sevp"
      end
    end
  end

  def caveats
    <<~EOS
      To get started with sevp, add the shellhook to your shell configuration:

        eval "$(sevp init <shell>)"

        for zsh:
          echo 'eval "$(sevp init zsh)"' >> ~/.zshrc

        for bash:
          echo 'eval "$(sevp init bash)"' >> ~/.bashrc

      For more details, visit the documentation:
      https://github.com/masamerc/sevp
    EOS
  end

  test do
    system "#{bin}/sevp --version"
  end
end
