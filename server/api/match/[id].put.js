import { serverSupabaseClient } from '#supabase/server'

const algCharToNumber= (char) => char.charCodeAt(0) - 97;

function boardToFen(board) {
    const fenRows = [];

    for (const row of board) {
        let fenRow = '';

        for (const cell of row) {
            if (cell === '') {
                fenRow += '1';
            } else {
                fenRow += cell;
            }
        }

        fenRows.push(fenRow);
    }

    return fenRows.join('/');
}

function calculateFenMove(fen , move) {
    const piece = move[0];
    const fromAlg = move.substring(1, 3);
    const toAlg = move.substring(3, 5);
    
    const [ fromX, fromY ] = [algCharToNumber(fromAlg[0]), fromAlg[1]];
    const [ toX, toY ] = [algCharToNumber(toAlg[0]), toAlg[1]];
    
    const fenRows = fen.split('/');
    const board = [];

    for (const rowFen of fenRows) {
        const row = [];

        for (const char of rowFen) {
            if (isNaN(char)) {
                row.push(char);
            } else {
                for (let i = 0; i < char; i++) {
                    row.push('');
                }
            }
        }

        board.push(row);
    }
    
    board[fromY-1][fromX] = '';
    board[toY-1][toX] = piece;

    return boardToFen(board);
}

async function makeMove(event) {
    const supabase = await serverSupabaseClient(event)
    const matchId = getRouterParam(event, 'id');
    const {move} = await readBody(event);

    if (!matchId) {
        console.log('No match id');
        return;
    }

    const { data } = await supabase
        .from('match')
        .select('*')
        .eq('id', matchId)
        .limit(1)
        .single();

    const board = calculateFenMove(data.board, move);
    const moves = data.moves;

    await supabase
        .from('match')
        .update({ moves: [...moves, move], board: board })
        .eq('id', matchId);
}


export default defineEventHandler(async (event) => {
    const contentType = getRequestHeader(event, 'Content-Type');
    if (!contentType) {
        return;
    }

    const domainModels = contentType
        .split(';')
        .filter(x => x.startsWith('domain-model'));

    if (domainModels.length === 0 || domainModels.length > 1) {
        return;
    }

    const domainModel = domainModels[0].toLowerCase()

    switch (domainModel) {
        case 'domain-model=makemove':
            return await makeMove(event);
        default:
            console.log('Unknown domain model', domainModel);
            return;
    }
});