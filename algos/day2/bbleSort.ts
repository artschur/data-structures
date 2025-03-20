function remakeBubbleSort(arr: number[]) {
  const arrLen = arr.length;
  for (let i = 0; i < arrLen; i++) {
    for (let j = 0; j < arrLen - i; j++) {
      if (arr[j] > arr[j + 1]) {
        let temp = arr[j + 1];
        arr[j + 1] = arr[j];
        arr[j] = temp;
      }
    }
  }
  console.log(arr);
}

remakeBubbleSort([1, 12, 4, 3, 8, 8]);
