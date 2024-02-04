<template>
  <div class="w-screen h-screen flex flex-col place-content-center">
    <h1 class="mx-auto font-black tracking-tight text-4xl">Chaoss</h1>

    <div v-if="pending" class="mx-auto bg-red-500 h-96 w-96"/>
    <ChessBoard ref="chessBoard" class="mx-auto p-8" :board="match.board" @move="move => selectedMove = move"/>

    <div class="mx-auto flex">
      <div class="p-2 mx-4 ">
        Your move <UTooltip text="This is your move in algebraic form" class="text-slate-400">{{ selectedMove }}</UTooltip>
      </div>

      <UButtonGroup size="md" orientation="horizontal">
        <UButton icon="i-heroicons-arrow-path"  label="Undo" color="gray" leading  :disabled="selectedMove === null" @click="chessBoard.reset()"/>
        <UButton icon="i-heroicons-arrow-right" label="Go Next" trailing :disabled="selectedMove === null" @click="submitMove()"/>
      </UButtonGroup>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const selectedMove = ref<string | null>(null);
const { data: match, pending } = useLazyFetch('/api/match');
const chessBoard = ref<null>(null);

async function submitMove() {
  if (selectedMove.value === null) return;

  await fetch(`/api/match/${match.value!.id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json;domain-model=MakeMove;' },
    body: JSON.stringify({ 
      move: selectedMove.value,
      token: match.value!.token, 
    }),
  });

  document.location.reload();
}
</script>
