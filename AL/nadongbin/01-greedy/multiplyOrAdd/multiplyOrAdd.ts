function multiplyOrAdd(numStr: string) {
  const nums = numStr.split('').map((s) => parseInt(s))

  let result = nums[0]
  for (let i = 1; i < nums.length; i++) {
    nums[i] < 2 || result < 1 ? (result += nums[i]) : (result *= nums[i])
  }

  return result
}

console.log(multiplyOrAdd('02984'))
