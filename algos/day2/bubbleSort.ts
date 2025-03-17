import { algoTester } from "./algorithmTester";

function bubbleSort(arr: number[]) {
    const arrLen = arr.length;
    for (let i = 0; i < arrLen; i++) {
        for (let j = 0; j < arrLen - i; j++) {
            if (arr[j] > arr[j + 1]) {
                let temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
            }
        }
    }
};

bubbleSort([3, 2, 1, 7]);

algoTester({
    arrSize: 1000,
    funct: bubbleSort,
});
