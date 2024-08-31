# **Stable Marriage Problem**

## **Overview:**

The Stable Marriage Problem involves matching members of two equal-sized sets based on preference rankings. The goal is to find a stable match where no pair of members would prefer each other over their current partners.

## **Python Example:**

```python
def stable_marriage(men_prefs, women_prefs):
    men = list(men_prefs.keys())
    women = list(women_prefs.keys())
    free_men = set(men)
    proposals = {man: [] for man in men}
    matches = {}

    while free_men:
        man = free_men.pop()
        woman = men_prefs[man][0]
        proposals[man].append(woman)
        men_prefs[man].pop(0)

        if woman not in matches:
            matches[woman] = man
        else:
            current_partner = matches[woman]
            if women_prefs[woman].index(man) < women_prefs[woman].index(current_partner):
                matches[woman] = man
                free_men.add(current_partner)
            else:
                free_men.add(man)
    
    return matches

# Example usage
men_prefs = {'A': ['X', 'Y', 'Z'], 'B': ['Y', 'X', 'Z'], 'C': ['Z', 'Y', 'X']}
women_prefs = {'X': ['B', 'A', 'C'], 'Y': ['A', 'B', 'C'], 'Z': ['C', 'A', 'B']}
print("Stable Matches:", stable_marriage(men_prefs, women_prefs))
```

## **Explanation:**
- **Stable Marriage Problem:** Uses the Gale-Shapley algorithm where each man proposes to women in his preference list and each woman tentatively accepts the best option based on her preference list, ensuring stability.

