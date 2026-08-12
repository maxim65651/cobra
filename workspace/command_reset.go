package cobra

import (
	"context"
	"reflect"

	"github.com/spf13/pflag"
)

// ResetFlags resets all flags (both local and persistent) of a command and its subcommands back to their default values.
func (c *Command) ResetFlags() {
	if c.flags != nil {
		c.flags.VisitAll(func(f *pflag.Flag) {
			v := reflect.ValueOf(f.Value)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				changedField := v.FieldByName("changed")
				if changedField.IsValid() && changedField.CanSet() && changedField.Kind() == reflect.Bool {
					changedField.SetBool(false)
				}
			}
			f.Value.Set(f.DefValue)
			f.Changed = false
		})
	}
	if c.persistentFlags != nil {
		c.persistentFlags.VisitAll(func(f *pflag.Flag) {
			v := reflect.ValueOf(f.Value)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				changedField := v.FieldByName("changed")
				if changedField.IsValid() && changedField.CanSet() && changedField.Kind() == reflect.Bool {
					changedField.SetBool(false)
				}
			}
			f.Value.Set(f.DefValue)
			f.Changed = false
		})
	}
	for _, cmd := range c.Commands() {
		cmd.ResetFlags()
	}
}

// ExecuteAndReset executes the command and automatically resets all flags
// to their default values before execution, preventing value leakage
// across repeated Execute() invocations.
func (c *Command) ExecuteAndReset() error {
	c.ResetFlags()
	return c.Execute()
}

// ExecuteContextAndReset executes the command with context and automatically resets
// all flags to their default values before execution.
func (c *Command) ExecuteContextAndReset(ctx context.Context) error {
	c.ResetFlags()
	return c.ExecuteContext(ctx)
}
