# piDayFinder

## Présentation

Le projet piDayFinder a pour objectif de trouver la position d'une date de naissance donnée (au format JJMMAAAA) parmi les nombreuses décimales de π (pi). La génération directe de π en TypeScript s'est révélée complexe. Nous avons donc adopté une approche alternative : utiliser un fichier "pi.txt" préexistant, contenant déjà un million de décimales de π.

## Développement

### Étape 1 : Choix du Projet

Nous avons opté pour un peu projet original visant à rechercher une date de naissance dans les décimales de π. La génération directe de π s'est avérée difficile, d'où l'utilisation d'un fichier "pi.txt" préparé au préalable avec un million de chiffres.

### Étape 2 : Mise en Œuvre en TypeScript

La mise en œuvre en TypeScript comprend les étapes suivantes :
- Lecture du fichier "pi.txt".
- Recherche de la date de naissance dans π.
- Mesure du temps d'exécution avec perf_hooks et benchmarks.js (lent à démarrer).
- Identification des problèmes de performance avec clinic doctor (qui est trés compliqué à mettre en place).
