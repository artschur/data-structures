function generateRandomArray(size: number): number[] {
    const array: number[] = [];
    for (let i = 0; i < size; i++) {
        array.push(Math.floor(Math.random() * 100000));
    }
    return array;
}

export function algoTester({ arrSize = 1000, funct, silent = true }: { arrSize: number; funct: (a: number[]) => void; silent?: boolean; }) {

    const testArray = generateRandomArray(arrSize ?? 1000);
    const startTime = Date.now();
    funct(testArray);
    const endTime = Date.now();
    console.log(`Sorting completed in ${endTime - startTime} ms`);

};