import {Component} from 'solid-js';
import {ChessBoard} from "./components/ChessBoard";


const App: Component = () => {
    return <div class="flex flex-col place-content-center w-screen h-screen gap-8 dark:text-white dark:bg-stone-900">
        <h1 class="mx-auto font-black tracking-tight text-4xl">Chaoss</h1>
        <div class="mx-auto w-fit h-fit">
            <ChessBoard/>
        </div>

        <button class="mx-auto w-24 bg-blue-500 hover:bg-blue-400 text-white font-bold py-2 px-4 border-b-4 border-blue-700 hover:border-blue-500 rounded">
            Go Next
        </button>
    </div>
};

export default App;
