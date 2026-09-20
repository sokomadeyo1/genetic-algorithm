#!/usr/bin/env python
import sys
from math import ceil, cos, pi, sin

import yaml
from PIL import Image, ImageDraw, ImageFont
from tqdm import tqdm

USAGE = f"Usage: {sys.argv[0]} data.yaml output.gif"
SIZE = (1250, 1200)
CENTER = (600, 600)
RADIUS = 500
TEXT_BEGIN = (950, 1150)
TEXT_Y_DIFF = -20
FONT = ImageFont.truetype("/usr/share/fonts/noto/NotoSansMono-Regular.ttf", 20)
STROKE = 3
TOTAL_DURATION = 5000
TOTAL_FRAMES = 250


def draw_generation(
    draw: ImageDraw.ImageDraw, gen: list[list[int]], n: int, param: dict[str][str]
):
    alpha = [[0 for _ in range(n)] for _ in range(n)]
    for path in gen:
        for i, j in zip(path, path[1:] + [path[0]]):
            alpha[i][j] += 1

    for i in range(n):
        for j in range(n):
            val = int(255 * (1 - alpha[i][j] / len(gen)))
            draw_edge(draw, i, j, n, (val, val, val))

    pos = TEXT_BEGIN
    for k, v in param.items():
        draw.text(pos, f"{k}: {v}", fill="black", font=FONT)
        pos = (pos[0], pos[1] + TEXT_Y_DIFF)


def draw_path(draw: ImageDraw.ImageDraw, points: list[int], n: int):
    """Get frames lazily."""
    for p1, p2 in zip(points, points[1:]):
        draw_edge(draw, p1, p2, n)


def draw_edge(draw: ImageDraw.ImageDraw, start: int, finish: int, n: int, color):
    point1 = pi / 2 - 2 * pi * start / n
    point2 = pi / 2 - 2 * pi * finish / n
    d1 = (cos(point1), -sin(point1))
    d2 = (cos(point2), -sin(point2))
    p1 = (CENTER[0] + RADIUS * d1[0], CENTER[1] + RADIUS * d1[1])
    p2 = (CENTER[0] + RADIUS * d2[0], CENTER[1] + RADIUS * d2[1])
    draw.line([p1, p2], fill=color, width=STROKE)


def gen_frames(data: list[list[list[int]]], n: int, param: dict[str][str]):
    for gen in tqdm(data):
        with Image.new("RGB", SIZE, "white") as img:
            draw = ImageDraw.Draw(img)
            draw_generation(draw, gen, n, param)
            yield img


def main():
    if len(sys.argv) < 3:
        print(USAGE)
        sys.exit()
    fi = sys.argv[1]
    with open(fi, "r") as f:
        data = yaml.safe_load(f)
    n = data["city_count"]
    sim = data["simulation"]
    # shorten data for faster gif generation
    sim = sim[:: ceil(len(sim) / TOTAL_FRAMES)]

    param = {
        par: data[par]
        for par in [
            "city_count",
            "population",
            "mutation_rate",
            "crossover_rate",
            "crossover_group",
            "seed",
        ]
    }
    duration = int(TOTAL_DURATION / len(sim))
    frames = gen_frames(sim, n, param)
    next(frames).save(sys.argv[2], append_images=frames, duration=duration, loop=0)


if __name__ == "__main__":
    main()
