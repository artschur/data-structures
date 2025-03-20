interface BagItem {
    id: number;
    utilidade: number;
    peso: number;
    quantidade: number;
}

interface Bag {
    itens: BagItem[];
    capacidade: number;
}

function readItemsFromInput(input: string): Bag {
    const lines = input.trim().split('\n');
    const itens: BagItem[] = [];

    for (let line = 0; line < lines.length; line++) {
        const [utilidade, peso, quantidade] = lines[line]
            .split(' ')
            .map(Number); //converte para num

        if (utilidade < 0) {
            const capacity = parseInt(lines[lines.length - 1].trim());
            return {
                itens: itens,
                capacidade: capacity
            };
        };

        itens.push({
            id: line,
            utilidade: utilidade,
            peso: utilidade,
            quantidade: quantidade
        });
    }
    return { itens: [], capacidade: 0 };
}

function selectionSort(arr: BagItem[]): BagItem[] {
    let arrLen = arr.length;
    for (let i = 0; i < arrLen - 1; i++) {
        let smallestIndex = i;
        for (let j = i + 1; j < arrLen; j++) {
            if (arr[j].utilidade / arr[j].peso < arr[smallestIndex].utilidade / arr[smallestIndex].peso) {
                smallestIndex = j;
            }
        }
        let temp = arr[smallestIndex];
        arr[smallestIndex] = arr[i];
        arr[i] = temp;
    }
    return arr;
};

function main() {
    const input = prompt("Insira (utility weight quantity) por linha, -1 -1 -1 para terminar, e capacidade no final):");

    if (!input) {
        console.log("No input received.");
        return;
    }

    const bag: Bag = readItemsFromInput(input);

    const itensOrdenadosPorMedia = selectionSort(bag.itens);

    // processar final (quantidade)
    // quantidade x peso  -> capacidade da mochila
    let pesoDosItensAdicionados = 0;
    for (let item of bag.itens) {
        for (let i = 1; i <= item.quantidade; i++) {
            if (i * item.peso) {

            }
        }
    }
    console.log(bag);
}

