"""Snowball module that calculates how much return on investment a
dividend stock will payout after certain periods, assuming the dividend
yield does not change
"""
from typing import NamedTuple, Union


# Type Definition for either int or float
Money = Union[int, float]


class ReinvestResult(NamedTuple):
    """ReinvestResult provides a clean return type containing
    the result properties from the dividend_reinvestment function

    Args:
        run (bool): ???
        final_shares (int): Final number of shares
        final_revenue (Money): Final revenue
        years (int): Number of years invested
        payout_terms (int): Number of payout terms
        stock_price (Money): Initial stock price either as int or float
        regular_additions (int): ???

    Returns:
        ReinvestResult
    """
    run: bool
    final_shares: int
    final_revenue: Money
    years: int
    payout_terms: int
    stock_price: Money
    regular_additions: int


# Snowball Effect
def dividend_reinvestment(
        starting_stocks: int,
        stock_price: Money,
        dividend_payout: int,
        payout_frequency: int,
        years: int,
        regular_additions: int
) -> ReinvestResult:
    """
    Takes in the information about a stock to easily calculate how quickly
    a stock with reinvestment will grow it's share size.
    Input:
        starting_stocks (int): The amount of stocks that are in your possession
        stock_price (float): The average/projected stock price for the year
        dividend_payout (float): The amount in dollars that that is paid
            per stock
        payout_frequency (int): The frequency, in number of months
            between payments,
            that a dividend is paid.
        years (int): The amount of years for this to continue
        regular_additions (int): The amount, if any, you plan to
            contribute between
        payout frequency (int): 0 if no additions to be made.
    Return:
        ReinvestResult
    """
    bank = 0
    new_stocks = 0
    if starting_stocks > 0:
        print('Stating with {} stocks, costing ${} while adding ${} per period yields:'.format(  # noqa
            starting_stocks, starting_stocks*stock_price, regular_additions))
        run = True
        payout_terms = (12/payout_frequency)*years
        for i in range(int(payout_terms)):
            initial_revenue = starting_stocks*dividend_payout
            new_revenue = new_stocks*dividend_payout
            bank = round((bank + initial_revenue +
                          new_revenue + regular_additions), 2)
            if bank >= stock_price:
                remainder = round((bank % stock_price), 2)
                # print(remainder)
                new_stocks = new_stocks+((bank-remainder)/stock_price)
                # print("new stocks : {} for run {}".format(new_stocks, i))
                bank = remainder
                # print("bank amount: {} for run {}".format(bank, i))
        final_shares = starting_stocks+new_stocks
        final_revenue = final_shares*dividend_payout

    else:
        run = False
        final_shares = 0
        final_revenue = 0

    result = ReinvestResult(run, final_shares, final_revenue,
                            years, payout_terms, stock_price,
                            regular_additions)
    return result


if __name__ == "__main__":
    res = dividend_reinvestment(  # noqa: E501
        starting_stocks=1,
        stock_price=12,
        dividend_payout=0.09,
        payout_frequency=1,
        years=20,
        regular_additions=25
    )

    if res.run is True:
        print('a final share count after {} years of {}'.format(
            res.years, res.final_shares))
        print('final revenue per month as ${}'.format(
            round(res.final_revenue, 2)))
        print("The final value of your portfolio is ${}".format(
            res.stock_price*res.final_shares
        ))
        print("With the value put into your account being ${}".format(res.payout_terms*res.regular_additions))  # noqa
    else:
        print('No starting shares, wrong tool')
