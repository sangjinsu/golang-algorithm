function adventurerGuild(adventurers: number[]) {
  adventurers.sort((a, b) => a - b)

  let result = 0
  let count = 0

  for (let i = 0; i < adventurers.length; i++) {
    count++
    if (count >= adventurers[i]) {
      result++
      count = 0
    }
  }

  return result
}

console.log(adventurerGuild([2, 1, 2, 3, 2]))
