const {
  commonLetters
} = require('./day2.part2');

describe('Day 2 Part 2', () => {
  it('calculates the correct common letters for "abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"', () => {
    const data = ["abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"]
    const expected = "fgij"
    const actual = commonLetters(data)

    expect(actual).toEqual(expected)
  })
})