init() 
======

Project struct:

```
├── main.go
└── pac
    ├── one.go
    ├── three.go
    └── two.go
```


`go run main.go` Output:

```
Call one init()
Call threee init()
Call two init()
Call main init()
Call main main()
I am Exportalbe!
```
