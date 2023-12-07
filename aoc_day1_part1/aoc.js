import fs from 'fs';

const lines = fs.readFileSync('input.txt', 'utf8').trim().split('\n');

const values = lines.map((line) => {
  const first = line.split('').find((char) => !Number.isNaN(Number(char)));
  const last = line.split('').findLast((char) => !Number.isNaN(Number(char)));

  console.log(line, first + last);
  return Number(first + last);
});

console.log(values.reduce((a, c) => a + c));
