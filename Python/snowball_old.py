# %% Snowball Effect
def dividend_reinvestment(
    starting_stocks: int,
    stock_price: int,
    dividend_payout: int,
    payout_frequency: int,
    years: int,
    regular_additions: int
):
    """
        Takes in the information about a stock to easily calculate how quickly
        a stock with reinvestment will grow it's share size.

        Input:
        starting_stocks: The amount of stocks that are in your possession

        stock_price: The average/projected stock price for the year

        dividend_payout: The amount in dollars that that is paid per stock

        payout_frequency: The frequency, in number of months between payments,
        that a dividend is paid.

        years: The amount of years for this to continue

        regular_additions: The amount, if any, you plan to contribute between
        payout frequency. 0 if no additions to be made.

        Return:
        final_shares: The total amount of shares owned at the end of the period

        final_revenue: The final amount that the dividends will be paying out
        per period.

        """
    # print(starting_stocks)
    # print(stock_price)
    # print(dividend_payout)
    # print(payout_frequency)
    # print(years)
    bank = 0
    new_stocks = 0

    if starting_stocks > 0:
        print('Stating with {} stocks, costing {}$ while adding {}$ per period yields:'.format(  # noqa
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

    return (run, final_shares, final_revenue, years, payout_terms, stock_price, regular_additions)  # noqa


run, shares, revenue, years, periods, price, additions = dividend_reinvestment(
    starting_stocks=0,
    stock_price=14.77,
    dividend_payout=0.09,
    payout_frequency=1,
    years=20,
    regular_additions=25

)

if run is True:
    print('a final share count after {} years of {}'.format(years, shares))
    print('final revenue per month as {}$'.format(
        round(revenue, 2)))
    print("The final value of your portfolio is {}".format(price*shares))
    print("With the value put into your account being {}".format(periods*additions))  # noqa
else:
    print('No starting shares, wrong tool')

# %%
