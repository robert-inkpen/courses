# %% Setup
# This is a catalog of items, categories and subcategories.

item_catalog = [
    {
        "item_id": 1,
        "item_description": "Water Bottle Large",
        "category_name": "Beverages",
        "sub_category": "Water"
    },
    {
        "item_id": 2,
        "item_description": "Soft Drink",
        "category_name": "Beverages",
        "sub_category": "Soda"
    },
    {
        "item_id": 3,
        "item_description": "Pound Cake",
        "category_name": "Cakes",
        "sub_category": "Fresh Cakes"
    },
    {
        "item_id": 4,
        "item_description": "Water Bottle Small",
        "category_name": "Beverages",
        "sub_category": "Water"
    },
]

# Below is a set of transcations

transcations = [
    {
        "week": 1,
        "transaction_id": 0,
        "item_id": 1,
        "spend": 10240
    },
    {
        "week": 1,
        "transaction_id": 1,
        "item_id": 1,
        "spend": 8500
    },
    {
        "week": 1,
        "transaction_id": 2,
        "item_id": 2,
        "spend": 7000
    },
    {
        "week": 2,
        "transaction_id": 0,
        "item_id": 1,
        "spend": 2300
    },
    {
        "week": 2,
        "transaction_id": 1,
        "item_id": 3,
        "spend": 1300
    },
    {
        "week": 2,
        "transaction_id": 2,
        "item_id": 4,
        "spend": 4500
    },
    {
        "week": 3,
        "transaction_id": 0,
        "item_id": 2,
        "spend": 2500
    }
]


# %% Q1 - Caluclate the % difference in spend between week 1 and week 2
spend_by_week = {}


def percent_difference(second_item: float, first_item: float) -> float:
    return ((second_item - first_item)/first_item)*100


for i in range(len(transcations)):
    if transcations[i]['week'] in spend_by_week:
        spend_by_week[transcations[i]['week']] += transcations[i]['spend']
    else:
        spend_by_week[transcations[i]['week']] = transcations[i]['spend']


diff = percent_difference(spend_by_week[2], spend_by_week[1])

print(f'spend_by_week: {spend_by_week}')

if diff > 0:
    print(
        f'There was {abs(diff):.2f} increase from week 1 to week 2')
else:
    print(
        f'There was {abs(diff):.2f} decrease from week 1 to week 2')


# %% # Q2 - Which item generates the most sales?
item_spend = {}

for trans in transcations:
    for items in item_catalog:
        if items["item_id"] == trans['item_id']:
            item_name = items['item_description']

    if item_name in item_spend:
        item_spend[item_name] += trans['spend']
    else:
        item_spend[item_name] = trans['spend']

# loop through item_spend and find the item with the highest spend
highest_spend = 0
highest_spend_item = ''
for key, value in item_spend.items():
    if value > highest_spend:
        highest_spend = value
        highest_spend_item = key

print(highest_spend_item)


# %% Q3 -Return the sub-categories with the most sales over 3000.
sale_limit = 3000
category_spend = {}

for trans in transcations:
    if trans['spend'] < 3000:
        break
    for items in item_catalog:
        if items["item_id"] == trans["item_id"]:
            item_category = items["sub_category"]

    if item_category in category_spend:
        category_spend[item_category] += trans['spend']
    else:
        category_spend[item_category] = trans['spend']

highest_spend = 0
highest_spend_category = ''

for key, value in category_spend.items():
    if value > highest_spend:
        highest_spend = value
        highest_spend_item = key

print(highest_spend_item)
# %%
