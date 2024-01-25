import * as fs from 'fs';
import * as readlineSync from 'readline-sync';
import { performance } from 'perf_hooks';
// @ts-ignore
import * as Benchmark from 'benchmark';

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

const birthDateToSearch = readlineSync.question('Enter your birth date (MMDD): ');

const suite = new Benchmark.Suite();

// Add a test to the suite
suite.add('findPosition', () => {
    findPosition(birthDateToSearch, './../../pi.txt');
});

suite.on('cycle', (event: any) => {
    console.log(String(event.target));
});

suite.on('complete', function () {
    // @ts-ignore
    console.log('Fastest is ' + this.filter('fastest').map('name'));
});

const startTime = performance.now();

suite.run({ async: false });

const endTime = performance.now();

const executionTime = endTime - startTime;
console.log(`Total execution time: ${executionTime} milliseconds`);
