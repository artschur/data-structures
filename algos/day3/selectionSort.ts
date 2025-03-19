import { algoTester } from "../day2/algorithmTester";

function selectionSort(arr: number[]) {
    const arrLen = arr.length;
    for (let i = 0; i + 1 < arrLen; i++) {
        let smallestIndex = i;
        for (let j = i + 1; j < arrLen; j++) {
            if (arr[j] < arr[smallestIndex]) {
                smallestIndex = j;
            }
        }
        let temp = arr[smallestIndex];
        arr[smallestIndex] = arr[i];
        arr[i] = temp;

    }
    return arr;
};

algoTester({
    arrSize: 20000,
    funct: selectionSort
});