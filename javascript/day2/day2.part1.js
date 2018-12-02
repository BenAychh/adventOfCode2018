function checksum(data) {
  const twosAndThrees = data.map(getCounts).map(get2sAnd3s).reduce((accumulator, count) => {
    accumulator[2] += count[2];
    accumulator[3] += count[3];
    return accumulator
  }, {
    2: 0,
    3: 0
  })

  return twosAndThrees[2] * twosAndThrees[3]
}

function getCounts(str) {
  const accumulator = {}
  for (let char of str) {
    if (accumulator[char]) {
      accumulator[char]++
    } else {
      accumulator[char] = 1
    }
  }
  return accumulator
}

function get2sAnd3s(count) {
  const twosAndThress = {
    2: 0,
    3: 0,
  }
  for (let key in count) {
    const value = count[key];
    if (value === 2) {
      twosAndThress[2] = 1;
    }
    if (value === 3) {
      twosAndThress[3] = 1;
    }
    if (twosAndThress[2] === 1 && twosAndThress[3] === 1) {
      break;
    }
  };
  return twosAndThress;
}

module.exports = {
  checksum
}