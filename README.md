<p align=center>
    <img src="./.assets/ripple-logo.png">
    <p align=center>A build automation tool written in <a href="https://go.dev/">Go</a></p>
    <p align=center>
    <a href="./LICENSE.md"><img src="https://img.shields.io/github/license/utkarshkrsingh/ripple?style=flat-square&logo=appveyor"></a>
    <img src="https://img.shields.io/badge/go-1.23.5-green?style=flat-square&logo=appveyor">
    <img src="https://img.shields.io/github/issues/utkarshkrsingh/ripple?style=flat-square&logo=appveyor">
    <img src="https://img.shields.io/github/forks/utkarshkrsingh/ripple?style=flat-square&logo=appveyor">
    <img src="https://img.shields.io/github/stars/utkarshkrsingh/ripple?style=flat-square&logo=appveyor">
    </p>
</p>

### Ripple
`ripple` is a build automation tool written in [Go](https://go.dev/). Executing a simgle task can trigger multiple dependent tasks.

### Installing and Building
1. Clone this repo via the following command:
```bash
git clone https://github.com/utkarshkrsingh/ripple.git
```

2. Get the dependencies downloaded (make sure Go is installed):
```bash
cd ripple
go mod tidy
go mod download
```
<strong>Note:</strong> You can even use `Makefile` for building the binary.


3. Finally compile the project to get a binary:
```bash
go build -o ripple ./cmd/ripple
```

##### For Linux & MacOS:
To use binary system-wide, place it in `/usr/local/bin/`:
```bash
sudo mv ripple /usr/local/bin/
sudo chmod +x /usr/local/bin/ripple
```

##### For Windows:
In Windows, to use ripple system-wide, you need to place it in a directory that is included in the `PATH` environment variable.

### Running
You can run the task by using the following command:
```bash
ripple task -r <taskname>   # To run the task by name
ripple task -s              # To show all the listed task in ripple.toml
```

### Configuration
The configuration for each project can be written in `ripple.toml` file at the root of the project.
Demo config is given - [here](./Demo-Config.md)
