# Appendix I: Setting up your workspace

In this appendix, you’ll find instructions for setting up your workspace and installing the tools we’ll use throughout the book.

## 1.1 Podman Desktop

Follow the instructions on the [official website](https://podman-desktop.io/docs/installation) to install Podman Desktop on your operating system.

Alternatively, you can use [Homebrew](https://brew.sh) on macOS or Linux:

```shell
brew install podman-desktop
```

On Windows, you can use the [Windows Package Manager](https://learn.microsoft.com/en-us/windows/package-manager/winget/):

```shell
winget install -e --id RedHat.Podman-Desktop
```

After the installation, configure Podman Desktop as described in the article [Podman Desktop for Java Development](https://www.thomasvitale.com/podman-desktop-for-java-development/) by Thomas Vitale.

## I.2 Visual Studio Code

Follow the instructions on the [official website](https://code.visualstudio.com/) to install Visual Studio Code on your operating system.

Alternatively, you can use [Homebrew](https://brew.sh) on macOS or Linux:

```shell
brew install visual-studio-code
```

On Windows, you can use the [Windows Package Manager](https://learn.microsoft.com/en-us/windows/package-manager/winget/):

```shell
winget install -e --id Microsoft.VisualStudioCode
```

After installing Visual Studio Code, no further configuration is requires since we'll automate the setup via Devcontainers or Flox.

## I.3 Flox

Follow the instructions on the [official website](https://flox.dev/docs/install-flox/) to install Flox on your operating system.

Alternatively, you can use [Homebrew](https://brew.sh) on macOS:

```shell
brew install flox
```
