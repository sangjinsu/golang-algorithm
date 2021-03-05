function theLawofLargeNumbers(n: number, m: number, k: number, nums: number[]) {
  nums.sort((a, b) => b - a)

  const count = Math.floor(m / (k + 1)) * k + (m % (k + 1))
  const sum = nums[0] * count + nums[1] * (m - count)

  return sum
}

const nums = [7, 6, 5, 1, 2]
const [n, m, k] = [5, 8, 4]

console.log(theLawofLargeNumbers(n, m, k, nums))
