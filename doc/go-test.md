# GO Test 


## using test package do `setup` or `teardown`


we can use `TestMain`. This function will be called instead of running the test . And in this function , we can define how the tests will run. For example  we can implement global `setup` and `teardown` 

```go
func TestMain(m *testing.M){
    setup()
    code := m.Run()
    shutdown()
    os.Exit() 
}
```