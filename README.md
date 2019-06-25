# GO
***
# Description
* This repository contains my **GO** codes.
* You need to install **Go Distribution** to your system. Please follow the [original source].
# Build
* You need **GO** on your system.
* Clone the repository:  
  * `git clone https://github.com/mertturkmenoglu/golang.git`
* After getting repository on your local environment, a simple compile process:  
```
    cd golang
    cd E001_HelloWorld
    go build main.go
    ./main
```  
* You can use `afgenerator.jar` file to create auto file generation.
* Simple process:
  * Add an alias:
    * `vim ~/.bashrc`
    * Add this line: `alias gen="./afgenerator.jar"`
    * Run this command: `source ~/.bashrc`
  * Program needs `settings.json` in the directory. When you open it first time, it will ask you necessary parameters and create the file.
  * Run the program via terminal with arguments:
    * `gen 123 GoTest`: Create folder
    * `gen -s`: Update settings
***
* **To build**: `go build fileName.go`
* **To run**: `go run fileName.go`
***
# Contributing
* All ideas and helps are welcome. Please contact with me.
# Authors
* Mert Türkmenoğlu

[original source]: https://golang.org/dl/