import re

text = "こんばんわ"
patterns = {
    "aisatsu":["こんにち[は|わ]", "こんばん[は|わ]", "よろしく"],
    "orei":["ありがとう", "[T|t]hank", "[T|t]hx", "どうも"]
}
for pt in patterns["aisatsu"]:
    if re.match(pt, text):
        print("やっふううううう")
        break
    else:
        continue
