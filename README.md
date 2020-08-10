# chi-gorm
Experimental project using Chi and GORM

## Starting the server


```
# install yolo - for hot reloading
go get github.com/azer/yolo

# starts development mode
make start
```

## Folder structure

Structure is inspired by https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/;
accepting only the principles that I like and getting rid of other notions that
I deem less important.

```
/
|- controllers
|- config
|- dal
|- interfaces (implemented in the future)
|- lib
|- models
|- services
main.go
router.go
```

**Controllers** - hosts the handlers of all requests coming in. Business logic
and data access layer should be done separately. This is done to decouple the
system.

**Config** - hosts setup for the sytem to connect to external data source or
any other things related to the infrastructure of the sytem. Hosts things like
database connection configurations.

**DAL** - data access layer. All queries and data operations should happen here.

**lib** - arbitrary libraries that do not fit in the other directories.

**Interfaces** - the bridge that allows two different entities to interact
with each other. For example, if the controller wants to talk to services,
UserController will implement IUserService ("I" indicating interface), and
in turn IUserService will be injected with UserService. This will be implemented
in the future since Dependency Injection requires a lot of operations overhead.

**models** - contains structs reflecting our data object from database.

**services** - where the business logic should live. Handles controller
requests and fetch data from data access access layer.

**main.go** - entry point of our system.

**router.go** - where we bind controllers to routes.
