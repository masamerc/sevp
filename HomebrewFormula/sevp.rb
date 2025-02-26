class Sevp < Formula
  desc "A lightweight TUI for seamlessly switching environment variable values."
  homepage "https://github.com/masamerc/sevp"
  version "0.0.5"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v0.0.5/sevp_0.0.5_darwin_arm64.tar.gz"
      sha256 "e57bca1325332f02b03a41a1632ee974c6e9e784be06cd99cb7086c81a40f5af"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v0.0.5/sevp_0.0.5_darwin_amd64.tar.gz"
      sha256 "f996a827bc5bac95b90ece563e1d4a48299ee68cf0a5fb892c73676ced909c48"

      def install
        bin.install "sevp"
      end
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/masamerc/sevp/releases/download/v0.0.5/sevp_0.0.5_linux_arm64.tar.gz"
      sha256 "725ea5442fa932eb2d8458f51f90ade00c6ca68917170a078bc9bc5a6666e2f4"

      def install
        bin.install "sevp"
      end
    end
    on_intel do
      url "https://github.com/masamerc/sevp/releases/download/v0.0.5/sevp_0.0.5_linux_amd64.tar.gz"
      sha256 "cf8dcd7b48c7b1eac634a6a3739bc0aa18b9e8c0bc498ba3da9d4e8591900232"

      def install
        bin.install "sevp"
      end
    end
  end

  test do
    system "#{bin}/sevp --version"
  end
end
