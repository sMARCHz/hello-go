# hello-go

## Go CLI

**go build** -> Compiles a bunch of go source code files  (Compiles .go to .exe)  
**go run** -> Compiles and executes one or two files  
**go fmt** -> Formats all the code in each file in the current directory  
**go install** -> Compiles and installs a package  
**go get** -> Downloads the raw source code of someone else's packages  
**go test** -> Runs any tests associated with the current project  

## Go Packages

Package is scope of the code. (Package == Project == Workspace). **A package contains many files which work together for similar purpose.**  
Go has 2 types of Packages  

1. **Executable** -> Generates a executable file that we use for running application.  
Package main is a **executable package** and **need func main()** inside it. If we use other package instead of main, nothing is going to come out when run go build.
2. **Reusable** -> Code used as 'helper' or 'dependency'. Good place to put reusable logic.

### Tips for Go Packages

1. New directory == New package
2. We can use go module to manage packages and dependencies easier  
``` go mod init <module name> ```

## Types

1. **Value** -> int, float, string, bool, array, struct
2. **Reference** -> slices, maps, function, pointer, channel

## Go != OOP
Go isn't OOP because **Go has no concept of class**. Although Go isn't OOP, we can extend the functionality of the type in Go by declaring the new type based on the  type which we want to extend from.  
``` type deck []string  ```  
According to the above code, it means the deck type has all functionilities of []string.  

## Reference
<https://www.youtube.com/watch?v=YS4e4q9oBaU>  
<https://www.udemy.com/course/go-the-complete-developers-guide/>
