import {Component, createResource, createSignal, Index, Show} from "solid-js";

const fetchRandomMatch = async () => {
    const response = await fetch(`/api/match`);
    const match = await response.json();
    match.board = boardFromFen(match.board);
    return match;
}

function boardFromFen(fen: string) {
    const rows = fen.split(' ')[0].split('/');
    return rows.map(row => {
        const newRow = [];
        for (let i = 0; i < row.length; i++) {
            const char = row[i];
            if (isNaN(parseInt(char))) {
                newRow.push(char);
            } else {
                for (let j = 0; j < parseInt(char); j++) {
                    newRow.push('');
                }
            }
        }
        return newRow;
    });
}

const pieceToImage = (piece: string) => {
    if (piece === '') return 'empty';
    const color = piece === piece.toLowerCase() ? 'b' : 'w';
    const name = piece.toLowerCase();
    return `${color}${name}`;
}

const numToLetter = (num: number) => String.fromCharCode(97 + num)

export const ChessBoard: Component = () => {
    const [selectedPiece, setSelectedPiece] = createSignal<{ piece: string, x: number, y: number }|null>(null);
    const [match] = createResource(fetchRandomMatch);

    async function onPieceClick(data: { piece: string, x: number, y: number }, event: Event) {
        if (selectedPiece() !== null) {
            const {piece, x: selectedX, y: selectedY} = selectedPiece()!;

            const startingPosition = `${numToLetter(selectedX)}${Math.abs(selectedY-8)}`;
            const endingPosition = `${numToLetter(data.x)}${Math.abs(data.y-8)}`;
            const move =  `${piece}${startingPosition}${endingPosition}`;
            console.debug("Move selected", move);

            const board = match()?.board;
            board[data.y][data.x] = piece;
            board[selectedY][selectedX] = '';

            await fetch(`/api/match/${match()?.id}`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json;domain-model=MakeMove;' },
                body: JSON.stringify({
                    move: move,
                }),
            });
            document.location.reload();
        } else {
            if (data.piece === '') return;
            setSelectedPiece(data);
        }
    }

    return <>
        <pre> {JSON.stringify(selectedPiece())}</pre>
        <div class="bg-amber-950 w-fit">
            <Index each={match()?.board}>{(row, x) =>
                <div class="flex">
                    <Index each={row()}>{(piece, y) =>
                        <div class="w-12 h-12 select-none hover:cursor-pointer" classList={{"bg-amber-600": !((x + y) % 2)}} onClick={[onPieceClick, {piece: piece(), x: x, y: y}]}>
                            <Show when={selectedPiece()?.x === x && selectedPiece()?.y === y}>
                                <div class="absolute w-12 h-12 bg-green-500 opacity-50"/>
                            </Show>
                            <Show when={piece()}>
                                <img src={`/chess/${pieceToImage(piece())}.svg`} class="relative select-none pointer-events-none"/>
                            </Show>

                        </div>}
                    </Index>
                </div>}
            </Index>
        </div>
    </>
}
