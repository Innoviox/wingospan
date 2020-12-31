def generate(s):
    _, *fields, _ = s.split("\n")
    typ = fields[0].split()[1]
    fields = [i.split()[0] for i in fields]
    k = typ[0].lower()
    s = f"func ({k} {typ}) String() string"
    s += " {\n\treturn[...]string{"
    for i in fields:
        s += '"'+i+'",'
    s = s[:-1] + "}" + f"[{k}]\n" + "}"
    return s

