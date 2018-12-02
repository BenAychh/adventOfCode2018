function processor(data) {
  return data.map(removePlusSignsAndConvertToNumber).reduce(sum);
}

function removePlusSignsAndConvertToNumber(entry) {
  return +entry.replace('+', '');
}

function sum(currentValue, entry) {
  return currentValue + entry;
}

function repeatFinder(data) {
  const freqMap = {
    0: true,
  }

  let repeated = null;

  let sum = 0;

  const dataAsNumbers = data.map(removePlusSignsAndConvertToNumber);
  let count = 0;
  while (repeated === null) {
    for (let i = 0; i < dataAsNumbers.length; i++) {
      const number = dataAsNumbers[i];
      sum = sum + number;
      if (freqMap[sum]) {
        repeated = sum;
        break;
      } else {
        freqMap[sum] = true;
      }
    }
  }
  return repeated;
}

module.exports = {
  processor,
  repeatFinder,
}