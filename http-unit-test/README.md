# Unit testing

we will be covering how to organise our tests and coverage

tests pas by default

```
go test -v
```

Subtests in Golang are a way to create hierarchical tests by defining multiple related tests under a single-parent test.

Subtests are created using the t. Run() method.
```
go test -v -run Ints/one_to_five  
```