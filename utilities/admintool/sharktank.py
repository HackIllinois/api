import sys
import random
import json
import argparse

import api

tier_thresholds = {}

def print_verbose(msg):
    if args.verbose:
        print(msg, file=sys.stderr, end="")

def get_tier_by_points(pts):
    for tier_name, threshold in sorted(tier_thresholds.items(), key = lambda x: x[1], reverse=True):
        if pts >= threshold:
            return tier_name
    
    return min(tier_thresholds, key=tier_thresholds.get) # lowest tier

def select_n_profiles_per_tier(n):
    print_verbose('Getting all profiles ... ')

    raw_profile_data, ok = api.make_request('GET', '/profile/list/')

    if not ok:
        print('Error occurred when trying to get all profiles', file=sys.stderr)
        sys.exit(1)
    
    print_verbose('Success\n')

    print_verbose('Getting tier thresholds ... ')

    raw_tier_thresholds, ok = api.make_request('GET', '/profile/tier/threshold/')

    if not ok:
        print('Error occurred when trying to get tier thresholds', file=sys.stderr)
        sys.exit(1)
    
    print_verbose('Success\n')

    for t in raw_tier_thresholds:
        tier_thresholds[t['name']] = t['threshold']

    profiles_by_tier = {}

    for p in raw_profile_data['profiles']:
        tier_name = get_tier_by_points(p['points'])
        if tier_name not in profiles_by_tier:
            profiles_by_tier[tier_name] = []
        profiles_by_tier[tier_name].append(p)
    
    print_verbose(f'Selecting at most {n} people from each tier ... ')

    chosen_profiles = {}

    for tier, profiles in sorted(profiles_by_tier.items(), key = lambda x: tier_thresholds[x[0]]):
        chosen_profiles[tier] = random.sample(profiles, min(n, len(profiles)))

    print_verbose('Success\n')
    print_verbose('Done\n')

    return chosen_profiles

if __name__ == '__main__':
    global args

    parser = argparse.ArgumentParser(description='Randomly pick a handful of people from each tier.')
    parser.add_argument('num_to_select', metavar='N', type=int,
                        help='number of people to select at most per tier')
    parser.add_argument("-v", "--verbose", help="change output verbosity", 
                        action = "store_true")
    
    args = parser.parse_args()

    chosen_profiles = select_n_profiles_per_tier(args.num_to_select)

    print(json.dumps(chosen_profiles, indent=4))
