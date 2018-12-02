const {
  checksum
} = require('./day2.part1');

describe('Day 2', () => {
  it('part 1', () => {
    const data = ["abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"]
    const expected = 12
    const actual = checksum(data)

    expect(actual).toEqual(expected)
  })
})