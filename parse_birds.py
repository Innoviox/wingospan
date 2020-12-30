import pdftotext

with open("birds.pdf", "rb") as f:
    pdf = pdftotext.PDF(f)

for page in pdf:
    for line in page.split("\n"):
         print(line)
