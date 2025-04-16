class Sevp < Formula
  desc "A lightweight TUI for seamlessly switching environment variable values."
  homepage "https://github.com/masamerc/sevp"
  version "1.0.3"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.3/sevp_1.0.3_darwin_arm64.tar.gz"
      sha256 "15632f44b3c9fd06dfe3db2924add5a83a2be0fc33f4795d6aa703bf864ae925"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.3/sevp_1.0.3_darwin_amd64.tar.gz"
      sha256 "96a5f87a33bf82c487a3091dfa28ceb5491273f65d797b448aaed823ca9aa70c"

      def install
        bin.install "sevp"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.3/sevp_1.0.3_linux_arm64.tar.gz"
      sha256 "d6fb43962d607c18d6d103b135d704689dbab3786c2491ea578404840286cf32"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.3/sevp_1.0.3_linux_amd64.tar.gz"
      sha256 "b1125d06cd361c1d323da8a5791a2fa7c0ec48069d44e3fd2f432c65e156a31f"

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
