let n = 1260

function change(n: number) {
  let cnt = 0
  const cointTypes = [500, 100, 50, 10]
  for (let index = 0; index < cointTypes.length; index++) {
    cnt += Math.floor(n / cointTypes[index])
    n %= cointTypes[index]
  }
  return cnt
}

console.log(change(n))
