package validate

import (
	"errors"
	"math"

	"github.com/urfave/cli"
)

// Validate validate cli argument
func Validate(c *cli.Context) error {
	if len(c.Args()) >= 2 || len(c.Args()) == 0 {
		return errors.New("required single argument")
	}
	if c.Float64("n") < 0 || c.Float64("n") != math.Trunc(c.Float64("n")) {
		return errors.New("-n option required Natural number or equal to 0")
	}
	return nil
}
