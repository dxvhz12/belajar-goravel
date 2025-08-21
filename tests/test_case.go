package tests

import (
	"github.com/goravel/framework/testing"

	"github.com/dxvhz12/belajar-goravel/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
