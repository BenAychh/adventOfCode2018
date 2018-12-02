const {
  processor,
  repeatFinder,
} = require('./day1');

describe('day 1, part 1', () => {
  test('+1, +1, +1 results in 4', () => {
    const data = ['+1', '+1', '+1'];
    expect(processor(data)).toEqual(3);
  })

  test('+1, +1, -2 results in 0', () => {
    const data = ['+1', '+1', '-2'];
    expect(processor(data)).toEqual(0);
  })

  test('-1, -2, -3 results in -6', () => {
    const data = ['-1', '-2', '-3'];
    expect(processor(data)).toEqual(-6);
  })
})

describe('day 1, part 2', () => {
  test('+1, -1 first reaches 0 twice.', () => {
    const data = ['+1', '-1'];
    expect(repeatFinder(data)).toEqual(0);
  })

  test('+3, +3, +4, -2, -4 first reaches 10 twice', () => {
    const data = ['+3', '+3', '+4', '-2', '-4'];
    expect(repeatFinder(data)).toEqual(10);
  })

  test('-6, +3, +8, +5, -6 first reaches 5 twice', () => {
    const data = ['-6', '+3', '+8', '+5', '-6'];
    expect(repeatFinder(data)).toEqual(5);
  })

  test('+7, +7, -2, -7, -4 first reaches 14 twice', () => {
    const data = ['+7', '+7', '-2', '-7', '-4'];
    expect(repeatFinder(data)).toEqual(14);
  })
})