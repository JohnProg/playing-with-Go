package error

import (
    "log"
)

func H(err error) {
    if err != nil {
        log.Fatal("error:", err)
        panic(err);
    }
}