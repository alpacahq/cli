package cmd

import (
	"fmt"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/alpacahq/cli/internal/api"
	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:   "schema [operation]",
	Short: "Show response schema for an API operation",
	Long:  "Display the JSON response fields for an API operation. Run without arguments to list all operations.",
	Example: `  alpaca schema
  alpaca schema GetAccount
  alpaca schema PostOrder
  alpaca schema News`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return listSchemas(cmd)
		}
		return showSchema(cmd, args[0])
	},
}

func listSchemas(cmd *cobra.Command) error {
	var ops []string
	for op := range api.ResponseSchemas {
		ops = append(ops, op)
	}
	sort.Strings(ops)

	w := cmd.OutOrStdout()
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "OPERATION\tFIELDS\tSUMMARY")
	for _, op := range ops {
		summary := api.OperationSummaries[op]
		fmt.Fprintf(tw, "%s\t%d\t%s\n", op, len(api.ResponseSchemas[op]), summary)
	}
	return tw.Flush()
}

func showSchema(cmd *cobra.Command, operation string) error {
	fields, ok := api.ResponseSchemas[operation]
	if !ok {
		for op, f := range api.ResponseSchemas {
			if strings.EqualFold(op, operation) {
				fields = f
				operation = op
				ok = true
				break
			}
		}
	}
	if !ok {
		return fmt.Errorf("unknown operation %q — run `alpaca schema` to list all", operation)
	}

	w := cmd.OutOrStdout()
	if summary := api.OperationSummaries[operation]; summary != "" {
		fmt.Fprintf(w, "%s — %s\n\n", operation, summary)
	}

	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "FIELD\tTYPE\tDESCRIPTION")
	for _, f := range fields {
		fmt.Fprintf(tw, "%s\t%s\t%s\n", f.Name, f.Type, f.Description)
	}
	return tw.Flush()
}
