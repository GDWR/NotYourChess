<template>
    <div>
      <div v-for="(row, y) in board" class="flex"> 
          <div v-for="(piece, x) in row" class="w-12 h-12 bg-emerald-400 select-none hover:cursor-pointer" :class="{'bg-emerald-900': (x+y)%2}" @click="onSquareClicked(piece, x, y)">
            <div v-if="x==selectedPiece?.x && y==selectedPiece?.y" class="absolute opacity-50 bg-yellow-500 w-12 h-12" />
            <img v-if="piece" :src="`/chess/${piece}.svg`" class="relative select-none pointer-events-none" />
          </div>  
      </div>
    </div>
</template>


<script setup lang="ts">
const emit = defineEmits<{
  (e: 'move', move: string): void
}>();

const props = defineProps<{
    board: string,
}>();

const numToLetter = (num: number) => String.fromCharCode(97 + num)
const selectedPiece = ref<{piece: string, x: number, y: number} | null>(null);
const canMove = ref(true);
const board = boardFromFen(props.board);

function boardFromFen(fen: string) {
  const rows = fen.split(' ')[0].split('/');
  return rows.map(row => {
    var newRow = [];
    for (var i = 0; i < row.length; i++) {
      var char = row[i];
      if (isNaN(parseInt(char))) {
        newRow.push(char);
      } else {
        for (var j = 0; j < parseInt(char); j++) {
          newRow.push('');
        }
      }
    }
    return newRow;
  }).reverse();
}


function onSquareClicked(piece: string, x: number, y: number) {
  if (canMove.value === false) return;

  if (selectedPiece.value) {
    const {piece, x: selectedX, y: selectedY} = selectedPiece.value;

    var startingPosition = `${numToLetter(selectedX)}${Math.abs(selectedY-8)}`;
    var endingPosition = `${numToLetter(x)}${Math.abs(y-8)}`;
    
    console.log({
      move: `${piece}${startingPosition}${endingPosition}`,
    });
    emit('move', `${piece}${startingPosition}${endingPosition}`);
    canMove.value = false;

    board[y][x] = piece;
    board[selectedY][selectedX] = '';
    selectedPiece.value = null;
  } else {
    if (!piece) return;
    selectedPiece.value = {piece, x, y};
  }
}

function reset() {
  canMove.value = true;
  selectedPiece.value = null;
}
</script>
