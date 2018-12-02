const {
  loadData
} = require('../helpers/loadData');
const {
  checksum,
} = require('./day2.part1')

const {
  commonLetters
} = require('./day2.part2')

async function main() {
  const data = await loadData('day2.data');
  let start = Date.now()
  const part1 = checksum(data)
  let end = Date.now()
  console.log('part1', part1, (end - start))

  start = Date.now()
  const part2 = commonLetters(data)
  end = Date.now()
  console.log('part2', part2, (end - start))
}

main()