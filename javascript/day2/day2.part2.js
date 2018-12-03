function commonLetters(data) {
  const {
    id1,
    index
  } = getSimilarIds(data)
  return id1.substring(0, index) + id1.substring(index + 1)
}

function getSimilarIds(ids) {
  const head = ids[0];
  const tail = ids.slice(1);
  for (let i = 0; i < tail.length; i++) {
    differncesIndexes = getDifferencesLimitTo2(head, tail[i]);
    if (differncesIndexes.length === 1) {
      return {
        id1: head,
        id2: tail[i],
        index: differncesIndexes[0]
      }
    }
  }
  return getSimilarIds(tail)
}

function getDifferencesLimitTo2(str1, str2) {
  const differncesIndexes = []
  for (let i = 0; i < str1.length; i++) {
    const str1Char = str1.charAt(i);
    const str2Char = str2.charAt(i);
    if (str1Char !== str2Char) {
      differncesIndexes.push(i);
    }
    if (differncesIndexes.length > 1) {
      return differncesIndexes;
    }
  }
  return differncesIndexes
}

module.exports = {
  commonLetters
}