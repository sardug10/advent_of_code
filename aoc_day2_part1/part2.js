import fs from 'fs';

const actualCubesCount = {
  red: 12,
  green: 13,
  blue: 14,
};

const games = fs.readFileSync('input.txt', 'utf8').trim().split('\n');

let result = 0;
games.forEach((game) => {
  let something = game.split(':');
  const gamevalues = something[1];

  const maxColorCube = {};
  gamevalues.split(';').map((el) => {
    el.split(',').map((innerEl) => {
      const actualgame = innerEl.trim().split(' ');
      const cubeColor = actualgame[1];
      const cubeNumber = +actualgame[0];

      if (cubeColor in maxColorCube) {
        if (maxColorCube[cubeColor] < cubeNumber) {
          maxColorCube[cubeColor] = cubeNumber;
        }
      } else {
        maxColorCube[cubeColor] = cubeNumber;
      }
    });
  });

  const product = Object.values(maxColorCube).reduce(
    (prod, el) => prod * el,
    1
  );

  result += product;
});

console.log(result);
