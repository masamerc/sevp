class Sevp < Formula
  desc "A lightweight TUI for seamlessly switching environment variable values."
  homepage "https://github.com/masamerc/sevp"
  version "0.0.4"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v0.0.4/sevp_0.0.4_darwin_arm64.tar.gz"
      sha256 "b472820c44737e50a22f0012df6e3379878d80d79351156c6628889983ed718a"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v0.0.4/sevp_0.0.4_darwin_amd64.tar.gz"
      sha256 "62889337b9527fc76e9d45933161df6180d8b2be053e9e56cd757c6140d4207d"

      def install
        bin.install "sevp"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v0.0.4/sevp_0.0.4_linux_arm64.tar.gz"
      sha256 "5f75f958bc90351027756f013bca9ef8e370f6d6a64437e9bd33fd9c2e65c864"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v0.0.4/sevp_0.0.4_linux_amd64.tar.gz"
      sha256 "50d210e5608be012b82be35548388431e8d89f53c71a2f338c3ed464a11902e1"

      def install
        bin.install "sevp"
      end
    end
  end

  test do
    system "#{bin}/sevp --version"
  end
end
