import sys
import random
import json
import argparse
import numpy as np

import api

tier_thresholds = {}

def print_verbose(msg):
    if args.verbose:
        print(msg, file=sys.stderr, end="")

def get_top_10_and_select_tiers(profiles, top_n, m):
    leftover_profiles = profiles.copy()

    leftover_profiles.sort(key = lambda x: x['points'], reverse = True)

    points = [p['points'] for p in leftover_profiles]

    profiles_by_tier = {}

    total_profiles = len(leftover_profiles)

    if total_profiles <= top_n:
        print_verbose(f'Less than {top_n} people total so everyone makes top {top_n}!\n')
        profiles_by_tier["top10"] = leftover_profiles
        return profiles_by_tier

    print_verbose(f'Generating top {top_n} ... ')
    l = top_n - np.searchsorted(points[top_n - 1::-1], points[top_n - 1], side='right')
    r = total_profiles - np.searchsorted(points[:top_n - 1:-1], points[top_n - 1], side='left')

    top10 = leftover_profiles[:l]

    leftover_profiles = leftover_profiles[l:]

    if r > top_n:
        print_verbose(f'{r - l} people tied, but cannot fit all of them within top 10 (only {top_n - l} places left)')
        if top_n - l > 1:
            print_verbose(f'. Randomly selecting {top_n - l} people ... ')
        else:
            print_verbose(f'. Randomly selecting 1 person ... ')

    for idx in sorted(random.sample(range(r - l), k = top_n - l), reverse=True):
        top10.append(leftover_profiles[idx])
        leftover_profiles.pop(idx)

    points = points[top_n:]

    profiles_by_tier["top10"] = top10

    print_verbose(f'Success\n')

    for tier_name, thres in sorted(tier_thresholds.items(), key = lambda x: x[1], reverse=True):
        print_verbose(f'Attempting to select {m} people from the {tier_name} tier and above ... ')
        r =  len(points) - np.searchsorted(points[::-1], thres, side='left')
        profiles_by_tier[tier_name] = []
        actual_k = min(m, r)
        if r != m and actual_k == r:
            print_verbose(f"Not enough people left to select {m}. ")
            if r > 1:
                print_verbose(f"{r} people can be picked, so they will all be selected ... ")
            if r == 1:
                print_verbose(f"1 person can be picked, so they will be selected ... ")
            else:
                print_verbose("0 people can be picked, so the tier will be empty ... ")
        for i in sorted(random.sample(range(r), actual_k), reverse=True):
            profiles_by_tier[tier_name].append(leftover_profiles[i])
            leftover_profiles.pop(i)
            points.pop(i)

        print_verbose(f'Success\n')

    return profiles_by_tier

def generate(top_n, m):
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

    profiles_by_tier = get_top_10_and_select_tiers(raw_profile_data['profiles'], top_n, m)

    print_verbose('Done\n')

    return profiles_by_tier

if __name__ == '__main__':
    global args

    parser = argparse.ArgumentParser(description='''Generate the top N users that have the most points,
                                                    and randomly pick a handful of people from each tier.
                                                    Outputs a JSON to stdout.''')
    parser.add_argument('top_n', metavar='N', type=int,
                        help='top N number of people to select based on overall points')
    parser.add_argument('num_to_select', metavar='M', type=int,
                        help='Randomly select at most M number of people per tier')
    parser.add_argument("-v", "--verbose", help="change output verbosity", 
                        action = "store_true")
    
    args = parser.parse_args()

    chosen_profiles = generate(args.top_n, args.num_to_select)

    print(json.dumps(chosen_profiles, indent=4))
