package main

import (
	"flag"
	"fmt"
	"time"

	"image_filter/filter"
	"image_filter/task"
)

func main() {
	// Options de ligne de commande
	dossierEntree := flag.String("src", "input", "Dossier des images d'entrée")
	dossierSortie := flag.String("dst", "output", "Dossier des images de sortie")
	typeFiltre := flag.String("filter", "grayscale", "Type de filtre : grayscale/blur")
	typeTache := flag.String("task", "waitgrp", "Type de tâche : waitgrp/channel")
	taillePool := flag.Int("poolsize", 4, "Taille du pool de travailleurs pour la tâche de canal")
	flag.Parse()

	var filtreSelectionne task.Filter
	// Sélection du filtre en fonction du type spécifié
	switch *typeFiltre {
	case "grayscale":
		filtreSelectionne = &filter.FiltreGris{}
	case "blur":
		filtreSelectionne = &filter.FiltreFlou{}
	default:
		fmt.Println("Type de filtre invalide")
		return
	}

	var tacheSelectionnee task.Tasker
	// Sélection de la tâche de traitement en fonction du type spécifié
	switch *typeTache {
	case "waitgrp":
		tacheSelectionnee = task.NewWaitGrpTask(*dossierEntree, *dossierSortie, filtreSelectionne)
	case "channel":
		tacheSelectionnee = task.NewChanTask(*dossierEntree, *dossierSortie, filtreSelectionne, *taillePool)
	default:
		fmt.Println("Type de tâche invalide")
		return
	}

	debut := time.Now()

	err := tacheSelectionnee.Process()
	if err != nil {
		fmt.Printf("Erreur pandant le traitement des images : %s\n", err)
		return
	}

	ecoule := time.Since(debut)
	fmt.Printf("Le traitement des images a pris %s\n", ecoule)
}
