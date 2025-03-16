This fork is used to improve the recompilation speed when using the `wails dev` command for development on Windows, aiming to enhance the development experience. For detailed information, please check https://github.com/wailsapp/wails/issues/4142. Relevant changes have been submitted as pull requests to the official repository. Before the official repository merges them, you can install the Wails CLI tool modified by this fork to improve the recompilation speed by using either
```bash
go install github.com/josStorer/wails/v2/cmd/wails@v2.9.2b
```
or
```bash
go install github.com/josStorer/wails/v2/cmd/wails@v2.10.1b
```
You only need to install the CLI tool and then `wails dev -m -skipembedcreate`. There is **no** need to redirect the wails url in `go.mod` to this repository.

v2.9.2 is the last Wails version that supports Golang 1.20, which may be useful for some users developing applications on Windows 7.
v2.10.1 is the current latest Wails version.