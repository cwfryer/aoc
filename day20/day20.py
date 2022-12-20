def _mix(tuples: list[tuple[int, int]]):
    for i in range(len(tuples)):
        i, entry = next((j, entry) for j, entry in enumerate(tuples) if entry[0] == i)
        j = (i + entry[1]) % (len(tuples) - 1)
        if i < j:
            tuples[i:j] = tuples[i + 1 : j + 1]
        elif i > j:
            tuples[j + 1 : i + 1] = tuples[j:i]
        tuples[j] = entry


def main():
    encFile = []
    with open("day20/day20.input") as f:
        for x in f:
            encFile.append(int(x))

    tuples = list(enumerate(811589153*x for x in encFile))
    for _ in range(10):
        _mix(tuples)
    i = next(i for i, entry in enumerate(tuples) if entry[1] == 0)
    answer = sum(tuples[(i + x) % len(tuples)][1] for x in range(1000, 3001, 1000))
    print(answer)


if __name__ == "__main__":
    main()
