import re

def main():
    lines = []
    num = 0
    with open("input", 'r') as fp:
        for line in fp:
            lines += [line.strip()]

    patt = "XMAS"

    # horizontal and reverse
    for line in lines:
        num += len(re.findall(patt, line))
        num += len(re.findall(patt, line[::-1]))

    # vertical and reverse
    for n in range(len(lines[0])):
        h = ""
        for line in lines:
            h += line[n]
        num += len(re.findall(patt, h))
        num += len(re.findall(patt, h[::-1]))

    # diagonal bottom left to top right and reverse
    w = len(lines[0])
    h = len(lines)
    diagonal = ["" for _ in range(w + h - 1)] 
    for i in range(h): 
        for j in range(w): 
            diagonal[i + j] += lines[j][i]

    for d in diagonal:
        num += len(re.findall(patt, d))
        num += len(re.findall(patt, d[::-1]))

    # diagonal top left to bottom right and reverse
    w = len(lines[0])
    h = len(lines)
    diagonal = ["" for _ in range(w + h - 1)] 
    for i in range(h): 
        for j in range(w): 
            diagonal[j - i + (h - 1)] += lines[j][i]

    for d in diagonal:
        num += len(re.findall(patt, d))
        num += len(re.findall(patt, d[::-1]))

    print(num)

if __name__ == "__main__":
    main()
