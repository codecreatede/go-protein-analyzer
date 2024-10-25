package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-10-25

A complete package for analyzing the protein PDB files and constructing the
protein sequences. This is a part of the modeller packages that i am writing
for the protein deep learning but these indiviual packages can also be used
as stand alone packages.


*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	os.Exit(1)
}

var structurefile string

var rootCmd = &cobra.Command{
	Use:  "protein analyzer",
	Long: "reconstructing the protein from the PDB files",
	Run:  structureFunc,
}

func init() {
	rootCmd.Flags().
		StringVarP(&structurefile, "structure file", "S", "protein atom file", "pdbfile")
}

func structureFunc(cmd *cobra.Command, srgs []string) {
	proteinNames := []string{
		"ALA",
		"ARG",
		"ASN",
		"ASP",
		"CYS",
		"GLU",
		"GLN",
		"GLY",
		"HIS",
		"ILE",
		"LEU",
		"LYS",
		"MET",
		"PHE",
		"PRO",
		"SER",
		"THR",
		"TRP",
		"TYR",
		"VAL",
		"SEC",
		"PYL",
	}
	proteinCap := []string{
		"A",
		"R",
		"N",
		"D",
		"C",
		"E",
		"Q",
		"G",
		"H",
		"I",
		"L",
		"K",
		"M",
		"F",
		"P",
		"S",
		"T",
		"W",
		"Y",
		"V",
		"U",
		"O",
	}

	type proteinC struct {
		position int
		protein  string
	}

	proteinCapNames := []proteinC{}

	fOpen, err := os.Open(structurefile)
	if err != nil {
		log.Fatal(err)
	}

	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if !strings.HasPrefix(string(line), "ATOM") {
			continue
		} else {
			positionInt, _ := strconv.Atoi(string(line)[10:11])
			proteinCapNames = append(proteinCapNames, proteinC{
				position: positionInt,
				protein:  string(line)[17:21],
			})
		}
	}

	constructProtein := []string{}

	for i := range proteinCapNames {
		constructProtein = append(constructProtein, proteinCapNames[i].protein)
	}

	constructSimplified := []string{}

	for i := range constructProtein {
		line := (strings.Split(string(proteinCapNames[i].protein), ""))
		appendline := strings.Join((line[0:3]), "")
		constructSimplified = append(constructSimplified, appendline)
	}

	proteinConstruct := []string{}

	for i := range constructSimplified {
		for j := range proteinNames {
			if constructSimplified[i] == proteinNames[j] {
				proteinConstruct = append(proteinConstruct, proteinCap[j])
			}
		}
	}
	fmt.Println(">", "constructed_protein", "\n", strings.Join(proteinConstruct, ""), "\n")
}
