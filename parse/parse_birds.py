import pdftotext
import csv

with open("birds.pdf", "rb") as f:
    pdf = pdftotext.PDF(f)


with open('birds.csv', 'w') as out:
    fieldnames = ['Name', 'Region', 'Cost', 'Points', 'Nest', 'Eggs', 'Wingspan', 'Action']

    writer = csv.DictWriter(out, fieldnames=fieldnames)
    writer.writeheader()

    for page in pdf:
        for line in page.split("\n"):
            bird = line.split()

            if not bird: continue
            if line.startswith("common"): continue

            try:         
                name = ''
                i = 0
                while not bird[i].isnumeric():
                    name += bird[i] + ' '
                    i += 1
                name = name.strip("+ /")
                points = bird[i]
                wingspan = bird[i + 1]
                action = ' '.join(bird[i + 3:])
                writer.writerow({'Name': name, 'Points': points, 'Wingspan': wingspan, 'Action': action})
            except:
                print("exception on line", line)
                input()
