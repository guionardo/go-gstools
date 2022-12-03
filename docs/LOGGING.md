# Logging

Add logging mixin to your types.

```golang
import (
    "log"
    "github.com/guionardo/go-gstools/logger"
)

type MyStruct struct {
    logger.Logger
    myField  int
    myField2 string
}

func NewMyStruct() *MyStruct {
    myStruct = &MyStruct{}
    // You can automatically setup the logging
    myStruct.AutoSetup("MyStruct")
    // or you can declare your logging settings
    myStruct.SetDebug(true).SetLogging(log.Default).UseColors(false)

    // Use the logging features
    myStruct.Infof("Starting %d", myStruct.myField)

}
```

```bash
2022/12/03 09:23:49 myStruct [INFO] Starting 0
```