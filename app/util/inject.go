package util

import (
	"testing"

	"github.com/etcinit/central/database"
	"github.com/etcinit/ohmygorm"
	"github.com/facebookgo/inject"
)

// TestingLogger represents something capable of logging while testing
type TestingLogger interface {
	Logf(format string, args ...interface{})
}

// InjectLogger provides functions for debugging dependency injection
type InjectLogger struct {
	T TestingLogger
}

// Debugf prints dependency injection debug information
func (i *InjectLogger) Debugf(f string, v ...interface{}) {
	i.T.Logf(f, v...)
}

// PopulateTest prepares a DI graph for testing.
func PopulateTest(t *testing.T, objects ...*inject.Object) *inject.Graph {
	g := inject.Graph{
		Logger: &InjectLogger{T: t},
	}

	if err := g.Provide(objects...); err != nil {
		t.Fatal(err)
	}

	if err := g.Populate(); err != nil {
		t.Fatal(err)
	}

	return &g
}

// PopulateDatabaseTest is similar to PopulateTest but it also sets up an
// in-memory database for testing.
func PopulateDatabaseTest(t *testing.T, objects ...*inject.Object) *inject.Graph {
	migrator := ohmygorm.DirectoryMigratorService{}

	objects = append(
		objects,
		&inject.Object{Value: &database.Directory{}},
		&inject.Object{Value: &migrator},
	)

	g := PopulateTest(t, objects...)

	migrator.Run()

	return g
}
