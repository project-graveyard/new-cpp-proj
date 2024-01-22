/*
Copyright © 2024 David Saah dasorange.hope@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:          "init",
	Args:         cobra.MinimumNArgs(1),
	Short:        "Initialise a new c++ project",
	Example:      "new-cpp-proj init directory_name",
	SilenceUsage: true,
	Version:      "0.1.0",
	Long: `
Initialise a new c++ project

Argument:
  directory    the name of your project directory
`,
	Run: func(_ *cobra.Command, args []string) {
		// create a new directory
		os.Mkdir(args[0], 0755)

		// create main.cpp
		mainFile, err := os.Create(args[0] + "/main.cpp")
		if err != nil {
			panic(err)
		}

		// write to main.cpp
		_, err = mainFile.WriteString("#include <iostream>\nusing namespace std;\n\nint main() {\n\tcout << \"Hello World!\" << endl;\n\treturn 0;\n}")
		if err != nil {
			panic(err)
		}

		// create makefile
		makeFile, err := os.Create(args[0] + "/Makefile")
		if err != nil {
			panic(err)
		}

		/* Makefile template
		all: run clean

		build:
		  @g++ main.cpp -o main.out

		run: build
			@./main.out

		clean:
			@rm *.out

		debug:
			@g++ -g main.cpp -o main.out
			@gdb main.out
		*/
		// write to makeFile
		_, err = makeFile.WriteString("all: run clean\n\nbuild:\n\t@g++ main.cpp -o main.out\n\nrun: build\n\t@./main.out\n\nclean:\n\t@rm *.out\n\ndebug:\n\t@g++ -g main.cpp -o main.out\n\t@gdb main.out")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
