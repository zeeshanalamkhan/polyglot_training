## Enable dependency tracking for your code
When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repository.

```go mod init lbg/hello```

Output: ``` go: creating new go.mod: module lbg/hello```

## To run the application you cam symply use 
``` go run .```
If your application has just one main program then it will automatically detect main methon (entry point) and it will execute

## If you wanted to compile and run then use below commands 
``` 
go build -o abc  
./abc
```
