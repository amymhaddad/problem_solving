
export const transform = (oldData) => {
  let newData = {}
  for (const [score, letters] of Object.entries(oldData)) {
    for (let letter of letters) {
      newData[letter.toLowerCase()] = Number(score)
  }
}
  return newData
};

