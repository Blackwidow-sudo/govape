/*
Copyright © 2024 Merlin Cornehl <bob_merlin92@web.de>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/blackwidow-sudo/govape/calculators/nicotine"
	"github.com/blackwidow-sudo/govape/validation"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// nicotineCmd represents the nicotine command
var nicotineCmd = &cobra.Command{
	Use:   "nicotine",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		model := struct {
			HaveQuantity string
			HaveNicotine string
			WantQuantity string
			WantNicotine string
		}{}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Used Base").
					Prompt("ml: ").
					Value(&model.HaveQuantity).
					Validate(validation.IsFloat),
				huh.NewInput().
					Title("Used nicotine").
					Prompt("mg/ml: ").
					Value(&model.HaveNicotine).
					Validate(validation.IsFloat),
			),
			huh.NewGroup(
				huh.NewInput().
					Title("Desired Base").
					Prompt("ml: ").
					Value(&model.WantQuantity).
					Validate(validation.IsFloat),
				huh.NewInput().
					Title("Desired nicotine").
					Prompt("mg/ml: ").
					Value(&model.WantNicotine).
					Validate(validation.IsFloat),
			),
		)

		if err := form.Run(); err != nil {
			switch err {
			case huh.ErrUserAborted:
				os.Exit(0)
			default:
				log.Fatal(err)
			}
		}

		inputs := nicotine.Inputs{
			HaveQuantity: toFloat(model.HaveQuantity),
			HaveNicotine: toFloat(model.HaveNicotine),
			WantQuantity: toFloat(model.WantQuantity),
			WantNicotine: toFloat(model.WantNicotine),
		}

		recipe, err := inputs.Calculate()
		if err != nil {
			log.Fatal(err)
		}

		tbl := renderTable("Recipe", [][]string{
			{"Quantity", fmt.Sprintf("%.2fml", recipe.Quantity)},
			{"NicotineBase", fmt.Sprintf("%.2fml", recipe.NicotineBase)},
			{"Rest", fmt.Sprintf("%.2fml", recipe.Rest)},
		})

		fmt.Println(tbl)
	},
}

func init() {
	rootCmd.AddCommand(nicotineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nicotineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nicotineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
