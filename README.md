# Clean Architechture - Go
A simple project using clean architechture following this diagram

<img src="https://user-images.githubusercontent.com/48635609/163501148-31074282-8f34-4a2e-8b08-cf01c6332e4f.jpeg" alt="clean architechture"/>

## Running Go project

`go run ./src/main.go`

## Installing dependencies

`go get ${dependency}`

## Go interfaces syntax

``` go
type CreateUserUseCase interface {
    CreateUser(user *entity.User) *entity.User
}
```

## Go struct syntax

``` go
type User struct {
    name string
    age uint
}
```

## Go struct method syntax

``` go
type CreateUserUseCaseImpl struct {
    createUserRepository *repository.CreateUserRepository
}

func (c *CreateUserUseCaseImpl) CreateUser(user *entity.User) *entity.User {
    ...
}
```

## Go struct constructor syntax
In Go we don't have constructors by default, but we can do that
``` go
type CreateUserUseCaseImpl struct {
    createUserRepository *repository.CreateUserRepository
}

func NewCreateUserUseCaseImpl(createUserRepository *repository.CreateUserRepository) *CreateUserUseCaseImpl {
    return &CreateUserUseCaseImpl {
        createUserRepository: createUserRepository
    }
}
```

## Go pointers

In Go we need to specify when we want a pointer the pointer is specified by the char `*`

This is a pointer of the type `entity.User`:
``` go
    *entity.User
```

This is a type `entity.User`
``` go
    entity.User
```

This is the memory reference of a pointer, specified by `&`:
``` go
   var *entity.User user = new(entity.User)
   fmt.Println(&user)
```

## Allocating a null memory address
When we will allocate the variable on the memory address in Go we need to say that
``` go
    var *entity.User user = new(entity.User)
    user := &entity.User{}
```
