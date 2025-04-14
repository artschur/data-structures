const fixedLenghtEmptyArr: Array<number> = new Array(2);

const dynamicLenghtFilledArr: Array<number> = [1, 2, 3, 4, 5];

console.log(fixedLenghtEmptyArr);

fixedLenghtEmptyArr.push(6);
fixedLenghtEmptyArr.push(6);
fixedLenghtEmptyArr.push(6);

const array2: number[] = new Array(2);
array2[0] = 1;
array2[1] = 2;

console.log(fixedLenghtEmptyArr);
console.log(array2);
dynamicLenghtFilledArr.push(6);
console.log(dynamicLenghtFilledArr);
