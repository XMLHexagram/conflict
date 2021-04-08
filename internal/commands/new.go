package commands

import (
	"fmt"
	"github.com/lmx-Hexagram/conflict/pkg"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

var NewNewProjectCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a conflict template",
	Long:  "",
	RunE:  newProject,
}

func newProject(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: project name is required.\033[m Example: conflict new helloworld\n")
		return nil
	}
	createpath, err := filepath.Abs(filepath.Clean(args[0]))
	pkg.Replaces[0].After = args[0]
	if err != nil {
		return err
	}
	err = pkg.GenerateDir("template", createpath)
	if err != nil {
		return err
	}

	c := exec.Command("go", "mod", "init", args[0])
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Dir = createpath
	err = c.Run()
	if err != nil {
		return err
	}

	c = exec.Command("go", "mod", "tidy")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Dir = createpath
	err = c.Run()
	if err != nil {
		return err
	}
	return nil
}
