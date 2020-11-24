# %%
import random


def check_for_vaccine():
    check = random.randint(0, 11)
    if check == 3:
        return True
    else:
        pass


def ontario_health_plan():
    covid_exists = True
    cure = False
    infection_rate = 0

    while covid_exists:
        infection_rate += 0.005

        if infection_rate > 0.025:
            infection_rate = 0.01
            print('lock down everything')
            cure = check_for_vaccine()
        elif cure:
            print('cure found: covid over')
            break
        else:
            cure = check_for_vaccine()
            print('doing fuck all')
            pass


ontario_health_plan()
# %%
