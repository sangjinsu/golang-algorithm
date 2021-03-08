function showDestination(n: number, moves: string) {
  const movesArr = moves.split(' ')

  // L R U D
  const dx = [0, 0, -1, 1]
  const dy = [-1, 1, 0, 0]

  let [x, y] = [1, 1]

  for (let i = 0; i < movesArr.length; i++) {
    let [nx, ny] = [x, y]
    switch (movesArr[i]) {
      case 'L':
        nx += dx[0]
        ny += dy[0]
        break
      case 'R':
        nx += dx[1]
        ny += dy[1]
        break
      case 'U':
        nx += dx[2]
        ny += dy[2]
        break
      case 'D':
        nx += dx[3]
        ny += dy[3]
        break
    }

    if (nx < 1 || nx > n || ny < 1 || ny > n) continue
    ;[x, y] = [nx, ny]
  }

  return [x, y].join(' ')
}

console.log(showDestination(5, 'R R R U D D'))
