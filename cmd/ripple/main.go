package main

import (
	"github.com/utkarshkrsingh/ripple/internal/cmd"
	"github.com/utkarshkrsingh/ripple/internal/log"
)

func main() {
    log.InitLogger()
    cmd.Execute()
}
