# 0. Getting started

## Setup

### Installing go

1. Download go binary from here: [https://golang.org/dl](https://golang.org/dl/) (please install a version `>= 1.11`)
2. Follow the install instructions here: [https://golang.org/doc/install](https://golang.org/doc/install#install)
3. (Optional) Install go extension for visual studio code: [https://code.visualstudio.com/docs/languages/go](https://code.visualstudio.com/docs/languages/go)

### `GOPATH` and go modules

This project uses go modules so you can clone it to wherever you like, it does _not_ have to been in your `GOPATH`. If you want to have it in your `GOPATH` anyway, you might have to set the environment variable `GO111MODULE=on` depending on your go version.

## Structure

This workshop is split up into 10 steps, 9 of which involve writing code. Each folder contains all the code written in the previous steps and the changes added in that step.

Each step comes with a README that gives some more background on what will happen in that step and which new concepts are being introduced. You will also find links to read up on all of these concepts. Whenever new concepts are introduced in the code there will be a comment that tells you how you could do the same thing in python.

If you are doing this workshop in a live session, we recommend creating an empty folder in here and trying to follow along from scratch. Should you get lost in-between, just copy all files from the respective step into your folder and jump right back in.
