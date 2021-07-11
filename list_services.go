package main

import (
	"context"
	"fmt"

	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/spf13/cobra"
)

type ListServicesCommand struct {
	cmd       *cobra.Command
	opts      *GlobalOptions
	userAgent string
	addr      string
	rcli      *grpcreflect.Client
	long      bool
	full      bool
}

func NewListServicesCommand(opts *GlobalOptions) *ListServicesCommand {
	c := &ListServicesCommand{
		cmd: &cobra.Command{
			Use:   "list_services ADDR [FULL_SERVICE_NAME]",
			Short: "List services and methods provided by gRPC server",
			Example: `
* List services
grpcurl ls localhost:8888

* List methods of the service
grpcurl ls localhost:8888 test.TestService
`,
			Aliases:      []string{"ls"},
			Args:         cobra.RangeArgs(1, 2),
			SilenceUsage: true,
		},
		opts: opts,
	}
	c.cmd.RunE = c.Run
	c.cmd.Flags().BoolVarP(&c.long, "long", "l", false, "list long")
	c.cmd.Flags().BoolVarP(&c.full, "full", "F", false, "fully qualified")
	c.cmd.Flags().StringVarP(&c.userAgent, "user-agent", "U", "", "user-agent")
	return c
}

func (c *ListServicesCommand) Command() *cobra.Command {
	return c.cmd
}

func (c *ListServicesCommand) Run(cmd *cobra.Command, args []string) error {
	nargs := len(args)
	ctx := context.Background()

	c.addr = args[0]
	conn, err := NewGRPCConnection(ctx, c.addr, c.userAgent, c.opts.Insecure)
	if err != nil {
		return err
	}
	defer conn.Close()
	c.rcli = NewServerReflectionClient(ctx, conn)

	if nargs == 1 {
		return c.listServices(ctx)
	} else if len(args) == 2 {
		return c.listMethods(ctx, args[1])
	}

	return nil
}

func (c *ListServicesCommand) listServices(ctx context.Context) error {
	svcs, err := c.rcli.ListServices()
	if err != nil {
		return err
	}

	for i := range svcs {
		fmt.Fprintf(c.opts.Output, "%s\n", svcs[i])
	}

	return nil
}

func (c *ListServicesCommand) listMethods(ctx context.Context, serviceName string) error {
	sdesc, err := c.rcli.ResolveService(serviceName)
	if err != nil {
		return err
	}

	for _, mdesc := range sdesc.GetMethods() {
		if c.long {
			inType := mdesc.GetInputType()
			outType := mdesc.GetOutputType()
			inRPCType := ""
			outRPCType := ""
			if mdesc.IsClientStreaming() {
				inRPCType = "streaming "
			}
			if mdesc.IsServerStreaming() {
				outRPCType = "streaming "
			}
			fmt.Fprintf(c.opts.Output, "%s(%s%s) return (%s%s)\n",
				mdesc.GetFullyQualifiedName(),
				inRPCType, inType.GetFullyQualifiedName(),
				outRPCType, outType.GetFullyQualifiedName())
		} else {
			fmt.Fprintf(c.opts.Output, "%s\n", mdesc.GetFullyQualifiedName())
		}
	}

	return nil
}
