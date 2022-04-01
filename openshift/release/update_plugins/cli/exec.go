package cli

import (
	"context"
	"hash/crc32"
	"io"
	"os"

	"github.com/spf13/cobra"
)

type Command struct {
	Out     io.Writer
	Exit    func(code int)
	Args    []string
	Context context.Context
}

func (c Command) Execute() {
	err := c.cobraCommand().ExecuteContext(c.context())
	c.reportError(err)
}

func (c Command) cobraCommand() *cobra.Command {
	cmd := root()
	if c.Out != nil {
		cmd.SetOut(c.Out)
		cmd.SetErr(c.Out)
	}
	if len(c.Args) > 0 {
		cmd.SetArgs(c.Args)
	}
	return cmd
}

func (c Command) context() context.Context {
	if c.Context != nil {
		return c.Context
	}
	return context.Background()
}

func (c Command) reportError(err error) {
	if err == nil {
		return
	}
	fn := os.Exit
	if c.Exit != nil {
		fn = c.Exit
	}
	fn(calcRetcode(err))
}

func calcRetcode(err error) int {
	if err == nil {
		return 0
	}
	return int(crc32.ChecksumIEEE([]byte(err.Error())))%254 + 1
}
