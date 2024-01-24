import * as fs from 'fs';
import * as readlineSync from 'readline-sync';

function findPosition(birthDate: string, filePath: string): number | undefined {
    const fileContent = fs.readFileSync(filePath, 'utf-8');
    const position = fileContent.indexOf(birthDate);

    if (position !== -1) {
        // Add 1 because positions start from 0
        return position + 1;
    } else {
        return undefined;
    }
}

// Get user input for birth date
const birthDateToSearch = readlineSync.question('Enter your birth date (MMDD): ');

const foundPosition = findPosition(birthDateToSearch, './../../pi.txt');

if (foundPosition !== undefined) {
    console.log(`The position of ${birthDateToSearch} in the file is: ${foundPosition}`);
} else {
    console.log(`The number was not found in the file.`);
}
