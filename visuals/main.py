#!/usr/bin/env python
import sys
from math import cos, pi, sin

import yaml
from PIL import Image, ImageDraw

USAGE = f"Usage: {sys.argv[0]} data.yaml output.png"
SIZE = (800, 800)
CENTER = (400, 400)
RADIUS = 300


def draw_generation(draw: ImageDraw.ImageDraw, gen: list[list[int]], n: int):
    alpha = [[0 for _ in range(n)] for _ in range(n)]
    for path in gen:
        for i, j in zip(path, path[1:] + [path[0]]):
            alpha[i][j] += 1

    for i in range(n):
        for j in range(n):
            val = int(255 * (1 - alpha[i][j] / len(gen)))
            draw_edge(draw, i, j, n, (val, val, val))


def draw_path(draw: ImageDraw.ImageDraw, points: list[int], n: int):
    for p1, p2 in zip(points, points[1:]):
        draw_edge(draw, p1, p2, n)


def draw_edge(draw: ImageDraw.ImageDraw, start: int, finish: int, n: int, color):
    point1 = pi / 2 - 2 * pi * start / n
    point2 = pi / 2 - 2 * pi * finish / n
    d1 = (cos(point1), -sin(point1))
    d2 = (cos(point2), -sin(point2))
    p1 = (CENTER[0] + RADIUS * d1[0], CENTER[1] + RADIUS * d1[1])
    p2 = (CENTER[0] + RADIUS * d2[0], CENTER[1] + RADIUS * d2[1])
    draw.line([p1, p2], fill=color, width=1)


def main():
    if len(sys.argv) < 3:
        print(USAGE)
        sys.exit()
    fi = sys.argv[1]
    with open(fi, "r") as f:
        data = yaml.safe_load(f)
    n = data["city_count"]
    with Image.new("RGBA", SIZE, "white") as img:
        draw = ImageDraw.Draw(img)
        gen = data["simulation"][0]
        draw_generation(draw, gen, n)
        img.save(sys.argv[2], "PNG")


if __name__ == "__main__":
    main()
