// Just some Jest stuff

const { afterUpdate } = require('svelte')

beforeEach(() => {
  // I run before each test
})

it('it works', async () => {
  let x = 10
  expect(x).toBe(10)
  expect(x).toEqual(10)
})

it('matches', async () => {
  let y = 'asdf'
  expect(y).toMatch(/a.*f/g)
})

it('arrays', async () => {
  let arr = ['asdf', 'qwer', 'zxcv']
  expect(arr).toContain('zxcv')
})

function throwsError() {
  throw new Error("I'm an error")
}

it('errors', async () => {
  expect(throwsError).toThrow()
  expect(throwsError).toThrow(Error)
  expect(throwsError).toThrow(/.*error/g)
})

function promiseReturnTest(cb) {
  return new Promise((res, rej) => {
    setTimeout(() => {
      if (cb) {
        cb('test')
      }
      res('test')
    }, 1000)
  })
}

it('callsback', async (done) => {
  function cb(data) {
    try {
      expect(data).toBe('test')
      done()
    } catch (err) {
      done(err)
    }
  }

  promiseReturnTest(cb)
})

it('promises', async () => {
  promiseReturnTest().then((resp) => expect(resp).toBeTruthy())
})

it('resolves', async () => {
  return expect(promiseReturnTest()).resolves.toBe('test')
})

it('async/awaits', async () => {
  expect.assertions(1)
  try {
    const val = await promiseReturnTest()
    expect(val).toBe('test')
  } catch (err) {
    expect(err).toMatch('error')
  }
})

afterEach(() => {
  // I run after each test
})
