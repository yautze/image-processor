# image-processor

## Govips 
[Github](https://github.com/davidbyttow/govips)

### 要求
- [libvips](https://github.com/libvips/libvips) 8.10+
- C compatible compiler such as gcc 4.6+ or clang 3.0+
- Go 1.14+

### 依賴關係
#### MacOS

使用 [homebrew](https://brew.sh/) 安裝 vips 和 pkg-config:

```bash
brew install vips pkg-config
```

#### Ubuntu

你需要一個最近的 libvips 來使用 govips。不斷添加新的 govips 功能，以利用新的 libvips 功能。Groovy (20.10) 和 Hirsute (21.04) 存儲庫有工作版本。但是在 Focal (20.04) 上，您需要從 backports 存儲庫安裝 libvips 和依賴項：

```bash
sudo add-apt-repository -y ppa:strukturag/libde265
sudo add-apt-repository -y ppa:strukturag/libheif
sudo add-apt-repository ppa:tonimelisma/ppa
```

Then:

```bash
sudo apt -y install libvips-dev
```

#### Windows

在 Windows 上推薦的方法是通過 WSL 和 Ubuntu 使用 Govips。

如果您需要在 Windows 上本地運行 Govips，這並不困難，但需要一些努力。我們目前沒有推薦的環境或設置。Windows 也不在我們的 CI/CD 目標列表中，因此不會定期測試 Govips 的兼容性。如果您願意設置和維護健壯的 CI/CD Windows 環境，請打開 PR，我們很高興接受您的貢獻並支持 Windows 作為平台。

### 安裝
```bash
go get -u github.com/davidbyttow/govips/v2/vips
```

#### MacOS note

在 MacOS 上，如果沒有首先設置環境變量，govips 可能無法編譯：

```bash
export CGO_CFLAGS_ALLOW="-Xpreprocessor"
```
