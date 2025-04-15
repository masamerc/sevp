class Sevp < Formula
  desc "A lightweight TUI for seamlessly switching environment variable values."
  homepage "https://github.com/masamerc/sevp"
  version "1.0.1"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.1/sevp_1.0.1_darwin_arm64.tar.gz"
      sha256 "7c321486367035cece2f70edf237866d177892cd4f4d9acc550ab90b813a8ce3"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.1/sevp_1.0.1_darwin_amd64.tar.gz"
      sha256 "cd2c30f12bc5eb9bf90a8f8654ad89c94e5735f3bec5e32f024eca93c76839f7"

      def install
        bin.install "sevp"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.1/sevp_1.0.1_linux_arm64.tar.gz"
      sha256 "613388d61f074d18cc76cc2681d19e563262cb81282d94bfd60491586f8da9ea"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v1.0.1/sevp_1.0.1_linux_amd64.tar.gz"
      sha256 "a6f6167df735dc8c321f11d719ea0fe8f964e18f4558804312e7e877278c5966"

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
