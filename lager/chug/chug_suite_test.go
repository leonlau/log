package chug_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestChug(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chug Suite")
}
