class Sevp < Formula
  desc "A lightweight TUI for seamlessly switching environment variable values."
  homepage "https://github.com/masamerc/sevp"
  version "1.0.2"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.2/sevp_1.0.2_darwin_arm64.tar.gz"
      sha256 "71ce318f96e0b7cbfc9b3ad68a3b51a18750f66915ae5febc79bbfdae326dd4b"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.2/sevp_1.0.2_darwin_amd64.tar.gz"
      sha256 "126f09d74d1cc340b5aec27c5022eaa0eca78ee594122bcf201478c9549d051c"

      def install
        bin.install "sevp"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.2/sevp_1.0.2_linux_arm64.tar.gz"
      sha256 "228eae1df5e49f0a70078a34a22da0e22fd0256e56b6bce78b7b6b23dea132c5"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.2/sevp_1.0.2_linux_amd64.tar.gz"
      sha256 "2f88f720027270d2309933f9fc7f3937da5cb5e5da3d2ca13d2c421bfb0f71c2"

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
