package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoTinyCrud(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoTinyCrud Suite")
}
