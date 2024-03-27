# CLI Todo List App

Cette application est une simple liste de tâches en ligne de commande (CLI) écrite en Go.
Avec l'aide du livre : powerful command line applications in go
Elle vous permet de gérer vos tâches quotidiennes facilement depuis votre terminal.

## Installation

Assurez-vous d'avoir Go installé sur votre système. Clonez ensuite ce dépôt :

`git clone https://github.com/Alexis3386/todocli`

Naviguez dans le répertoire nouvellement créé :

`cd todocli`

Comppilez l'application

`go build`

## Utilisation

Pour ajouter une tâche à votre liste, utilisez la commande suivante :

`./todocli -task "Nom de la tâche"`

Pour afficher la liste de toutes vos tâches, utilisez la commande suivante :

`./todocli -list`

Pour afficher l'aide :

`./todocli -h`
