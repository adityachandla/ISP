## Elderly-family interaction
total_elderly=77000 # In 2020
elderly_visited_once_a_week = 0.6 * total_elderly
elderly_visited_once_a_month = 0.4 * total_elderly
total_visits_yearly = (elderly_visited_once_a_week*52) + \
        (elderly_visited_once_a_month*12)
average_visit_duration_hours = 2
total_visit_hours_yearly = average_visit_duration_hours * total_visits_yearly

## Nurses
total_pay_yearly = 32000
total_pay_per_shift = 32000/300 ## Assuming 300 working days

## Money
avoidable_shifts = total_visit_hours_yearly/8 ## Assuming 8 hour work day
profit = avoidable_shifts*total_pay_per_shift
print(f"Total profit {profit}")
profit_millions = profit//1000000
print(f"Profit in millions {profit_millions}")
