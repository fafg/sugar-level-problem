const results = require('./endpoints/results.js');
const amount = process.env.DATASET || 300;

module.exports = () => {
  return {
    results: results(amount),
  };
};