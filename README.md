# go-test-example

```
________________________________________   
|        _________________________       | 
|        |    _____________      |       | 
|        |    |           |      |       | 
|        |    |   Random  |      |       | 
|        |    |___________|      |       | 
|        |                       |       | 
|        |         Number        |       | 
|        |_______________________|       | 
|                                        | 
|                   Sum                  | 
|________________________________________| 
```

Dependency injection aka [inversion of control](https://en.wikipedia.org/wiki/Inversion_of_control) is probably better known as a solution to the [dependency inversion principle](https://en.wikipedia.org/wiki/Dependency_inversion_principle) which is the letter “D” in SOLID. Both principles touch on writing well-architected, modularized programs.

Dependency injection containers have proved to be useful and even though [Wire](https://github.com/google/wire/blob/main/_tutorial/README.md) seems like the most promising DIC solution in Go, right now, it does not appear to add significant value in most Go use cases. It’s also important to remember that dependency injections do not require a container. Surprisingly, many people tend to forget about this, maybe because they get used to frameworks first.

## Reference 

https://appliedgo.net/di/ <br/>
https://banzaicloud.com/blog/dependency-injection-go/
