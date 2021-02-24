let x=0.01,y=0,z=0;
let a=10,b=28,c=8.0/3.0;
let points = new Array;
let angle=1;
function setup(){
    createCanvas(800,600,WEBGL);
    colorMode(HSB);
}

function draw(){
    background(0);
    let dt=0.01;
    let dx=a*(y-x),dy=x*(b-z)-y,dz=x*y-c*z;dx*=dt;dy*=dt;dz*=dt;
    x+=dx;y+=dy;z+=dz;
    let h=0;
    //let camX = map(mouseX, 0, width, -200, 200);
    //let camY = map(mouseY, 0, height, -200, 200);
    //camera(camX, camY,(height/2.0) / tan(PI*30.0 / 180.0), 0, 0, 0, 0, 1, 0);
    points.push(new p5.Vector(x, y, z));
    scale(5);
    noFill();
    strokeWeight(2);
    //translate(width/2,height/2);
    beginShape();
    points.forEach(function(v){
        stroke(h,255,255);
        vertex(v.x,v.y,v.z);
        h+=0.1;
        if(h>255){
            h=0;
        }
    })
    endShape();
}