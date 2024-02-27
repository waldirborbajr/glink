
<p align="center">
  <img width="256" height="256" src="./assets/glink.png" />
</p>

<h1 align="center">glink - GOLang symlink manager</h1>

<p align="center">
  <a href="https://github.com/waldirborbajr/glink/actions/workflows/ci-cd.yaml">
    <img alt="tests" src="https://github.com/waldirborbajr/glink/actions/workflows/ci-cd.yaml/badge.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/waldirborbajr/glink">
    <img alt="goreport" src="https://goreportcard.com/badge/github.com/waldirborbajr/glink" />
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

**BETA:** This project is in active development. Please check out the issues and contribute if you're interested in helping out.

## glink
`tl;dr:` `glink`, a.k.a GOlang Symbolic Link (symlink), is an open-source software built-in with the main aim of being a personal alternative to **GNU Stow**.

As `GNU Stow`, `glink` is a symlink farm manager which takes distinct packages of software and/or data located in separate directories on the filesystem, and makes them appear to be installed in the same place. 

With `glink` it is eeasy to track and manage configuration files in the user's home directory, especially when coupled with version control systems. 

## How to install

### Homebrew

To install glink, run the following [homebrew](https://brew.sh/) command:

```sh
brew install waldirborbajr/glink/glink
```

### Go

Alternatively, you can install glink using Go's go install command:

```sh
go install github.com/waldirborbajr/glink@latest
```

This will download and install the latest version of glink. Make sure that your Go environment is properly set up.

**Note:** Do you want this on another package manager? [Create an issue](https://github.com/waldirborbajr/glink/issues/new) and let me know!

## How to use

### glink for symblinks

```sh
glink .... #TODO
```

## .glink-ignore`

You can add files/directories to ignore list, so when execute `glink` the content will no be linked.

```sh
touch .glink-ignore
```

## Contributors

<a href="https://github.com/waldirborbajr/glink/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=waldirborbajr/glink" />
</a>

Made with [contrib.rocks](https://contrib.rocks).


