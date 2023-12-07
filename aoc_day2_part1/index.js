import fs from 'fs';

const actualCubesCount = {
  red: 12,
  green: 13,
  blue: 14,
};

const games = fs.readFileSync('input.txt', 'utf8').trim().split('\n');
console.log(games);

let result = 0;
games.forEach((game) => {
  let something = game.split(':');
  const gameLabel = something[0];
  const gamevalues = something[1];

  const gameNum = +gameLabel.split(' ')[1];
  let isValidgame = true;
  gamevalues.split(';').map((el) => {
    el.split(',').map((innerEl) => {
      const actualgame = innerEl.trim().split(' ');
      const cubeColor = actualgame[1];
      const cubeNumber = +actualgame[0];

      if (actualCubesCount[cubeColor] < cubeNumber) {
        isValidgame = false;
      }
    });
  });

  if (isValidgame) {
    result += gameNum;
  }
});

console.log(result);
