function untilValueReachesOne(N: number, K: number) {
  let count = 0

  while (true) {
    let target = Math.floor(N / K) * K
    count += N - target
    N = target

    if (N < K) {
      break
    }

    count++
    N = Math.floor(N / K)
  }

  count += N - 1
  return count
}

console.log(untilValueReachesOne(15, 5))
