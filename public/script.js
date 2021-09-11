let points = [];
const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");

// Draw canvas and points
const redraw = () => {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;
    ctx.fillStyle = "lightgray";
    ctx.fillRect(0, 0, canvas.width, canvas.height);
    points.forEach((p) => {
        ctx.fillStyle = "black";
        ctx.fillText(`${p.x}; ${p.y}`, p.x + 5, p.y - 5);
        ctx.beginPath();
        ctx.arc(p.x, p.y, 5, 0, Math.PI * 2);
        ctx.fill();
    });
}
redraw();

const convex = () => {
    if(points.length >= 3) {
        fetch("/convex2d", {
            method: "POST",
            body: JSON.stringify({ points }),
            headers: {
                "Content-Type": "application/json"
            }
        }).then(res => res.json()).then(res => {
            ctx.beginPath();
            res.forEach((p) => {
                ctx.lineTo(p.x, p.y);
            });
            ctx.lineTo(res[0].x, res[0].y);
            ctx.stroke();
        });
    }
}

// Update canvas to fit window size
window.addEventListener("resize", function () {
    redraw();
    convex();
});

// Draw point on click on canvas and get convex hull
canvas.addEventListener("click", function (e) {
    const x = e.pageX - this.offsetLeft;
    const y = e.pageY - this.offsetTop;
    points.push({ x, y });
    redraw();
    convex();
});