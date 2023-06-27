package task

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)

type WaitGrpTask struct {
	RepertoireSource   string
	RepertoireDest     string
	Filtre             Filter
	NombreTravailleurs int
}

// NewWaitGrpTask : crée une nouvelle instance de WaitGrpTask avec les paramètres spécifiés.
func NewWaitGrpTask(repertoireSource, repertoireDest string, filtre Filter) *WaitGrpTask {
	return &WaitGrpTask{
		RepertoireSource:   repertoireSource,
		RepertoireDest:     repertoireDest,
		Filtre:             filtre,
		NombreTravailleurs: 1, // Par défaut, 1 travailleur
	}
}

// Process exécute le traitement des images en parallèle en utilisant WaitGroup.
// Chaque image est traitée individuellement et le temps d'exécution est affiché.
func (t *WaitGrpTask) Process() error {
	listeFichiers, err := ioutil.ReadDir(t.RepertoireSource)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, fichierInfo := range listeFichiers {
		if !fichierInfo.IsDir() {
			wg.Add(1)
			go func(nomFichier string) {
				defer wg.Done()

				cheminSource := filepath.Join(t.RepertoireSource, nomFichier)
				cheminDest := filepath.Join(t.RepertoireDest, nomFichier)

				heureDebut := time.Now() // Mesure de l'heure de début pour chaque image

				err := t.Filtre.Process(cheminSource, cheminDest)
				if err != nil {
					fmt.Printf("Erreur pendant du traitement de l'image %s : %s\n", nomFichier, err)
					return
				}

				tempsEcoule := time.Since(heureDebut) // Calcul du temps écoulé pour chaque image
				fmt.Printf("Durée d'exécution pour l'image %s : %s\n", nomFichier, tempsEcoule)
			}(fichierInfo.Name())
		}
	}

	wg.Wait()

	return nil
}
