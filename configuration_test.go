package main

import (
	"reflect"
	"testing"

	"github.com/elwinar/rambler/env"
)

func TestLoad(t *testing.T) {
	var cases = []struct {
		input  string
		err    bool
		output Configuration
	}{
		{
			input:  "test/notfound.json",
			err:    true,
			output: Configuration{},
		},
		{
			input:  "test/invalid.json",
			err:    true,
			output: Configuration{},
		},
		{
			input: "test/valid.json",
			err:   false,
			output: Configuration{
				Environment: env.Environment{
					Driver:    "mysql",
					Protocol:  "tcp",
					Host:      "localhost",
					Port:      3306,
					User:      "root",
					Password:  "",
					Database:  "rambler_default",
					Directory: ".",
					Table:     "migrations",
				},
				Environments: map[string]env.Environment{
					"testing": {
						Database: "rambler_testing",
					},
					"development": {
						Database: "rambler_development",
					},
					"production": {
						Database: "rambler_production",
					},
				},
			},
		},
		{
			input: "test/valid.hjson",
			err:   false,
			output: Configuration{
				Environment: env.Environment{
					Driver:    "mysql",
					Protocol:  "tcp",
					Host:      "localhost",
					Port:      3306,
					User:      "root",
					Password:  "",
					Database:  "rambler_default",
					Directory: ".",
					Table:     "migrations",
				},
				Environments: map[string]env.Environment{
					"testing": {
						Database: "rambler_testing",
					},
					"development": {
						Database: "rambler_development",
					},
					"production": {
						Database: "rambler_production",
					},
				},
			},
		},
	}

	for n, c := range cases {
		cfg, err := Load(c.input)
		if (err != nil) != c.err {
			t.Error("case", n, "got unexpected error:", err)
			continue
		}

		if !reflect.DeepEqual(cfg, c.output) {
			t.Error("case", n, "got unexpected output: wanted", c.output, "got", cfg)
		}
	}
}

func TestConfigurationEnv(t *testing.T) {
	var cases = []struct {
		input  string
		err    bool
		output env.Environment
	}{
		{
			input:  "unknown",
			err:    true,
			output: env.Environment{},
		},
		{
			input: "default",
			err:   false,
			output: env.Environment{
				Driver:    "mysql",
				Protocol:  "tcp",
				Host:      "localhost",
				Port:      3306,
				User:      "root",
				Password:  "",
				Database:  "rambler_default",
				Directory: ".",
				Table:     "migrations",
			},
		},
		{
			input: "testing",
			err:   false,
			output: env.Environment{
				Driver:    "mysql",
				Protocol:  "tcp",
				Host:      "localhost",
				Port:      3306,
				User:      "root",
				Password:  "",
				Database:  "rambler_testing",
				Directory: ".",
				Table:     "migrations",
			},
		},
	}

	for n, c := range cases {
		cfg := Configuration{
			Environment: env.Environment{
				Driver:    "mysql",
				Protocol:  "tcp",
				Host:      "localhost",
				Port:      3306,
				User:      "root",
				Password:  "",
				Database:  "rambler_default",
				Directory: ".",
				Table:     "migrations",
			},
			Environments: map[string]env.Environment{
				"testing": {
					Database: "rambler_testing",
				},
				"development": {
					Database: "rambler_development",
				},
				"production": {
					Database: "rambler_production",
				},
			},
		}

		env, err := cfg.Env(c.input)
		if (err != nil) != c.err {
			t.Error("case", n, "got unexpected error:", err)
			continue
		}

		if !reflect.DeepEqual(env, c.output) {
			t.Error("case", n, "got unexpected output:", cfg)
		}
	}
}
