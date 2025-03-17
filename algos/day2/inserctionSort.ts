import { algoTester } from "./algorithmTester";

function inserctionSort(array: number[]) {
    const arrLen = array.length;
    for (let i = 1; i < arrLen; i++) {
        let pivot = array[i];
        let j = i;
        while (j >= 0 && array[j - 1] > pivot) {
            array[j] = array[j - 1];
            j--;
        }
        array[j] = pivot; //tem que ser com o pivot por que senao ele vai subtituir o numero repetido
    }
}
// inserctionSort([3, 2, 1, 7, 4, 12, 11, 1]);

algoTester({
    arrSize: 1000,
    funct: inserctionSort
});