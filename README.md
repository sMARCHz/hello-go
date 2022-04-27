# hello-go

## Go CLI

**go build** -> Compiles a bunch of go source code files  (Compiles .go to .exe)  
**go run** -> Compiles and executes one or two files  
**go fmt** -> Formats all the code in each file in the current directory  
**go install** -> Compiles and installs a package  
**go get** -> Downloads the raw source code of someone else's packages  
**got test** -> Runs any tests associated with the current project  

## Go Packages

Package is scope of the code. (Package == Project == Workspace). **A package contains many files which work together for similar purpose.**  
Go has 2 types of Packages  

1. **Executable** -> Generates a executable file that we use for running application.  
Package main is a **executable package** and **need func main()** inside it. If we use other package instead of main, nothing is going to come out when run go build.
2. **Reusable** -> Code used as 'helper' or 'dependency'. Good place to put reusable logic.

## Types
1. **Value** -> int, float, string, bool, array, struct
2. **Reference** -> slices, maps, function, pointer, channel
