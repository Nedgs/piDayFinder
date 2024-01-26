# piBirthFinder

## Présentation

Le projet piBirthFinder a pour objectif de trouver la position d'une date de naissance donnée (au format JJMMAAAA) parmi les nombreuses décimales de π (pi). La génération directe de π en TypeScript s'est révélée complexe vu la taille des décimales qu'on veut générer (1 million). Nous avons donc adopté une approche alternative : utiliser un fichier "pi.txt" préexistant, contenant déjà un million de décimales de π.

## Développement

### Étape 1 : Choix du Projet

Nous avons opté pour un peu projet original visant à rechercher une date de naissance dans les décimales de π. La génération directe de π s'est avérée difficile, d'où l'utilisation d'un fichier "pi.txt" préparé au préalable avec un million de chiffres.

### Étape 2 : Recherche de Pi

Mise en place d'un fichier .txt avec le premier million de décimales. Possibilité de récupérer la valeur de Pi 
uniquement en Python (time out en python, + inf en python). Cependant le but était de se baser sur la même valeur.

### Etape 3 : Mise en oeuvre dans les différentes technologies 

## Typescript
La mise en œuvre en TypeScript comprend les étapes suivantes :
- Lecture du fichier "pi.txt".
- Recherche de la date de naissance dans π.
- Mesure du temps d'exécution avec perf_hooks et benchmarks.js (lent à démarrer).
- Identification des problèmes de performance avec clinic doctor (qui est trés compliqué à mettre en place).

## GO
La mise en œuvre en TypeScript comprend les étapes suivantes :
- Récupération du fichier .txt
- Recherche de la date de naissance dans pi
- Mise en place de goroutine
- Mesure du temps d'execution time / pprof
- Lancement de benchmarks

Implémentation envisager mais pas mis en place :
- Memoization : lancement de la méthode une unique fois avant fermeture du programme donc pas d'utilité
- Utilisation du design pattern semaphore. Fichier trop petit, gain de temps insignifiant + gestion des intersections compliquées pour pas grand chose