class Sevp < Formula
  desc "A lightweight TUI for seamlessly switching environment variable values."
  homepage "https://github.com/masamerc/sevp"
  version "1.0.4"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.4/sevp_1.0.4_darwin_arm64.tar.gz"
      sha256 "b9abaf7ebd5194a60c0678a96e5eacc939e1d95912aa694692f29123a53c944b"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.4/sevp_1.0.4_darwin_amd64.tar.gz"
      sha256 "da4c2198ad1cfc4a7a369ff9f821f606284b297b61f3a031daced687ec4cc8d3"

      def install
        bin.install "sevp"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.4/sevp_1.0.4_linux_arm64.tar.gz"
      sha256 "a5d07015c1264a7ee9d7b7db0578f74bcba7f0fb90c70d2ef7a0e14fa8cbed76"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.4/sevp_1.0.4_linux_amd64.tar.gz"
      sha256 "ac802686e2c6991ce6579ead0c4a33eecd83081a922f39221f0562a8dad65959"

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
