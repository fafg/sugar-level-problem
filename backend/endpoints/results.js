const faker = require('faker');

function generateRandomNumber(min, max) {
  var val = Math.random() * (max - min) + min;
  return Math.round(val * 100) / 100;
}

module.exports = (amount) => {
  let list = [];
  for (let i = 0; i < amount; i++) {
    list.push({
      id: faker.random.uuid(),
      name: faker.name.findName(),
      email: faker.internet.email(),
      samples: [
        {
          time: 'morning',
          value: generateRandomNumber(0.2, 0.9)
        },
        {
          time: 'afternoon',
          value: generateRandomNumber(0.0, 0.5)
        },
        {
          time: 'evening',
          value: generateRandomNumber(0.0, 0.4)
        }
      ]
    });
  }
  return list;
};