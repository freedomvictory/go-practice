# GO MOD 


## Instruction 

Refer the link [GO BLOG](https://blog.golang.org/using-go-modules)


## Basic use  

Open the Dir [go-module-test](../go-module-test) ***(outside `$PATH/src`)***

1. Create a new module 

    - construct test project `mkdir -p ./go-module-test && touch ./go-module-test/{hello.go,hello_test.go}`
    - edit file 
    ```go
    //hello.go

    package hello
    func Hello() string {
        return "Hello, world."
    }

    ```

    ```go
    //hello_test.go 
     
    package hello
    import "testing"

    func TestHello(t *testing.T) {
        want := "Hello, world."
        if got := Hello(); got != want {
            t.Errorf("Hello() = %q, want %q", got, want)
        }
    }
    ```

    - open the `go-module-test` direcotry, and run `go test`

    ```bash

    $ go test
    PASS
    ok  	_/home/gopher/hello	0.020s
    $
    ```
  
    - make the current directory the root of a module by using `go mod init`

    ```bash
    $ go mod init example.com/hello
 
    ```




    



2. Adding a denpendency 



3. Upgrading dependencies














##