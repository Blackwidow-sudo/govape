/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/blackwidow-sudo/govape/calculators/fill"
	"github.com/blackwidow-sudo/govape/validation"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// fillCmd represents the fill command
var fillCmd = &cobra.Command{
	Use:   "fill",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		model := struct {
			HaveVpg      string
			HaveNicotine string
			WantAroma    string
			WantNicotine string
		}{}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Used VPG").
					Prompt("ml: ").
					Value(&model.HaveVpg).
					Validate(validation.IsFloat),
				huh.NewInput().
					Title("Used nicotine").
					Prompt("mg/ml: ").
					Value(&model.HaveNicotine).
					Validate(validation.IsFloat),
			),
			huh.NewGroup(
				huh.NewInput().
					Title("Desired Aroma").
					Prompt("%: ").
					Value(&model.WantAroma).
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

		inputs := fill.Inputs{
			HaveVpg:      toFloat(model.HaveVpg),
			HaveNicotine: toFloat(model.HaveNicotine),
			WantAroma:    toFloat(model.WantAroma),
			WantNicotine: toFloat(model.WantNicotine),
		}

		recipe, err := inputs.Calculate()
		if err != nil {
			log.Fatal(err)
		}

		tbl := renderTable("Recipe", [][]string{
			{"Quantity", fmt.Sprintf("%.2fml", recipe.Quantity)},
			{"VPG", fmt.Sprintf("%.2fml", recipe.Vpg)},
			{"NicotineBase", fmt.Sprintf("%.2fml", recipe.NicotineBase)},
			{"Aroma", fmt.Sprintf("%.2fml", recipe.Aroma)},
		})

		fmt.Println(tbl)
	},
}

func init() {
	rootCmd.AddCommand(fillCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fillCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fillCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
