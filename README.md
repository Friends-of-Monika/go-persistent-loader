# Go Persistent Loader

This project is an experiment on loading/deserializing
[Monika After Story](https://github.com/monika-after-story/monikamoddev) persistent (save) file into memory.

Currently it contains only the most basic wrapping types just to make it work plus some typed mappings (see [renpy_persistent_persistent.go](https://github.com/Friends-of-Monika/go-persistent-loader/blob/master/persistent/types/renpy_persistent_persistent.go)
and [renpy_preferences_preferences.go](https://github.com/Friends-of-Monika/go-persistent-loader/blob/master/persistent/types/renpy_preferences_preferences.go).)

This is proven to work, however, there is a possibility (~~however, this hasn't ever happened on testing stage~~ and it
happened once) some of these mappings may fail to apply due to type incompatibilities between dynamically typed Python
and statically typed Go.

For this case and also for impossibility to know the value type for sure there are two alias types to `interface{}`,
`any` and `unknown`. For clarity, fields with type that was not yet seen as anything other than `nil` are
typed `unknown`
and fields that are proven to be dynamically typed are typed `any`.

## Example

```go
import (
    "github.com/friends-of-monika/go-persistent-loader"
)

func main() {
    save, err := persistent.Load("/home/user/.renpy/Monika After Story/persistent")
    if err != nil {
        panic(err)
    }
    
    fmt.Println(save.Playername)
}
```