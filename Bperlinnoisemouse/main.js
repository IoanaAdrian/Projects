

    let rows, cols, scl = 20;
    let w = 2000;
    let h = 1600;

    let flying = 0;
    var terrain=[];
    let R=200;
    function setup() {
        createCanvas(600, 600, WEBGL);
        cols = h / scl;
        rows = w / scl;
        

    }

    function draw() {
        background(0);
        stroke('red');
        noFill();
        //print(mouseX/scl," ",mouseY/scl);
        flying -= 0.1;
        let xoff = 0;
        for (var x = 0; x < rows; x++) {
            terrain[x]=[];
            let yoff = flying;
            for (var y = 0; y < cols; y++) {
                
                let cx=Math.abs(mouseX/scl-x+20)*scl*2,cy=Math.abs(mouseY/scl-y+8)*scl*2;
                if(cx*cx+cy*cy<=R*R){
                    //print("pula"," ",x," ",y," ",colors[x][y]);
                    terrain[x][y]=random(10, 200);
                }else{
                    terrain[x][y] = map(noise(xoff, yoff), 0, 1, -100, 100);
                }
                yoff += 0.2;
            }
            xoff += 0.2;
        }
        translate(width / 2, height / 2);
        //rotateX(PI / 3);

        translate(-w / 2, -h / 2+50);
        for (var x = 0; x < rows - 1; x++) {
            beginShape(TRIANGLE_STRIP);
            for (var y = 0; y < cols; y++) {
                vertex(x * scl, y * scl, terrain[x][y]);
                vertex((x + 1) * scl, y * scl, terrain[x + 1][y]);
            }
            endShape();
        }

    }

