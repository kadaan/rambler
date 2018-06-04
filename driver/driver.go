package driver

import (
	"fmt"

	"github.com/elwinar/rambler/env"
)

// Driver is the interface used by the program to initialize the database connection.
type Driver interface {
	New(env.Environment) (Conn, error)
}

var drivers = make(map[string]Driver)

// Register register a driver
func Register(name string, driver Driver) error {
	if _, found := drivers[name]; found {
		return fmt.Errorf(`driver "%s" already registered`, name)
	}

	if driver == nil {
		return fmt.Errorf(`not a valid driver`)
	}

	drivers[name] = driver
	return nil
}

// Get initialize a driver from the given environment
func Get(name string, environment env.Environment) (Conn, error) {
	driver, found := drivers[name]
	if !found {
		return nil, fmt.Errorf(`driver "%s" not registered`, name)
	}

	conn, err := driver.New(environment)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
