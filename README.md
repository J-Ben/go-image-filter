# Projet de traitement d'images

Ce projet a été réalisé dans le cadre de mon cours sur le traitement d'images. L'objectif était de créer un programme capable d'appliquer des filtres sur des images de manière parallèle en utilisant deux méthodes différentes : WaitGroup et Channel.

## Structure du projet

Le projet est organisé en plusieurs fichiers et dossiers :

- Le dossier `images` contient les images sources sur lesquelles les filtres seront appliqués.
- Le dossier `output` est l'emplacement où les images traitées seront enregistrées.
- Le dossier `filtre` contient les implémentations des filtres.
- Le dossier `tache` contient les implémentations des tâches de traitement.
- Le fichier `main.go` est le point d'entrée du programme.

## Filtres

Deux filtres ont été implémentés :

- Filtre de conversion en niveaux de gris (`FiltreNiveauxGris`) : Ce filtre convertit une image en niveaux de gris en utilisant la bibliothèque `github.com/disintegration/imaging`.
- Filtre de flou (`FiltreFlou`) : Ce filtre applique un flou gaussien à une image en utilisant la bibliothèque `github.com/disintegration/imaging`.

## Tâches de traitement

Deux tâches de traitement ont été mises en œuvre :

- Tâche WaitGroup (`TacheWaitGroup`) : Cette tâche utilise la synchronisation WaitGroup pour exécuter les traitements d'images en parallèle. Chaque image est traitée individuellement, et le temps d'exécution est affiché.
- Tâche Channel (`TacheChannel`) : Cette tâche utilise des goroutines et des channels pour répartir les traitements d'images en parallèle. Les images sont réparties entre plusieurs travailleurs, et le temps d'exécution est affiché.

## Utilisation du programme

Avant d'exécuter le programme, assurez-vous d'avoir les bibliothèques requises. Vous pouvez les installer en exécutant la commande suivante :

Une fois les bibliothèques installées, vous pouvez exécuter le programme en utilisant la commande suivante :

- L'option `-src` spécifie le répertoire source contenant les images à traiter.
- L'option `-dst` spécifie le répertoire de destination où les images traitées seront enregistrées.
- L'option `-filter` spécifie le type de filtre à appliquer (`niveauxgris` pour le filtre de niveaux de gris, `flou` pour le filtre de flou).
- L'option `-task` spécifie le type de tâche à utiliser (`waitgroup` pour la tâche WaitGroup, `channel` pour la tâche Channel).

Le programme affichera le temps d'exécution individuel et total du traitement d'images.

N'hésitez pas à explorer le code source pour en savoir plus sur l'implémentation détaillée.

## Auteur

Ce projet a été réalisé par Ben-Jamin MAMFOUMBI KOUELY.
