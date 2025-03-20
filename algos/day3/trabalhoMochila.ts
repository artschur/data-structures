interface BagItem {
  id: number;
  utilidade: number;
  peso: number;
  quantidadeDisponivel: number;
  quantidadeImplementada: number;
}

interface Bag {
  itens: BagItem[];
  capacidade: number;
}

function readItemsFromInput(input: string): Bag {
  const lines = input.trim().split("\n");

  if (!lines) {
    console.log("not worked");
    return { itens: [], capacidade: 0 };
  }
  const itens: BagItem[] = [];

  console.log(lines);
  for (let line = 0; line < lines.length; line++) {
    const [utilidade, peso, quantidadeDisponivel] = lines[line]
      .split(" ")
      .map(Number); //converte para num

    if (utilidade < 0) {
      const capacity = parseInt(lines[lines.length - 1].trim());
      return {
        itens: itens,
        capacidade: capacity,
      };
    }

    itens.push({
      id: line,
      utilidade: utilidade,
      peso: peso,
      quantidadeDisponivel: quantidadeDisponivel,
      quantidadeImplementada: 0,
    });
  }
  return { itens: [], capacidade: 0 };
}

function selectionSort(arr: BagItem[]): BagItem[] {
  let arrLen = arr.length;
  for (let i = 0; i < arrLen - 1; i++) {
    let smallestIndex = i;
    for (let j = i + 1; j < arrLen; j++) {
      if (
        arr[j].utilidade / arr[j].peso <
        arr[smallestIndex].utilidade / arr[smallestIndex].peso
      ) {
        smallestIndex = j;
      }
    }
    let temp = arr[smallestIndex];
    arr[smallestIndex] = arr[i];
    arr[i] = temp;
  }
  return arr;
}

async function main() {
  console.log("Enter bag items (format: utilidade peso quantidade)");
  console.log("Enter -1 -1 -1 to mark the end of items");
  console.log("Then enter the bag capacity:");

  const file = Bun.file("./inputTest.txt");

  const input = await file.text();

  const bag: Bag = readItemsFromInput(input);

  const itensOrdenadosPorMedia = selectionSort(bag.itens);

  let pesoDosItensAdicionados = 0;
  for (let item of itensOrdenadosPorMedia.reverse()) {
    for (let i = 1; i <= item.quantidadeDisponivel; i++) {
      if (i * item.peso + pesoDosItensAdicionados <= bag.capacidade) {
        pesoDosItensAdicionados += i * item.peso;
        item.quantidadeImplementada++;
      }
    }
    bag.itens.push(item);
  }
  console.log(bag.itens);
}

main();
