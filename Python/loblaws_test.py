# %%

# Write a function to return the factorial of a given number?


# Factorial 5! = 5.4.3.2.1 = 120
# """
# Algorithm


# """
# range 1-1000

# -----------
seen = {}


def factorial(fact: int) -> int:
    """
    input:
    fact: Number to be factorialized

    Rerturn:

    factorialzed number
    """

    if fact == 0:
        return 1

    elif fact in seen.keys():
        return seen[fact]

    else:
        seen[fact] = fact*factorial(fact-1)
        return seen[fact]

# %%

# -----------------------------------------------

# Given a non empty list of reviews and a set of keywords, return the K most frequent keywords.

# If two words have the same frequency, then the word with the lower alphabetical order comes first.
# Note: Multiple occurrences of a keyword within the same review only counts as one. 
# Input:


# k = 2


# Output:
# [“zehrs”, “loblaws ”]


# Explanation:
# “zehrs” occurs in 2 different reviews and “loblaws” and “nofrills” both occur in 1 review each, “loblaws” is returned as per the alphabetical order.

def word_count(review: str) -> dict:
    indi_words = review.split(" ")

    set = {}

    for word in indi_words:
        if word in set:
            pass
        elif word in count_dict:
            set.add(word)
            count_dict[word] += 1

        elif len(set) == len(keywords):
            break


if __name__ == __main__:

    keywords = ['loblaws', 'nofrills', 'zehrs', 'shoppers']

    count_dict = {}

    for i in keywords:
        count_dict[i] = 0

    for indi_review in reviews:
        word_count(indi_review)


# %%
