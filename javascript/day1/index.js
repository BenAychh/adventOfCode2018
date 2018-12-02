const data = require('./day1.json');
const {
  processor,
  repeatFinder,
} = require('./day1');

// console.log(processor(data))

const start = Date.now()
const results = repeatFinder(data)
const end = Date.now();

console.log(repeatFinder(data), end - start);