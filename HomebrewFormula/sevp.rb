class Sevp < Formula
  desc "A lightweight TUI for seamlessly switching environment variable values."
  homepage "https://github.com/masamerc/sevp"
  version "1.0.2"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.2/sevp_1.0.2_darwin_arm64.tar.gz"
      sha256 "f6ee4c99c17ffc7f370ef8a2c2aa1b3603980d6514342e78c778f34f8be179c4"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.2/sevp_1.0.2_darwin_amd64.tar.gz"
      sha256 "b95ac38181319fadfe569ea9d7571e972c71404cacee49888483430e21879435"

      def install
        bin.install "sevp"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.2/sevp_1.0.2_linux_arm64.tar.gz"
      sha256 "f94eb07a52ef0c5c1179399581866f68cfd04e2081993e026ea7a0539c957d8b"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.2/sevp_1.0.2_linux_amd64.tar.gz"
      sha256 "a2d1055de9d12d2019be1dafb05af588ab1bc7df0631b27866ec61612d09a6d6"

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
