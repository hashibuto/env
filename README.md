# env
environment config loader

Initialize a struct from the process environment.

```
type Example struct {
    NumItems    int     `env:PROC_NUM_ITEMS,33`     // With a default value of 33
    KeyName     string  `env:PROC_KEY_NAME`         // With no default
    EnableThing bool    `env:PROC_ENABLE_THING,f`   // Accepts true,t,1 for true and false,f,0 for false
}

config := &Example{}
err := env.Initialize(config)
```