import { GameState, Player } from "../../server/src/models";

const BG_COLOR = "#231f20";
const FOOD_COLOR = "#e66916";

const initialScreen = document.getElementById("initialScreen") as HTMLDivElement;
const gameScreen = document.getElementById("gameScreen") as HTMLDivElement;
const playerNameInput = document.getElementById("playerNameInput") as HTMLInputElement;
const joinGameBtn = document.getElementById("joinGameButton") as HTMLButtonElement;

let canvas: HTMLCanvasElement, ctx: CanvasRenderingContext2D;
let playerName: string;
let gameActive = false;
let socket: WebSocket;

joinGameBtn.addEventListener("click", () => {
    
    // const code = gameCodeInput.value;
    playerName = playerNameInput.value;

    if (!playerName) {
        alert("Please enter a name");
        return;
    }

    initialScreen.style.display = "none";
    gameScreen.style.display = "block";
    
    socket.send(JSON.stringify( { type: "joinGame", name: playerName } );
    
    console.log(`joinGame ${socket.id} ${playerName}`);
    
    init();

});

// let die: boolean = false;

socket.onmessage = (event) => {
   
    const msg = JSON.parse(event.data);
    
    if (msg.type === "state"){
        
        requestAnimationFrame( () => paintGame(msg.state) );
    
    } else if (msg.type === "die"){
        
        handleDie();
    
    }

});

socket.on("die",()=>{
    
    console.log("die");
    gameActive = false;
    
    if(confirm("You are dead, restart?")){

        socket.send(JSON.stringify( { type: "joinGame", name: playerName } ) );
        gameActive = true;
    }

});

// function reset() {
//     // playerNumber = null;
//     // gameCodeInput.value = "";
//     // gameCodeDisplay.innerText = "";
//     initialScreen.style.display ="block";
//     gameScreen.style.display = "none";
//     playerNameInput.value = "";
//     gameActive = false;
// }

let canvas: HTMLCanvasElement, ctx: CanvasRenderingContext2D;
let playerName;
let gameActive = false;

function init() {

    initialScreen.style.display = "none";
    gameScreen.style.display = "block";
    canvas = document.getElementById("canvas") as HTMLCanvasElement;
    ctx = canvas.getContext("2d");

    const socket = new WebSocket("ws://192.168.1.161:8080/ws"); 
    
    document.addEventListener("keydown", (e) => {
        
        let direction;
        
        switch(e.code){
            case "KeyW":
                direction = "Up";
                break;
            case "KeyA":
                direction = "Left";
                break;
            case "KeyS":
                direction = "Down";
                break;
            case "KeyD":
                direction = "Right";
                break;
            case "ArrowLeft":
            case "ArrowRight":
            case "ArrowDown":
            case "ArrowUp":
                direction = e.code.replace("Arrow", "");
                break;
            default:
                return;
        }
        socket.send(JSON.stringify( { type: "turn", name: playerName,direction } ) ); 
    })
    
    gameActive = true;
}

function paintGame(state: GameState) {
    
    ctx.fillStyle = BG_COLOR;
    canvas.width = canvas.height = 600;
    ctx.fillRect(0, 0, canvas.width, canvas.height);

    const gridsize = state.gridsize;
    const size = canvas.width / gridsize;
    ctx.fillStyle = FOOD_COLOR;
    
    for(let food of state.foods){
        
        ctx.beginPath();
        ctx.arc(
            food.pos.x * size + size / 2, 
            food.pos.y * size + size / 2, 
            size / 2, 
            0, 
            Math.PI * 2
        );
        ctx.fill();
    
    }
    
    ctx.closePath();
    
    for( const [key, player] of Object.entries(state.players) ) {
        paintPlayer(player as any, size);
    }
}

function paintPlayer(playerState: any, size: number) {
    
    ctx.fillStyle = playerState.color;
    
    for(let cell of playerState.snake){

        ctx.fillRect(cell.x * size, cell.y * size, size, size);

    }
    
    ctx.fillStyle = "white";
    ctx.font = "16px verdana";

    ctx.fillText(
        playerState.name,
        playerState.heading.x * size,
        playerState.heading.y * size - 8
    );

}
