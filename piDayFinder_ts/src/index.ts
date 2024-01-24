import { Big } from 'big.js';
import * as readlineSync from 'readline-sync';

// Fonction pour calculer les deux premiers millions de chiffres de pi
function calculateTwoMillionDigitsOfPi() {
    let pi = new Big(0);
    let k = 0;
    let precision = 1000000;
    let kMax = precision;
    let numerator = new Big(1);
    let denominator = new Big(1);
    let temp = new Big(0);
    let sign = 1;

    while (k < kMax) {
        temp = numerator.div(denominator);
        temp = temp.times(sign);
        pi = pi.plus(temp);
        numerator = numerator.plus(2);
        denominator = denominator.plus(2);
        sign = sign * -1;
        k++;
    }

    pi = pi.times(4);
    return pi;
}

//afficher la valeur de pi
const pi = calculateTwoMillionDigitsOfPi();
console.log(pi);

// // Fonction pour trouver la position d'une date dans les chiffres de pi
// function findPositionOfDate(day: string, month: string, piString: string): number {
//     const dateDigits = `${month}${day}`;
//     const piSubstring = piString.indexOf(dateDigits);

//     if (piSubstring !== -1) {
//         return piSubstring + 1; // La position commence à 1, pas à 0
//     } else {
//         return -1; // La date n'a pas été trouvée
//     }
// }

// Calculer les deux premiers millions de chiffres de pi
// const pi = calculateTwoMillionDigitsOfPi();
// console.log(pi);
// const piString = pi.toString();

// // Saisir la date de naissance depuis la console
// const day = readlineSync.question('Entrez le jour de votre date de naissance (ex. 01): ');
// const month = readlineSync.question('Entrez le mois de votre date de naissance (ex. 01): ');

// // Trouver la position de la date dans les chiffres de pi
// const position = findPositionOfDate(day, month, piString);

// // Afficher les résultats
// if (position !== -1) {
//     console.log(`Votre date de naissance (${day}-${month}) a été trouvée à la position ${position} dans les deux premiers millions de chiffres de pi.`);
// } else {
//     console.log(`Votre date de naissance (${day}-${month}) n'a pas été trouvée dans les deux premiers millions de chiffres de pi.`);
// }


