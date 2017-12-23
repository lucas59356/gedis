# GEDIS: Like redis but not redis

[![Build Status](https://travis-ci.org/lucas59356/gedis.svg?branch=master)](https://travis-ci.org/lucas59356/gedis)

GEDIS is a in-memory database that works by basic commands and can be easily integrated with golang. It haves an api if I, You or somebody want to integrate externally

# Basic use example
```golang
package main

import (
    "github.com/lucas59356/gedis/core"
)

func init() {
    DB := core.NewThread()
    DB.Set("test", true) // It returns the same value as interface{}, int8 that represents the data type and a error if any
    DB.Get("test") // It returns the same values as DB.Get()
    DB.Del("test") // Deletes the key in the DB, if the system need ram the space will be unallocated. More information about ram release see /core/gcman.go
    // You can create infinite instances of the DB, the limit is your hardware.
    // Basic supported types: DB.Set() can guess the type of the interface, if not one of these it will return an error and tp will be core.TypeWhatever ("*")
    tp := core.Types[core.TypeBool] // tp will be "bool"
    tp := core.Types[core.TypeString] // tp will be "string"
    tp := core.Types[core.TypeInt] // tp will be "int"
}
```

# TODO
- [ ] More types support for DB, API support too
- [ ] More benchmark tests
- [ ] Improve core/gcman.go
- [ ] Persistence support (backup). If it stops all the data is deleted
- [ ] Wrappers for the api
- [ ] Master/slave support for scalability
- [ ] Docker support