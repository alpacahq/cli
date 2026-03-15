package cmdutil

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/alpacahq/cli/internal/api"
	"github.com/spf13/cobra"
)

// FlagOpts configures RegisterFlags behavior. All map keys use OAS names.
type FlagOpts struct {
	Defaults map[string]string // OAS name → default value override
}

// RegisterFlags registers CLI flags from generated FlagDef definitions.
// It also stores the operation name (from FlagDef.OpName) as a Cobra
// annotation so --schema can look it up at runtime, and wraps the
// command's Args validator to allow --schema to bypass arg checks.
func RegisterFlags(cmd *cobra.Command, defs []api.FlagDef, opts *FlagOpts) {
	if len(defs) > 0 && defs[0].OpName != "" {
		if cmd.Annotations == nil {
			cmd.Annotations = map[string]string{}
		}
		cmd.Annotations["op"] = defs[0].OpName

		if origArgs := cmd.Args; origArgs != nil {
			cmd.Args = func(cmd *cobra.Command, args []string) error {
				if v, _ := cmd.Flags().GetBool("schema"); v {
					return nil
				}
				return origArgs(cmd, args)
			}
		}
	}

	for _, d := range defs {
		name := d.Name
		defaultVal := d.Default
		if opts != nil {
			if def, ok := opts.Defaults[d.OASName]; ok {
				defaultVal = def
			}
		}

		switch d.Type {
		case "bool":
			cmd.Flags().Bool(name, defaultVal == "true", d.Description)
		case "int":
			defInt := 0
			if defaultVal != "" {
				var err error
				defInt, err = strconv.Atoi(defaultVal)
				if err != nil {
					log.Fatalf("RegisterFlags: invalid int default %q for flag %q: %v", defaultVal, name, err)
				}
			}
			cmd.Flags().Int(name, defInt, d.Description)
		default:
			cmd.Flags().String(name, defaultVal, d.Description)
		}

		if len(d.Completions) > 0 {
			_ = cmd.RegisterFlagCompletionFunc(name, cobra.FixedCompletions(d.Completions, cobra.ShellCompDirectiveNoFileComp))
		}
	}
}

func Str(cmd *cobra.Command, name string) string {
	v, _ := cmd.Flags().GetString(name)
	return v
}

func RequireStr(cmd *cobra.Command, name string) (string, error) {
	v := Str(cmd, name)
	if v == "" {
		return "", fmt.Errorf("--%s is required (see '%s --help' for examples)", name, cmd.CommandPath())
	}
	return v, nil
}

func RequireAll(cmd *cobra.Command, names ...string) error {
	var missing []string
	for _, n := range names {
		if Str(cmd, n) == "" {
			missing = append(missing, "--"+n)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("%s required (see '%s --help' for examples)", strings.Join(missing, ", "), cmd.CommandPath())
	}
	return nil
}

func Bool(cmd *cobra.Command, name string) bool {
	v, _ := cmd.Flags().GetBool(name)
	return v
}

func Int(cmd *cobra.Command, name string) int {
	v, _ := cmd.Flags().GetInt(name)
	return v
}

func Changed(cmd *cobra.Command, name string) bool {
	return cmd.Flags().Changed(name)
}

// PatchHelper tracks which flags were explicitly set, reducing boilerplate
// in PATCH command handlers that conditionally set body fields.
type PatchHelper struct {
	cmd     *cobra.Command
	changed bool
}

func NewPatchHelper(cmd *cobra.Command) *PatchHelper {
	return &PatchHelper{cmd: cmd}
}

func (p *PatchHelper) Str(flag string, target *string) {
	if p.cmd.Flags().Changed(flag) {
		*target = Str(p.cmd, flag)
		p.changed = true
	}
}

func (p *PatchHelper) Bool(flag string, target *bool) {
	if p.cmd.Flags().Changed(flag) {
		*target = Bool(p.cmd, flag)
		p.changed = true
	}
}

func (p *PatchHelper) Int(flag string, target *int) {
	if p.cmd.Flags().Changed(flag) {
		*target = Int(p.cmd, flag)
		p.changed = true
	}
}

func (p *PatchHelper) Strs(flag string, target *[]string) {
	if p.cmd.Flags().Changed(flag) {
		if s := Str(p.cmd, flag); s != "" {
			*target = strings.Split(s, ",")
		}
		p.changed = true
	}
}

func (p *PatchHelper) AnyChanged() bool {
	return p.changed
}
