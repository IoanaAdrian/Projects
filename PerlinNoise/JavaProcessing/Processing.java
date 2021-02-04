import processing.core.PApplet;

public class Processing extends PApplet {

    public static void main(String[] args) {
        PApplet.main("Processing");
    }

    int rows, cols, scl = 20;
    int w = 2000;
    int h = 1600;

    float flying = 0;
    float[][] terrain;

    public void settings() {
        size(600, 600, P3D);
    }

    public void setup() {
        cols = h / scl;
        rows = w / scl;
        terrain = new float[rows][cols];

    }

    public void draw() {
        flying -= 0.1;
        float xoff = 0;
        for (int x = 0; x < rows; x++) {
            float yoff = flying;
            for (int y = 0; y < cols; y++) {
                terrain[x][y] = map(noise(xoff, yoff), 0, 1, -100, 100);
                yoff += 0.2;
            }
            xoff += 0.2;
        }
        background(0);
        stroke(255);
        noFill();

        translate(width / 2, height / 2);
        rotateX(PI / 3);

        translate(-w / 2, -h / 2+50);
        for (int x = 0; x < rows - 1; x++) {
            beginShape(TRIANGLE_STRIP);
            for (int y = 0; y < cols; y++) {
                vertex(x * scl, y * scl, terrain[x][y]);
                vertex((x + 1) * scl, y * scl, terrain[x + 1][y]);
            }
            endShape();
        }

    }

}
