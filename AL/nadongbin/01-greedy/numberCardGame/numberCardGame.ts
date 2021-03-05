function numberCardGame(cards: number[][]) {
  let max = 0
  for (let i = 0; i < cards.length; i++) {
    const min = Math.min(...cards[i])
    max = Math.max(max, min)
  }
  return max
}

const cards = [
  [3, 1, 2],
  [4, 1, 4],
  [2, 2, 2],
]

console.log(numberCardGame(cards))
